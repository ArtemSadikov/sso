package config

import "time"

type Environment string

const (
	Local Environment = "local"
	Dev   Environment = "development"
	Prod  Environment = "production"
)

type Config struct {
	Env            Environment    `yaml:"env" env-default:"local"`
	GRPC           *GrpcConfig    `yaml:"grpc"`
	Storage        *StorageConfig `yaml:"postgres"`
	MigrationsPath string
	TokenTTL       time.Duration `yaml:"token_ttl" env-default:"1h"`
}

type GrpcConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type StorageConfig struct {
	Port     int    `yaml:"port" env-default:"5432"`
	Host     string `yaml:"host" env-default:"127.0.0.1"`
	User     string `yaml:"user" env-default:"postgres"`
	Password string `yaml:"password" env-default:"password"`
	Database string `yaml:"database" env-default:"postgres"`
}
