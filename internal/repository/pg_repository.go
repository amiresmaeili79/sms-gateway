package repository

import (
	"github.com/jackc/pgx/v5"
)

type PostgresRepository struct {
	db *pgx.Conn
}
