TAG = $(shell git tag -l --points-at HEAD)
COMMIT = $(shell git rev-parse --short HEAD)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
BUILD_TIME = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

CORE_MOD_NAME = github.com/gbh007/hgraber-next-agent-core
LDFLAGS = -ldflags "-X '$(CORE_MOD_NAME)/version.Version=$(TAG)' -X '$(CORE_MOD_NAME)/version.Commit=$(COMMIT)' -X '$(CORE_MOD_NAME)/version.BuildAt=$(BUILD_TIME)' -X '$(CORE_MOD_NAME)/version.Branch=$(BRANCH)'"

SERVICE_BIN = $(PWD)/_build

create_build_dir:
	mkdir -p $(SERVICE_BIN)

.PHONY: build
build: create_build_dir
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -trimpath -o $(SERVICE_BIN)/agent-linux-arm64  ./cmd/agent
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -trimpath -o $(SERVICE_BIN)/agent-linux-amd64  ./cmd/agent
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -trimpath -o $(SERVICE_BIN)/agent-windows-amd64.exe  ./cmd/agent
	CGO_ENABLED=0 go build $(LDFLAGS) -trimpath -o $(SERVICE_BIN)/hgraber-agent ./cmd/agent

.PHONY: docker
docker: build
	docker build -f Dockerfile \
		--build-arg "BINARY_PATH=$(SERVICE_BIN)/agent-linux-arm64" \
		-t hgraber-next-agent:arm64 .
	docker save hgraber-next-agent:arm64 -o _build/hgraber-next-agent_arm64.tar

	docker build -f Dockerfile \
		--build-arg "BINARY_PATH=$(SERVICE_BIN)/agent-linux-amd64" \
		-t hgraber-next-agent:amd64 .
	docker save hgraber-next-agent:amd64 -o _build/hgraber-next-agent_amd64.tar

.PHONY: run-example
run-example: build
	$(SERVICE_BIN)/hgraber-agent --config config-example.yaml

.PHONY: run
run: build
	$(SERVICE_BIN)/hgraber-agent

.PHONY: scan
scan: build
	$(SERVICE_BIN)/hgraber-agent --scan


.PHONY: config
config: create_build_dir
	go build $(LDFLAGS) -trimpath -o $(SERVICE_BIN)/configremaper  ./cmd/configremaper

	$(SERVICE_BIN)/configremaper --out config-generated.yaml
	$(SERVICE_BIN)/configremaper --out config-generated.env
	$(SERVICE_BIN)/configremaper --out config-generated.toml


.PHONY: update-dep
update-dep:
	go get -u github.com/gbh007/hgraber-next@master
	go mod tidy