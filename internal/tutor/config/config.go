package config

import "github.com/spf13/viper"

type JWTConfig struct {
	SecretKey string
}

type Config struct {
	JwtConfig JWTConfig
}

func New() Config {
	return Config{JwtConfig: JWTConfig{SecretKey: viper.GetString("secret-key")}}
}
