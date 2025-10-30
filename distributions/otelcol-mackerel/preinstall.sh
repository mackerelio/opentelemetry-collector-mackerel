#!/bin/sh

getent passwd otelcol-mackerel >/dev/null || useradd --system --user-group --no-create-home --shell /sbin/nologin otelcol-mackerel
