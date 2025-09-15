#!/bin/sh

set -e

BIN_NAME="fndn"
for GOOS in linux windows darwin; do
  for GOARCH in amd64 arm64; do
    OUTDIR="./bin/dist/${GOOS}_${GOARCH}"
    mkdir -p "$OUTDIR"
    OUTFILE="${OUTDIR}/${BIN_NAME}"
    [ "$GOOS" = "windows" ] && OUTFILE="${OUTFILE}.exe"
    CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -ldflags "-s -w" -o "$OUTFILE" ./cmd/fndn
    if ! { [ "$GOOS" = "windows" ] && [ "$GOARCH" = "arm64" ]; } && [ "$GOOS" != "darwin" ]; then
      upx --best --lzma "$OUTFILE"
    fi
  done
done

