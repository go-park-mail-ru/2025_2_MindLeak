package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	repository "github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/chat"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/chat/dto"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type ChatUsecase interface {
	HandleConnection(ctx context.Context, conn *websocket.Conn, sessionID uuid.UUID)
	GetUserRooms(ctx context.Context, sessionID uuid.UUID) ([]models.Room, error)
	GetMessages(ctx context.Context, sessionID, roomID uuid.UUID, limit int) ([]models.Message, error)
	CreateRoom(ctx context.Context, sessionID uuid.UUID, members []uuid.UUID, name *string) (uuid.UUID, error)
}

type chatUsecase struct {
	sessionRepo session.SessionRepository
	realtime    RealtimeUsecase
	repo        repository.ChatRepository
}

func NewChatUsecase(s session.SessionRepository, repo repository.ChatRepository) ChatUsecase {
	uc := &chatUsecase{sessionRepo: s, repo: repo}
	uc.realtime = NewRealtimeUsecase(repo)
	return uc
}

func (uc *chatUsecase) HandleConnection(ctx context.Context, conn *websocket.Conn, sessionID uuid.UUID) {
	sess, err := uc.sessionRepo.GetSessionById(ctx, sessionID)
	if err != nil {
		conn.WriteJSON(map[string]string{"error": "invalid session"})
		conn.Close()
		return
	}

	rooms, _ := uc.repo.GetUserRooms(ctx, sess.UserId)
	roomIDs := make([]uuid.UUID, len(rooms))
	for i, r := range rooms {
		roomIDs[i] = r.ID
	}

	uc.realtime.Register(conn, sess.UserId, roomIDs)

	for _, room := range rooms {
		msgs, _ := uc.repo.GetMessages(ctx, room.ID, 50)
		for i := len(msgs) - 1; i >= 0; i-- {
			conn.WriteJSON(dto.ServerEvent{Type: "message", Data: msgs[i], Ts: time.Now()})
		}
	}
}

func (uc *chatUsecase) GetUserRooms(ctx context.Context, sessionID uuid.UUID) ([]models.Room, error) {
	session, _ := uc.sessionRepo.GetSessionById(ctx, sessionID)
	userID := session.UserId

	return uc.repo.GetUserRooms(ctx, userID)
}

func (uc *chatUsecase) GetMessages(ctx context.Context, sessionID, roomID uuid.UUID, limit int) ([]models.Message, error) {
	session, _ := uc.sessionRepo.GetSessionById(ctx, sessionID)
	userID := session.UserId

	if ok, _ := uc.repo.IsUserInRoom(ctx, roomID, userID); !ok {
		return nil, fmt.Errorf("forbidden")
	}
	return uc.repo.GetMessages(ctx, roomID, limit)
}

func (uc *chatUsecase) CreateRoom(ctx context.Context, sessionID uuid.UUID, members []uuid.UUID, name *string) (uuid.UUID, error) {
	session, _ := uc.sessionRepo.GetSessionById(ctx, sessionID)
	userID := session.UserId

	n := ""
	if name != nil {
		n = *name
	}
	all := append([]uuid.UUID{userID}, members...)
	return uc.repo.CreateRoom(ctx, n, len(all) > 2, all)
}
