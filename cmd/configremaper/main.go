package main

import (
	"github.com/gbh007/hgraber-next-agent-core/config"
	"github.com/gbh007/hgraber-next/application/configremaper"
)

func main() {
	configremaper.Run(config.DefaultConfigWrapped(config.DefaultParsers))
}
