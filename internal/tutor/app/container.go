package app

import (
	"github.com/jmoiron/sqlx"
)

type Container struct {
	pgConnect *sqlx.DB
}
