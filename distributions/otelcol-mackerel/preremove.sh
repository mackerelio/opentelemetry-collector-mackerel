#!/bin/sh

if [ "$1" != "1" ]; then
  if command -v systemctl >/dev/null 2>&1; then
    systemctl stop otelcol-mackerel.service
    systemctl disable otelcol-mackerel.service
  fi
fi
