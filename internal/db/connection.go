package db

import (
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Configuration struct {
	User     string
	Password string
	Address  string
	Port     string
	Database string
}

func NewDB(config Configuration) *bun.DB {
	return bun.NewDB(
		sql.OpenDB(
			pgdriver.NewConnector(
				pgdriver.WithInsecure(true),
				pgdriver.WithAddr(fmt.Sprintf("%s:%s", config.Address, config.Port)),
				pgdriver.WithUser(config.User),
				pgdriver.WithPassword(config.Password),
				pgdriver.WithDatabase(config.Database),
			),
		),
		pgdialect.New(),
	)
}
