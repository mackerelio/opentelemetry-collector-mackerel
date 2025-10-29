#!/bin/sh

set -e

if [ -z "$TAGPR_CURRENT_VERSION" ]; then
  echo "Error: TAGPR_CURRENT_VERSION environment variable is not set"
  exit 1
fi

if [ -z "$TAGPR_NEXT_VERSION" ]; then
  echo "Error: TAGPR_NEXT_VERSION environment variable is not set"
  exit 1
fi

for manifest in distributions/*/manifest.yaml; do
  if [ -f "$manifest" ]; then
    sed -i '' "s|github.com/mackerelio/opentelemetry-collector-mackerel/\([^ ]*\) ${TAGPR_CURRENT_VERSION}|github.com/mackerelio/opentelemetry-collector-mackerel/\1 ${TAGPR_NEXT_VERSION}|g" "$manifest"
  fi
done

echo "Done."
