package handler

import (
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/google/uuid"
)

type Article struct {
	Id        string    `json:"id"`
	AuthorId  string    `json:"author_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func FeedHandler(w http.ResponseWriter, r *http.Request, sessions *repository.InMemorySession, articles *repository.InMemoryArticle) {
	session, err := sessions.CreateSession()
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	cookies.SetCookie(w, session.SessionId)

	// моки постов, надо заполнить нормальными данными
	mockArticles := []Article{
		{
			Id:        uuid.New().String(),
			AuthorId:  uuid.New().String(),
			Title:     "Первый пост",
			Content:   "Это пример поста для ленты.",
			CreatedAt: time.Now(),
		},
	}

	err = json.Write(w, http.StatusOK, mockArticles)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

}
