package chat

import (
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	usecase "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/chat"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

type Handler struct{ uc usecase.ChatUsecase }

func NewHandler(uc usecase.ChatUsecase) *Handler { return &Handler{uc: uc} }

func (h *Handler) GetRooms(w http.ResponseWriter, r *http.Request) {
	sessionID := h.getSessionId(r)
	if sessionID == uuid.Nil {
		json.WriteError(w, http.StatusUnauthorized, "unauthorized")
		return
	}
	rooms, _ := h.uc.GetUserRooms(r.Context(), sessionID)
	json.Write(w, http.StatusOK, rooms)
}

func (h *Handler) GetMessages(w http.ResponseWriter, r *http.Request) {
	sessionID := h.getSessionId(r)
	if sessionID == uuid.Nil {
		json.WriteError(w, http.StatusUnauthorized, "unauthorized")
		return
	}
	roomID, _ := uuid.Parse(mux.Vars(r)["room_id"])
	msgs, err := h.uc.GetMessages(r.Context(), sessionID, roomID, 50)
	if err != nil {
		json.WriteError(w, http.StatusForbidden, err.Error())
		return
	}
	json.Write(w, http.StatusOK, msgs)
}

func (h *Handler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	sessionID := h.getSessionId(r)
	if sessionID == uuid.Nil {
		json.WriteError(w, http.StatusUnauthorized, "unauthorized")
		return
	}
	var in struct {
		Members []uuid.UUID `json:"members"`
		Name    *string     `json:"name"`
	}
	if err := json.Read(r, &in); err != nil {
		json.WriteError(w, http.StatusBadRequest, "bad json")
		return
	}
	roomID, err := h.uc.CreateRoom(r.Context(), sessionID, in.Members, in.Name)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.Write(w, http.StatusCreated, map[string]string{"room_id": roomID.String()})
}

func (h *Handler) Connect(w http.ResponseWriter, r *http.Request) {
	cookie, err := cookies.GetCookie(r)
	if err != nil {
		http.Error(w, "no cookie", http.StatusUnauthorized)
		return
	}
	sessionID, err := uuid.Parse(cookie.Value)
	if err != nil {
		http.Error(w, "bad session", http.StatusUnauthorized)
		return
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	h.uc.HandleConnection(r.Context(), conn, sessionID)
}

func (h *Handler) getSessionId(r *http.Request) uuid.UUID {
	cookie, _ := cookies.GetCookie(r)
	if cookie == nil {
		return uuid.Nil
	}
	sessionId, err := uuid.Parse(cookie.Value)
	if err != nil {
		return uuid.Nil
	}

	return sessionId
}
