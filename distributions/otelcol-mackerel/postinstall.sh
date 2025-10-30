#!/bin/sh

if command -v systemctl >/dev/null 2>&1; then
  if [ -d /run/systemd/system ]; then
    systemctl daemon-reload
  fi
  if [ -d /run/systemd/system ]; then
    systemctl enable --now otelcol-mackerel.service
  fi
fi
