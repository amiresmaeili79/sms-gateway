package repository

import (
	"context"
	"github.com/amir79esmaeili/sms-gateway/internal/model"
	"github.com/jackc/pgx/v5"
)

const (
	QueryAll = "SELECT * FROM messages ORDER BY date"
)

type MessageRepository struct {
	PostgresRepository
}

func NewMessageRepository(db *pgx.Conn) *MessageRepository {
	return &MessageRepository{
		PostgresRepository{
			db: db,
		},
	}
}

func (m MessageRepository) List(ctx context.Context) (interface{}, error) {
	rows, err := m.db.Query(ctx, QueryAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*model.Message

	for rows.Next() {
		var m model.Message
		err := rows.Scan(&m.Sender, &m.Recipient, &m.Body, &m.Date)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &m)
	}
	return messages, nil
}

func (m MessageRepository) Get(id interface{}, ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (m MessageRepository) Create(entity interface{}, ctx context.Context) (interface{}, error) {
	return nil, nil
}
