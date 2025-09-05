package config

import "time"

type Config[T any] struct {
	Log         Log         `toml:"log" yaml:"log" envconfig:"LOG"`
	API         API         `toml:"api" yaml:"api" envconfig:"API"`
	DebugServer DebugServer `toml:"debug_server" yaml:"debug_server" envconfig:"DEBUG_SERVER"`
	Application Application `toml:"application" yaml:"application" envconfig:"APPLICATION"`
	Parsers     *T          `toml:"parsers" yaml:"parsers" envconfig:"PARSERS"`
	FSBase      FSBase      `toml:"fs_base" yaml:"fs_base" envconfig:"FS_BASE"`
	Sqlite      Sqlite      `toml:"sqlite" yaml:"sqlite" envconfig:"SQLITE"`
	ZipScanner  ZipScanner  `toml:"zip_scanner" yaml:"zip_scanner" envconfig:"ZIP_SCANNER"`
	Highway     Highway     `toml:"highway" yaml:"highway" envconfig:"HIGHWAY"`
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

func DefaultConfigWrapped[T any](defaultParsers func() *T) func() Config[T] {
	return func() Config[T] {
		return DefaultConfig(defaultParsers)
	}
}

type Application struct {
	TraceEndpoint   string        `toml:"trace_endpoint" yaml:"trace_endpoint" envconfig:"TRACE_ENDPOINT"`
	ClientTimeout   time.Duration `toml:"client_timeout" yaml:"client_timeout" envconfig:"CLIENT_TIMEOUT"`
	ServiceName     string        `toml:"service_name" yaml:"service_name" envconfig:"SERVICE_NAME"`
	UseUnsafeCloser bool          `toml:"use_unsafe_closer" yaml:"use_unsafe_closer" envconfig:"USE_UNSAFE_CLOSER"`
	Pyroscope       Pyroscope     `toml:"pyroscope" yaml:"pyroscope" envconfig:"PYROSCOPE"`
}

type Pyroscope struct {
	Endpoint string `toml:"endpoint" yaml:"endpoint" envconfig:"ENDPOINT"`
	Debug    bool   `toml:"debug" yaml:"debug" envconfig:"DEBUG"`
	Rate     int    `toml:"rate" yaml:"rate" envconfig:"RATE"`
}

func DefaultApplication() Application {
	return Application{
		TraceEndpoint: "",
		ClientTimeout: time.Minute,
		ServiceName:   "hgraber-next-agent",
	}
}
