package app

import (
	"fmt"

	"github.com/spf13/viper"
)

type JWTConfig struct {
	Exp       uint
	SecretKey string
}

type SQLConfig struct {
	DatabaseName string
	Username     string
	Password     string
	SSLMode      bool
}

type HTTPConfig struct {
	Host string
	Port uint
}

func (sqlConf *SQLConfig) GetStringSSLMode() string {
	if sqlConf.SSLMode {
		return "enable"
	}

	return "disable"
}

type Config struct {
	JwtConfig  JWTConfig
	SQLConfig  SQLConfig
	HTTPConfig HTTPConfig
}

func newConfig() (Config, error) {
	viper.SetDefault("jwt.token-exp", 3600)
	viper.SetDefault("db.ssl-mode", false)
	viper.SetDefault("http.port", 3000)
	viper.SetDefault("http.host", "")

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
			Port: viper.GetUint("http.port"),
			Host: viper.GetString("http.host"),
		},
	}, nil
}
