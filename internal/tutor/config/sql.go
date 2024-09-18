package config

type SQLConfig struct {
	DatabaseName string
	Username     string
	Password     string
	SSLMode      bool
}

func (sqlConf *SQLConfig) GetStringSSLMode() string {
	if sqlConf.SSLMode {
		return "enable"
	}

	return "disable"
}
