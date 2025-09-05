package config

import "time"

type FSBase struct {
	ImportPath          string `toml:"import_path" yaml:"import_path" envconfig:"IMPORT_PATH"`
	FilePath            string `toml:"file_path" yaml:"file_path" envconfig:"FILE_PATH"`
	EnableDeduplication bool   `toml:"enable_deduplication" yaml:"enable_deduplication" envconfig:"ENABLE_DEDUPLICATION"`
	ImportLimitOnFolder int    `toml:"import_limit_on_folder" yaml:"import_limit_on_folder" envconfig:"IMPORT_LIMIT_ON_FOLDER"`
}

func DefaultFSBase() FSBase {
	return FSBase{}
}

type Sqlite struct {
	FilePath string `toml:"file_path" yaml:"file_path" envconfig:"FILE_PATH"`
}

func DefaultSqlite() Sqlite {
	return Sqlite{}
}

type ZipScanner struct {
	MasterAddr  string `toml:"master_addr" yaml:"master_addr" envconfig:"MASTER_ADDR"`
	MasterToken string `toml:"master_token" yaml:"master_token" envconfig:"MASTER_TOKEN"`
}

func DefaultZipScanner() ZipScanner {
	return ZipScanner{}
}

type Highway struct {
	PrivateKey    string        `toml:"private_key" yaml:"private_key" envconfig:"PRIVATE_KEY"`
	TokenLifetime time.Duration `toml:"token_lifetime" yaml:"token_lifetime" envconfig:"TOKEN_LIFETIME"`
}

func DefaultHighway() Highway {
	return Highway{}
}
