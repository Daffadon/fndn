package config_template

const BuildConfigTemplate string = `#!/bin/bash

set -e

v=$(cat VERSION)
docker build -t {{.ModuleName}}:$v .
`

const BinaryBuildConfigTemplate string = `#!/bin/bash

set -e

BIN_NAME={{.ProjectName}}
for GOOS in linux windows darwin; do
  for GOARCH in amd64 arm64; do
    OUTDIR="./bin/dist/${GOOS}_${GOARCH}"
    mkdir -p "$OUTDIR"
    OUTFILE="${OUTDIR}/${BIN_NAME}"
    [ "$GOOS" = "windows" ] && OUTFILE="${OUTFILE}.exe"
    CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -ldflags "-s -w" -o "$OUTFILE" ./cmd/
  done
done

`
