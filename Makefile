TAG = $(shell git tag -l --points-at HEAD)
COMMIT = $(shell git rev-parse --short HEAD)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
BUILD_TIME = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

CORE_MOD_NAME = github.com/gbh007/hgraber-next-agent-core
LDFLAGS = -ldflags "-X '$(CORE_MOD_NAME)/version.Version=$(TAG)' -X '$(CORE_MOD_NAME)/version.Commit=$(COMMIT)' -X '$(CORE_MOD_NAME)/version.BuildAt=$(BUILD_TIME)' -X '$(CORE_MOD_NAME)/version.Branch=$(BRANCH)'"


create_build_dir:
	mkdir -p ./_build

.PHONY: build
build: create_build_dir
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -trimpath -o ./_build/agent-linux-arm64  ./cmd/agent
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -trimpath -o ./_build/agent-linux-amd64  ./cmd/agent
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -trimpath -o ./_build/agent-windows-amd64.exe  ./cmd/agent

.PHONY: docker
docker: build
	docker build -f Dockerfile \
		--build-arg "BINARY_PATH=./_build/agent-linux-arm64" \
		-t hgraber-next-agent:arm64 .
	docker save hgraber-next-agent:arm64 -o _build/hgraber-next-agent_arm64.tar

	docker build -f Dockerfile \
		--build-arg "BINARY_PATH=./_build/agent-linux-amd64" \
		-t hgraber-next-agent:amd64 .
	docker save hgraber-next-agent:amd64 -o _build/hgraber-next-agent_amd64.tar

.PHONY: run-example
run-example: create_build_dir
	CGO_ENABLED=0 go build $(LDFLAGS) -trimpath -o ./_build/hgraber-agent ./cmd/agent

	./_build/hgraber-agent --config config-example.yaml

.PHONY: run
run: create_build_dir
	CGO_ENABLED=0 go build $(LDFLAGS) -trimpath -o ./_build/hgraber-agent ./cmd/agent

	./_build/hgraber-agent

.PHONY: scan
scan: create_build_dir
	CGO_ENABLED=0 go build $(LDFLAGS) -trimpath -o ./_build/hgraber-agent ./cmd/agent

	./_build/hgraber-agent --scan

.PHONY: update-dep
update-dep:
	go get -u github.com/gbh007/hgraber-next@master