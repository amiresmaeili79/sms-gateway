package repository

import (
	"context"
	"github.com/amir79esmaeili/sms-gateway/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log"
)

const (
	QueryAll    = "SELECT * FROM messages ORDER BY date"
	QueryById   = "SELECT * FROM messages WHERE id = $1"
	CreateTable = `CREATE TABLE IF NOT EXISTS messages (
    	id UUID primary key,
    	sender varchar(20),
    	recipient varchar(20) not null,
    	body varchar(100) not null,
    	date date not null,
    	provider varchar(20)
    );`
)

type MessageRepositoryImpl struct {
	PostgresRepository
}

func NewMessageRepository(db *pgx.Conn) *MessageRepositoryImpl {
	_, err := db.Exec(context.Background(), CreateTable)
	if err != nil {
		log.Fatalf("Could not prepare table creation query, %v\n", err)
		return nil
	}
	return &MessageRepositoryImpl{
		PostgresRepository{
			db: db,
		},
	}
}

func (m MessageRepositoryImpl) List(ctx context.Context) ([]*model.Message, error) {
	rows, err := m.db.Query(ctx, QueryAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*model.Message

	for rows.Next() {
		var m model.Message
		err := rows.Scan(&m.Id, &m.Sender, &m.Recipient, &m.Body, &m.Date, &m.Provider)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &m)
	}
	return messages, nil
}

func (m MessageRepositoryImpl) Get(id uuid.UUID, ctx context.Context) (*model.Message, error) {
	row := m.db.QueryRow(ctx, QueryById, id)
	var msg *model.Message
	err := row.Scan(&msg.Id, &msg.Sender, &msg.Recipient, &msg.Body, &msg.Date, &msg.Provider)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (m MessageRepositoryImpl) Create(entity *model.Message, ctx context.Context) error {
	return nil
}
