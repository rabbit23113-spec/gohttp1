package cfg

import "fmt"

type Config struct {
	DbName     string `env:"POSTGRES_DB"`
	DbUser     string `env:"POSTGRES_USER"`
	DbPassword string `env:"POSTGRES_PASSWORD"`
	DbPort     string `env:"POSTGRES_PORT"`
	DbHost     string `env:"POSTGRES_HOST"`
}

func (c *Config) DatabaseUrl() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		c.DbUser,
		c.DbPassword,
		c.DbHost,
		c.DbPort,
		c.DbName,
	)
}
