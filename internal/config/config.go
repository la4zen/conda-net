package config

type Http struct {
	IP   string `json:"ip,omitempty"`
	Port string `json:"port,omitempty"`
}

type Database struct {
	IP       string `json:"ip,omitempty"`
	Name     string `json:"name,omitempty"`
	Port     string `json:"port,omitempty"`
	User     string `json:"user,omitempty"`
	Password string
}

type Config struct {
	Http     Http
	Database Database
}

func New() *Config {
	return &Config{
		Http: Http{
			IP:   "0.0.0.0",
			Port: "8000",
		},
		Database: Database{
			IP:       "127.0.0.1",
			Port:     "5432",
			Password: "postgres",
			User:     "postgres",
			Name:     "postgres",
		},
	}
}
