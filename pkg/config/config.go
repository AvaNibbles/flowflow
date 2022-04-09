package config

import (
	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
)

type Http struct {
	Host string `mapstructure:"host" default:""`
	Port int    `mapstructure:"port" default:"9002"`
}

type Config struct {
	Http    *Http           `mapstructure:"http"`
	Storage *StorageConfig  `mapstructure:"storage"`
	Db      *DatabaseConfig `mapstructure:"db"`
	Client  *ClientConfig   `mapstructure:"client"`
}

type StorageConfig struct {
	Type string           `mapstructure:"type" default:"s3"`
	S3   *S3StorageConfig `mapstructure:"s3"`
}

type S3StorageConfig struct {
	Endpoint       *string `mapstructure:"endpoint"`
	Region         *string `mapstructure:"region"`
	BucketName     string  `mapstructure:"bucketName"`
	DisableSSL     *bool   `mapstructure:"disablessl"`
	ForcePathStyle *bool   `mapstructure:"forcepathstyle"`
}

type DatabaseConfig struct {
	Type     string                  `mapstructure:"type" default:"postgres"`
	Postgres *PostgresDatabaseConfig `mapstructure:"postgres"`
}

type PostgresDatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     uint16 `mapstructure:"port" default:"5432"`
	Username string `mapstructure:"username" default:"postgres"`
	Password string `mapstructure:"password" default:"postgres"`
	Database string `mapstructure:"database" default:"default"`
	SslMode  string `mapstructure:"sslmode" default:"disable"`
}

type ClientConfig struct {
	Host string `mapstructureL:"host"`
}

func NewConfig() *Config {
	ret := &Config{
		Http: &Http{},
		Storage: &StorageConfig{
			Type: "",
			S3:   &S3StorageConfig{},
		},
		Db: &DatabaseConfig{
			Type:     "",
			Postgres: &PostgresDatabaseConfig{},
		},
		Client: &ClientConfig{},
	}

	defaults.SetDefaults(ret.Http)
	defaults.SetDefaults(ret.Storage)
	defaults.SetDefaults(ret.Storage.S3)
	defaults.SetDefaults(ret.Db)
	defaults.SetDefaults(ret.Db.Postgres)

	return ret
}

func LoadConfigFromViper() (*Config, error) {
	ret := NewConfig()

	if err := viper.Unmarshal(ret); err != nil {
		return nil, err
	}

	if err := ret.Validate(); err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *Config) Validate() error {
	return nil
}
