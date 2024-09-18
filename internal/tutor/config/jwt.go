package config

type JWTConfig struct {
	SecretKey string
	Exp       uint
}
