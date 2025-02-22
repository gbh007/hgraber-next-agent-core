package config

type DebugServer struct {
	Addr            string `yaml:"addr" envconfig:"ADDR"`
	LogErrorHandler bool   `yaml:"log_error_handler" envconfig:"LOG_ERROR_HANDLER"`
	Debug           bool   `yaml:"debug" envconfig:"DEBUG"`
}

func DefaultDebugServer() DebugServer {
	return DebugServer{}
}

type Parsers struct {
	HG4Token string   `yaml:"hg4_token" envconfig:"HG4_TOKEN"`
	Enabled  []string `yaml:"enabled" envconfig:"ENABLED"`
}

func DefaultParsers() *Parsers {
	return &Parsers{
		HG4Token: "",
		Enabled: []string{
			"mock",
			"hgraber_local",
		},
	}
}
