package config

type API struct {
	Addr            string `toml:"addr" yaml:"addr" envconfig:"ADDR"`
	Token           string `toml:"token" yaml:"token" envconfig:"TOKEN"`
	LogErrorHandler bool   `toml:"log_error_handler" yaml:"log_error_handler" envconfig:"LOG_ERROR_HANDLER"`
	Debug           bool   `toml:"debug" yaml:"debug" envconfig:"DEBUG"`
}

func (a API) GetAddr() string {
	return a.Addr
}

func (a API) GetToken() string {
	return a.Token
}

func (a API) GetLogErrorHandler() bool {
	return a.LogErrorHandler
}

func (a API) GetDebug() bool {
	return a.Debug
}

func DefaultAPI() API {
	return API{
		Addr:  ":8080",
		Token: "",
	}
}
