package postgres

import (
	"context"
	"fmt"
	"github.com/amir79esmaeili/sms-gateway/internal/cfg"
	"github.com/jackc/pgx/v5"
	"log"
)

const UrlSample = "postgres://%s:%s@%s:%s/%s"

// ConnectToDB takes configuration struct and creates a connection to the given Database
func ConnectToDB(cfg *cfg.Config) (*pgx.Conn, error) {
	dbUrl := fmt.Sprintf(UrlSample, cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}
	err = conn.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	log.Println("[INFO] Database connected...")
	return conn, nil
}

func KillConnection(conn *pgx.Conn) error {
	if !conn.IsClosed() {
		return conn.Close(context.Background())
	}
	return nil
}
