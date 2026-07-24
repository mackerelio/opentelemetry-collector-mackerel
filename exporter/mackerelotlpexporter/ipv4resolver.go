package mackerelotlpexporter

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc/resolver"
)

const ipv4ResolverScheme = "mdotipv4"

var minResolveInterval = 30 * time.Second

func init() {
	resolver.Register(&ipv4ResolverBuilder{})
}

type lookupIPFunc func(ctx context.Context, network, host string) ([]net.IP, error)

type ipv4ResolverBuilder struct {
	lookupIP        lookupIPFunc
	resolveInterval time.Duration
}

func (b *ipv4ResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, _ resolver.BuildOptions) (resolver.Resolver, error) {
	endpoint := target.Endpoint()
	host, port, err := net.SplitHostPort(endpoint)
	if err != nil {
		return nil, fmt.Errorf("ipv4 resolver: parse endpoint %q: %w", endpoint, err)
	}

	if net.ParseIP(host) != nil {
		if err := cc.UpdateState(resolver.State{
			Addresses: []resolver.Address{{Addr: endpoint}},
		}); err != nil {
			return nil, fmt.Errorf("ipv4 resolver: update state for endpoint %q: %w", endpoint, err)
		}
		return &deadResolver{}, nil
	}

	lookupIP := b.lookupIP
	if lookupIP == nil {
		lookupIP = net.DefaultResolver.LookupIP
	}

	resolveInterval := b.resolveInterval
	if resolveInterval <= 0 {
		resolveInterval = minResolveInterval
	}

	ctx, cancel := context.WithCancel(context.Background())
	r := &ipv4Resolver{
		host:       host,
		port:       port,
		ctx:        ctx,
		cancel:     cancel,
		cc:         cc,
		rn:         make(chan struct{}, 1),
		lookupIP:   lookupIP,
		interval:   resolveInterval,
		lastLookup: time.Now(),
	}
	r.lookup()
	r.wg.Add(1)
	go r.watcher()
	return r, nil
}

func (b *ipv4ResolverBuilder) Scheme() string {
	return ipv4ResolverScheme
}

type deadResolver struct{}

func (deadResolver) ResolveNow(resolver.ResolveNowOptions) {}
func (deadResolver) Close()                                {}

type ipv4Resolver struct {
	host       string
	port       string
	ctx        context.Context
	cancel     context.CancelFunc
	cc         resolver.ClientConn
	rn         chan struct{}
	lookupIP   lookupIPFunc
	interval   time.Duration
	lastLookup time.Time
	wg         sync.WaitGroup
}

// watcher re-resolves on ResolveNow signals, enforcing a minimum interval.
func (r *ipv4Resolver) watcher() {
	defer r.wg.Done()
	for {
		select {
		case <-r.ctx.Done():
			return
		case <-r.rn:
		}
		nextAllowed := r.lastLookup.Add(r.interval)
		if delay := time.Until(nextAllowed); delay > 0 {
			select {
			case <-r.ctx.Done():
				return
			case <-time.After(delay):
			}
		}
		r.lookup()
		r.lastLookup = time.Now()
	}
}

func (r *ipv4Resolver) lookup() {
	ips, err := r.lookupIP(r.ctx, "ip4", r.host)
	if err != nil {
		r.cc.ReportError(fmt.Errorf("ipv4 resolver: resolve %s: %w", r.host, err))
		return
	}
	if len(ips) == 0 {
		r.cc.ReportError(fmt.Errorf("ipv4 resolver: no IPv4 address found for %s", r.host))
		return
	}

	addrs := make([]resolver.Address, len(ips))
	for i, ip := range ips {
		addrs[i] = resolver.Address{Addr: net.JoinHostPort(ip.String(), r.port)}
	}
	if err := r.cc.UpdateState(resolver.State{Addresses: addrs}); err != nil {
		r.cc.ReportError(fmt.Errorf("ipv4 resolver: update state for %s: %w", r.host, err))
	}
}

func (r *ipv4Resolver) ResolveNow(resolver.ResolveNowOptions) {
	select {
	case r.rn <- struct{}{}:
	default:
	}
}

func (r *ipv4Resolver) Close() {
	r.cancel()
	r.wg.Wait()
}
