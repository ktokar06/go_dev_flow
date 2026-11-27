package config

type Config struct {
	Port        string
	DatabaseURL string
	JWTSecret   string
	CORSOrigin  string
	LogLevel    string
	LogFile     string
	Environment string
	ConfigFile  string
	LogJSON     bool
}