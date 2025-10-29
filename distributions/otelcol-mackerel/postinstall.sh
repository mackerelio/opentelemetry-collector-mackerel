#!/bin/sh

if command -v systemctl >/dev/null 2>&1; then
  if [ -d /run/systemd/system ]; then
    systemctl daemon-reload
  fi
  systemctl enable otelcol-mackerel.service
  if [ -d /run/systemd/system ]; then
    systemctl restart otelcol-mackerel.service
  fi
fi
