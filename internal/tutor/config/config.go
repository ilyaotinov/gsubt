package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	defaultTimeoutOnStop = 5
	defaultJWTExp        = 3600
	defaultHTTPPort      = 3000
)

type Config struct {
	JwtConfig  JWTConfig
	SQLConfig  SQLConfig
	HTTPConfig HTTPConfig
}

func New() (Config, error) {
	viper.SetDefault("jwt.token-exp", defaultJWTExp)
	viper.SetDefault("db.ssl-mode", false)
	viper.SetDefault("http.port", defaultHTTPPort)
	viper.SetDefault("http.host", "")
	viper.SetDefault("http.timeout-on-stop", defaultTimeoutOnStop)

	viper.SetConfigFile(viper.GetString("config-file-path"))
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("invalid config file provided: %w", err)
	}

	return Config{
		JwtConfig: JWTConfig{
			SecretKey: viper.GetString("jwt.secret-key"),
			Exp:       viper.GetUint("jwt.token-exp"),
		},
		SQLConfig: SQLConfig{
			DatabaseName: viper.GetString("db.name"),
			Username:     viper.GetString("db.username"),
			Password:     viper.GetString("db.password"),
			SSLMode:      viper.GetBool("db.ssl-mode"),
		},
		HTTPConfig: HTTPConfig{
			Port:          viper.GetUint("http.port"),
			Host:          viper.GetString("http.host"),
			TimeoutOnStop: viper.GetInt("http.timeout-on-stop"),
		},
	}, nil
}
