package config

import "time"

type FSBase struct {
	ImportPath          string `yaml:"import_path" envconfig:"IMPORT_PATH"`
	FilePath            string `yaml:"file_path" envconfig:"FILE_PATH"`
	EnableDeduplication bool   `yaml:"enable_deduplication" envconfig:"ENABLE_DEDUPLICATION"`
	ImportLimitOnFolder int    `yaml:"import_limit_on_folder" envconfig:"IMPORT_LIMIT_ON_FOLDER"`
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
