package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect string
	Username string
	Password string
	Name string
	Charset string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect: "mysql",
			Username: "admin",
			Password: "tonjootonjoo",
			Name:"go_rest",
			Charset:"utf8",
		},
	}
}