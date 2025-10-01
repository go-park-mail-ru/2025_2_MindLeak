package feed

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/article"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type ArticleResponse struct {
	Title        string `json:"title"`
	Content      string `json:"content"`
	Image        string `json:"image"`
	AuthorName   string `json:"author_name"`
	AuthorAvatar string `json:"author_avatar"`
}

type UserResponse struct {
	Email  string `json:"email"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func TestFeedHandlerStatus(t *testing.T) {
	type test struct {
		name          string
		method        string
		setupSession  func() (*session.InMemorySession, uuid.UUID)
		setupArticles func() *article.InMemoryArticle
		setCookie     bool
		cookieValue   string
		wantStatus    int
	}

	tests := []test{
		{
			name:          "invalid method",
			method:        http.MethodPost,
			setupSession:  func() (*session.InMemorySession, uuid.UUID) { return session.NewInMemorySession(), uuid.New() },
			setupArticles: article.NewInMemoryArticle,
			setCookie:     false,
			wantStatus:    http.StatusMethodNotAllowed,
		},
		{
			name:          "no cookie",
			method:        http.MethodGet,
			setupSession:  func() (*session.InMemorySession, uuid.UUID) { return session.NewInMemorySession(), uuid.New() },
			setupArticles: article.NewInMemoryArticle,
			setCookie:     false,
			wantStatus:    http.StatusOK,
		},
		{
			name:   "invalid cookie",
			method: http.MethodGet,
			setupSession: func() (*session.InMemorySession, uuid.UUID) {
				return session.NewInMemorySession(), uuid.New()
			},
			setupArticles: article.NewInMemoryArticle,
			setCookie:     true,
			cookieValue:   "invalid-uuid",
			wantStatus:    http.StatusOK,
		},
		{
			name:   "valid session",
			method: http.MethodGet,
			setupSession: func() (*session.InMemorySession, uuid.UUID) {
				sessions := session.NewInMemorySession()
				sessionID := uuid.New()
				_, _ = sessions.CreateSession()
				return sessions, sessionID
			},
			setupArticles: article.NewInMemoryArticle,
			setCookie:     true,
			cookieValue:   "",
			wantStatus:    http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/feed", nil)
			w := httptest.NewRecorder()

			sessions, sessionID := tt.setupSession()
			articles := tt.setupArticles()

			if tt.setCookie {
				cookieValue := tt.cookieValue
				if tt.cookieValue == "" {
					cookieValue = sessionID.String()
				}
				req.AddCookie(&http.Cookie{
					Name:  cookies.SessionID,
					Value: cookieValue,
				})
			}

			FeedHandler(w, req, sessions, articles)

			resp := w.Result()
			defer resp.Body.Close()

			assert.Equal(t, tt.wantStatus, resp.StatusCode, "status code mismatch in case %s", tt.name)
		})
	}
}

func TestFeedHandlerResponse(t *testing.T) {
	type test struct {
		name          string
		method        string
		setupSession  func() (*session.InMemorySession, uuid.UUID)
		setupArticles func() *article.InMemoryArticle
		setCookie     bool
		cookieValue   string
		wantStatus    int
		wantError     bool
		wantErrorText string
		checkCookie   bool
		wantArticles  bool
	}

	tests := []test{
		{
			name:          "invalid method",
			method:        http.MethodPost,
			setupSession:  func() (*session.InMemorySession, uuid.UUID) { return session.NewInMemorySession(), uuid.New() },
			setupArticles: article.NewInMemoryArticle,
			setCookie:     false,
			wantStatus:    http.StatusMethodNotAllowed,
			wantError:     true,
			wantErrorText: "method not allowed",
		},
		{
			name:   "no cookie creates session",
			method: http.MethodGet,
			setupSession: func() (*session.InMemorySession, uuid.UUID) {
				return session.NewInMemorySession(), uuid.New()
			},
			setupArticles: func() *article.InMemoryArticle {
				articles := article.NewInMemoryArticle()
				return articles
			},
			setCookie:    false,
			wantStatus:   http.StatusOK,
			checkCookie:  true,
			wantArticles: true,
		},
		{
			name:   "invalid cookie creates new session",
			method: http.MethodGet,
			setupSession: func() (*session.InMemorySession, uuid.UUID) {
				return session.NewInMemorySession(), uuid.New()
			},
			setupArticles: func() *article.InMemoryArticle {
				articles := article.NewInMemoryArticle()
				return articles
			},
			setCookie:    true,
			cookieValue:  "invalid-uuid",
			wantStatus:   http.StatusOK,
			checkCookie:  true,
			wantArticles: true,
		},
		{
			name:   "valid session returns articles",
			method: http.MethodGet,
			setupSession: func() (*session.InMemorySession, uuid.UUID) {
				sessions := session.NewInMemorySession()
				sessionID := uuid.New()
				_, _ = sessions.CreateSession()
				return sessions, sessionID
			},
			setupArticles: func() *article.InMemoryArticle {
				articles := article.NewInMemoryArticle()
				return articles
			},
			setCookie:    true,
			cookieValue:  "",
			wantStatus:   http.StatusOK,
			wantArticles: true,
			checkCookie:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/feed", nil)
			w := httptest.NewRecorder()

			sessions, sessionID := tt.setupSession()
			articles := tt.setupArticles()

			if tt.setCookie {
				cookieValue := tt.cookieValue
				if tt.cookieValue == "" {
					cookieValue = sessionID.String()
				}
				req.AddCookie(&http.Cookie{
					Name:  cookies.SessionID,
					Value: cookieValue,
				})
			}

			FeedHandler(w, req, sessions, articles)

			resp := w.Result()
			defer resp.Body.Close()

			assert.Equal(t, tt.wantStatus, resp.StatusCode, "status code mismatch")

			data, _ := io.ReadAll(resp.Body)

			if tt.wantError {
				var errResp ErrorResponse
				assert.NoError(t, json.Unmarshal(data, &errResp))
				assert.Equal(t, tt.wantErrorText, errResp.Error, "error message mismatch")
			} else if tt.wantArticles {
				var articlesResp []ArticleResponse
				assert.NoError(t, json.Unmarshal(data, &articlesResp))
				assert.NotEmpty(t, articlesResp, "articles should be returned")
				assert.Equal(t, "ИИ в 2025: Как нейросети меняют бизнес-процессы", articlesResp[0].Title, "article title mismatch")
				assert.Equal(t, "Искусственный интеллект в 2025 году стал неотъемлемой частью бизнеса...", articlesResp[0].Content, "article content mismatch")
				assert.Equal(t, "Алексей Владимиров", articlesResp[0].AuthorName, "author name mismatch")
				assert.Equal(t, "https://st4.depositphotos.com/36740986/38337/i/450/depositphotos_383375990-stock-photo-collection-hundred-dollar-banknotes-female.jpg", articlesResp[0].Image, "image mismatch")
			}

			if tt.checkCookie {
				cookies := resp.Cookies()
				assert.NotEmpty(t, cookies, "cookie must be set")
				assert.Equal(t, "session_id", cookies[0].Name, "cookie name mismatch")
				_, err := uuid.Parse(cookies[0].Value)
				assert.NoError(t, err, "cookie value must be valid UUID")
			}
		})
	}
}
