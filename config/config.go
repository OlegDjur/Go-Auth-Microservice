package config

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
}

type ServerConfig struct{}

type PostgresConfig struct{}
