package repository

import (
	"context"
	"github.com/amir79esmaeili/sms-gateway/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

const (
	QueryAll = "SELECT * FROM messages ORDER BY date"
)

type MessageRepositoryImpl struct {
	PostgresRepository
}

func NewMessageRepository(db *pgx.Conn) *MessageRepositoryImpl {
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
		err := rows.Scan(&m.Sender, &m.Recipient, &m.Body, &m.Date, &m.Provider)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &m)
	}
	return messages, nil
}

func (m MessageRepositoryImpl) Get(id uuid.UUID, ctx context.Context) (*model.Message, error) {
	return nil, nil
}

func (m MessageRepositoryImpl) Create(entity *model.Message, ctx context.Context) error {
	return nil
}
