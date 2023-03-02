package db

import (
	"database/sql"
	"fmt"

	"github.com/OlegDjur/Go-Auth-Microservice/config"
	_ "github.com/lib/pq"
)

func NewPsqlDB(c *config.Config) (*sql.DB, error) {
	dataSource := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.Postgres.PostgresHost,
		c.Postgres.PostgresPort,
		c.Postgres.PostgresUser,
		c.Postgres.PostgresDBName,
		c.Postgres.PostgresPassword,
	)

	db, err := sql.Open(c.Postgres.PgDriver, dataSource)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
