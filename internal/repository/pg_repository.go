package repository

import (
	"github.com/amir79esmaeili/sms-gateway/internal/cfg"
	"github.com/amir79esmaeili/sms-gateway/internal/postgres"
	"github.com/jackc/pgx/v5"
	"log"
)

type PostgresRepository struct {
	db *pgx.Conn
}

func (p *PostgresRepository) Configure(config *cfg.Config) {
	var err error
	p.db, err = postgres.ConnectToDB(config)
	log.Fatalf("could not configure db, %v \n", err)
}
