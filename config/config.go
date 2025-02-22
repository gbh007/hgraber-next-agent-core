package config

import "time"

type Config[T any] struct {
	Log         Log         `yaml:"log" envconfig:"LOG"`
	API         API         `yaml:"api" envconfig:"API"`
	DebugServer DebugServer `yaml:"debug_server" envconfig:"DEBUG_SERVER"`
	Application Application `yaml:"application" envconfig:"APPLICATION"`
	Parsers     *T          `yaml:"parsers" envconfig:"PARSERS"`
	FSBase      FSBase      `yaml:"fs_base" envconfig:"FS_BASE"`
	Sqlite      Sqlite      `yaml:"sqlite" envconfig:"SQLITE"`
	ZipScanner  ZipScanner  `yaml:"zip_scanner" envconfig:"ZIP_SCANNER"`
	Highway     Highway     `yaml:"highway" envconfig:"HIGHWAY"`
}

func DefaultConfig[T any](defaultParsers func() *T) Config[T] {
	return Config[T]{
		Log:         LogDefault(),
		API:         DefaultAPI(),
		DebugServer: DefaultDebugServer(),
		Parsers:     defaultParsers(),
		Application: DefaultApplication(),
		FSBase:      DefaultFSBase(),
		Sqlite:      DefaultSqlite(),
		Highway:     DefaultHighway(),
	}
}

type Application struct {
	TraceEndpoint   string        `yaml:"trace_endpoint" envconfig:"TRACE_ENDPOINT"`
	ClientTimeout   time.Duration `yaml:"client_timeout" envconfig:"CLIENT_TIMEOUT"`
	ServiceName     string        `yaml:"service_name" envconfig:"SERVICE_NAME"`
	UseUnsafeCloser bool          `yaml:"use_unsafe_closer" envconfig:"USE_UNSAFE_CLOSER"`
	Pyroscope       Pyroscope     `yaml:"pyroscope" envconfig:"PYROSCOPE"`
}

type Pyroscope struct {
	Endpoint string `yaml:"endpoint" envconfig:"ENDPOINT"`
	Debug    bool   `yaml:"debug" envconfig:"DEBUG"`
	Rate     int    `yaml:"rate" envconfig:"RATE"`
}

func DefaultApplication() Application {
	return Application{
		TraceEndpoint: "",
		ClientTimeout: time.Minute,
		ServiceName:   "hgraber-next-agent",
	}
}
