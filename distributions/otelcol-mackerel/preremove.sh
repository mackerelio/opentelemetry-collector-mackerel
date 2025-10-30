#!/bin/sh

if [ "$1" != "1" ]; then
  if command -v systemctl >/dev/null 2>&1; then
    systemctl disable --now otelcol-mackerel.service
  fi
fi
