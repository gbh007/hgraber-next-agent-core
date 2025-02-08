#! /bin/bash

TAG=$(git tag -l --points-at HEAD)
COMMIT=$(git rev-parse --short HEAD)
BRANCH=$(git rev-parse --abbrev-ref HEAD)
BUILD_TIME=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

CORE_MOD_NAME="github.com/gbh007/hgraber-next-agent-core"
LDFLAGS="-X '${CORE_MOD_NAME}/version.Version=${TAG}' -X '${CORE_MOD_NAME}/version.Commit=${COMMIT}' -X '${CORE_MOD_NAME}/version.BuildAt=${BUILD_TIME}' -X '${CORE_MOD_NAME}/version.Branch=${BRANCH}'"

go build -ldflags "${LDFLAGS}" -trimpath -o ./_build/agent  ./cmd/agent

CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "${LDFLAGS}" -trimpath -o ./_build/agent-linux-arm64  ./cmd/agent
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "${LDFLAGS}" -trimpath -o ./_build/agent-linux-amd64  ./cmd/agent
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "${LDFLAGS}" -trimpath -o ./_build/agent-windows-amd64.exe  ./cmd/agent