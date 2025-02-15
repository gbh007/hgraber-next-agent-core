package config

import "time"

type Config[T any] struct {
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
	Debug           bool          `yaml:"debug" envconfig:"DEBUG"`
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
		Debug:         false,
		TraceEndpoint: "",
		ClientTimeout: time.Minute,
		ServiceName:   "hgraber-next-agent",
	}
}

type API struct {
	Addr            string `yaml:"addr" envconfig:"ADDR"`
	Token           string `yaml:"token" envconfig:"TOKEN"`
	LogErrorHandler bool   `yaml:"log_error_handler" envconfig:"LOG_ERROR_HANDLER"`
	Debug           bool   `yaml:"debug" envconfig:"DEBUG"`
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

type FSBase struct {
	ExportPath          string `yaml:"export_path" envconfig:"EXPORT_PATH"`
	FilePath            string `yaml:"file_path" envconfig:"FILE_PATH"`
	EnableDeduplication bool   `yaml:"enable_deduplication" envconfig:"ENABLE_DEDUPLICATION"`
	ExportLimitOnFolder int    `yaml:"export_limit_on_folder" envconfig:"EXPORT_LIMIT_ON_FOLDER"`
}

func DefaultFSBase() FSBase {
	return FSBase{}
}

type Sqlite struct {
	FilePath string `yaml:"file_path" envconfig:"FILE_PATH"`
}

func DefaultSqlite() Sqlite {
	return Sqlite{}
}

type ZipScanner struct {
	MasterAddr  string `yaml:"master_addr" envconfig:"MASTER_ADDR"`
	MasterToken string `yaml:"master_token" envconfig:"MASTER_TOKEN"`
}

func DefaultZipScanner() ZipScanner {
	return ZipScanner{}
}

type Highway struct {
	PrivateKey    string        `yaml:"private_key" envconfig:"PRIVATE_KEY"`
	TokenLifetime time.Duration `yaml:"token_lifetime" envconfig:"TOKEN_LIFETIME"`
}

func DefaultHighway() Highway {
	return Highway{}
}
