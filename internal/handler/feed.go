package handler

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	"net/http"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/google/uuid"
)

func FeedHandler(w http.ResponseWriter, r *http.Request, sessions *session.InMemorySession, articles *repository.InMemoryArticle) {
	cookie, err := cookies.GetCookie(r)
	if err == nil {
		if sessionID, parseErr := uuid.Parse(cookie.Value); parseErr == nil {
			if _, sessErr := sessions.GetSessionById(sessionID); sessErr == nil {
				returnFeed(w, articles)
				return
			}
		}
	}

	session, err := sessions.CreateSession()
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	cookies.SetCookie(w, session.SessionId)

	returnFeed(w, articles)
}

func returnFeed(w http.ResponseWriter, articles *repository.InMemoryArticle) {
	mockArticles, err := articles.GetAllArticles()
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(mockArticles) == 0 {
		authorId := uuid.New()

		_, _ = articles.CreateArticle(authorId,
			"ИИ в 2025: Как нейросети меняют бизнес-процессы",
			"Искусственный интеллект в 2025 году стал неотъемлемой частью бизнеса...")

		_, _ = articles.CreateArticle(authorId,
			"Как российский стартап привлёк $10M на рынке SaaS",
			"Российский стартап CloudPeak разработал SaaS-платформу...")

		_, _ = articles.CreateArticle(authorId,
			"Тренды контент-маркетинга: Что работает в 2025 году",
			"Контент-маркетинг в 2025 году переживает новый виток...")

		_, _ = articles.CreateArticle(authorId,
			"Почему 80% стартапов терпят неудачу в первый год",
			"Запуск стартапа — это всегда риск...")

		_, _ = articles.CreateArticle(authorId,
			"Как мы увеличили конверсию на 30% с помощью UX",
			"Компания BrightPath переработала интерфейс...")

		_, _ = articles.CreateArticle(authorId,
			"Экспериментальный сверхдлинный заголовок статьи, в котором мы попробуем уместить сразу и суть, и интригу, и даже немного юмора, чтобы проверить, как фронтенд справится с рендерингом текста...",
			`Это тестовое содержимое статьи, которое специально сделано очень длинным, чтобы проверить работу фронтенда с большими объёмами текста... (длинный текст)`)

		mockArticles, _ = articles.GetAllArticles()
	}

	if err := json.Write(w, http.StatusOK, mockArticles); err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error())
	}
}
