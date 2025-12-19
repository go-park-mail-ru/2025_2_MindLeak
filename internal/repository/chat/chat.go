package repository

import (
	"context"
	"database/sql"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/google/uuid"
)

type ChatRepository interface {
	GetUserRooms(ctx context.Context, userID uuid.UUID) ([]models.Room, error)
	GetMessages(ctx context.Context, roomID uuid.UUID, limit int) ([]models.Message, error)
	CreateMessage(ctx context.Context, msg *models.Message) error
	IsUserInRoom(ctx context.Context, roomID, userID uuid.UUID) (bool, error)
	CreateRoom(ctx context.Context, name string, isGroup bool, members []uuid.UUID) (uuid.UUID, error)
}

type chatRepo struct{ db *sql.DB }

func NewChatRepository(db *sql.DB) ChatRepository { return &chatRepo{db} }

func (r *chatRepo) GetUserRooms(ctx context.Context, userID uuid.UUID) ([]models.Room, error) {
	q := `SELECT r.room_id, COALESCE(r.name,''), r.is_group, r.created_at
	      FROM chat_rooms r
	      JOIN chat_room_members m ON r.room_id = m.room_id
	      WHERE m.user_id = $1 ORDER BY r.created_at DESC`
	rows, err := r.db.QueryContext(ctx, q, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var rooms []models.Room
	for rows.Next() {
		var r models.Room
		if err := rows.Scan(&r.ID, &r.Name, &r.IsGroup, &r.CreatedAt); err != nil {
			return nil, err
		}
		rooms = append(rooms, r)
	}
	return rooms, rows.Err()
}

func (r *chatRepo) GetMessages(ctx context.Context, roomID uuid.UUID, limit int) ([]models.Message, error) {
	q := `SELECT m.message_id, m.room_id, m.user_id, u.name, COALESCE(u.avatar,''), m.text, m.created_at
	      FROM chat_messages m
	      JOIN "user" u ON m.user_id = u.user_id
	      WHERE m.room_id = $1 ORDER BY m.created_at DESC LIMIT $2`
	rows, err := r.db.QueryContext(ctx, q, roomID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var msgs []models.Message
	for rows.Next() {
		var m models.Message
		if err := rows.Scan(&m.ID, &m.RoomID, &m.UserID, &m.UserName, &m.Avatar, &m.Text, &m.CreatedAt); err != nil {
			return nil, err
		}
		msgs = append(msgs, m)
	}
	return msgs, rows.Err()
}

func (r *chatRepo) CreateMessage(ctx context.Context, msg *models.Message) error {
	q := `INSERT INTO chat_messages (room_id, user_id, text, created_at) VALUES ($1,$2,$3,$4) RETURNING message_id`
	return r.db.QueryRowContext(ctx, q, msg.RoomID, msg.UserID, msg.Text, msg.CreatedAt).Scan(&msg.ID)
}

func (r *chatRepo) IsUserInRoom(ctx context.Context, roomID, userID uuid.UUID) (bool, error) {
	var ok bool
	err := r.db.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM chat_room_members WHERE room_id=$1 AND user_id=$2)`, roomID, userID).Scan(&ok)
	return ok, err
}

func (r *chatRepo) CreateRoom(ctx context.Context, name string, isGroup bool, members []uuid.UUID) (uuid.UUID, error) {
	var roomID uuid.UUID
	err := r.db.QueryRowContext(ctx, `INSERT INTO chat_rooms (name, is_group) VALUES ($1,$2) RETURNING room_id`, name, isGroup).Scan(&roomID)
	if err != nil {
		return uuid.Nil, err
	}
	for _, uid := range members {
		_, err := r.db.ExecContext(ctx, `INSERT INTO chat_room_members (room_id, user_id) VALUES ($1,$2)`, roomID, uid)
		if err != nil {
			return uuid.Nil, err
		}
	}
	return roomID, nil
}
