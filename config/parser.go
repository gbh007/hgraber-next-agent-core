package config

import "time"

type DebugServer struct {
	Addr            string `yaml:"addr" envconfig:"ADDR"`
	LogErrorHandler bool   `yaml:"log_error_handler" envconfig:"LOG_ERROR_HANDLER"`
	Debug           bool   `yaml:"debug" envconfig:"DEBUG"`
}

func DefaultDebugServer() DebugServer {
	return DebugServer{}
}

type Parsers struct {
	HG4Token string       `yaml:"hg4_token" envconfig:"HG4_TOKEN"`
	Enabled  []string     `yaml:"enabled" envconfig:"ENABLED"`
	Cache    ParsersCache `yaml:"cache" envconfig:"CACHE"`
}

func DefaultParsers() *Parsers {
	return &Parsers{
		HG4Token: "",
		Enabled: []string{
			"mock",
			"hgraber_local",
		},
		Cache: DefaultParsersCache(),
	}
}

type ParsersCache struct {
	Enabled       bool          `yaml:"enabled" envconfig:"ENABLED"`
	Path          string        `yaml:"path" envconfig:"PATH"`
	TTL           time.Duration `yaml:"ttl" envconfig:"TTL"`
	CleanInterval time.Duration `yaml:"clean_interval" envconfig:"CLEAN_INTERVAL"`
}

func DefaultParsersCache() ParsersCache {
	return ParsersCache{
		TTL:           time.Hour,
		CleanInterval: time.Minute * 5,
	}
}
