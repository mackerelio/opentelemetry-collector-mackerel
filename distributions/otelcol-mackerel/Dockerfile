# Build with GoReleaser
# cf.) https://goreleaser.com/customization/dockers_v2/

FROM gcr.io/distroless/base-debian13:nonroot

ARG TARGETPLATFORM

COPY --chown=nonroot:nonroot $TARGETPLATFORM/otelcol-mackerel /usr/bin/

ENTRYPOINT ["/usr/bin/otelcol-mackerel"]
CMD ["--config", "mackerel:default"]

EXPOSE 4317 4318
