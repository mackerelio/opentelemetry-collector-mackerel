#!/bin/sh

set -e

REMOTE_MAIN_SHA=$(git ls-remote origin main | awk '{print $1}')
CURRENT_SHA=$(git rev-parse HEAD)

if [ "$CURRENT_SHA" != "$REMOTE_MAIN_SHA" ]; then
  echo "Error: current commit ($CURRENT_SHA) does not match remote main HEAD ($REMOTE_MAIN_SHA)" >&2
  exit 1
fi

if [ -n "$(git status --porcelain)" ]; then
  echo "Error: working directory has uncommitted changes" >&2
  exit 1
fi

go tool -modfile=tool.mod multimod verify

for tag in $(go tool -modfile=tool.mod multimod tag -c "$CURRENT_SHA" -m beta -p --print-tags)
do
  git push origin "$tag"
done
