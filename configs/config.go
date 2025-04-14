package configs

type Config struct {
	Database *DatabaseConfig
}

func LoadConfig() *Config {
	return &Config{
		Database: LoadDatabaseConfig(),
	}
}
