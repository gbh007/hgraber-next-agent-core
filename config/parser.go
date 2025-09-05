package config

import "time"

type DebugServer struct {
	Addr            string `toml:"addr" yaml:"addr" envconfig:"ADDR"`
	LogErrorHandler bool   `toml:"log_error_handler" yaml:"log_error_handler" envconfig:"LOG_ERROR_HANDLER"`
	Debug           bool   `toml:"debug" yaml:"debug" envconfig:"DEBUG"`
}

func DefaultDebugServer() DebugServer {
	return DebugServer{}
}

type Parsers struct {
	HG4Token string       `toml:"hg4_token" yaml:"hg4_token" envconfig:"HG4_TOKEN"`
	Enabled  []string     `toml:"enabled" yaml:"enabled" envconfig:"ENABLED"`
	Cache    ParsersCache `toml:"cache" yaml:"cache" envconfig:"CACHE"`
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
	Enabled       bool          `toml:"enabled" yaml:"enabled" envconfig:"ENABLED"`
	Path          string        `toml:"path" yaml:"path" envconfig:"PATH"`
	TTL           time.Duration `toml:"ttl" yaml:"ttl" envconfig:"TTL"`
	CleanInterval time.Duration `toml:"clean_interval" yaml:"clean_interval" envconfig:"CLEAN_INTERVAL"`
}

func DefaultParsersCache() ParsersCache {
	return ParsersCache{
		TTL:           time.Hour,
		CleanInterval: time.Minute * 5,
	}
}
