package comment

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
	"net/http"
)

// GetCommentsHandler обрабатывает HTTP-запрос для получения комментариев.
//
// Поддерживаются параметры запроса:
//   - authorId: UUID автора, чьи комментарии нужно получить.
//   - articleId: UUID статьи, для которой нужно получить комментарии.
//
// Если параметры не указаны, то возвращается список комментариев пользователя текущей сессии.
// Если указаны оба параметра, то возвращается код HTTP 501.
//
// Возвращает JSON-массив объектов CommentIODto со статусом 200.
// В случае ошибок - HTTP 500.
func (h *Handler) GetCommentsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	authorIdString := r.URL.Query().Get("authorId")

	articleIdString := r.URL.Query().Get("articleId")

	if articleIdString != "" && authorIdString != "" {
		json.WriteError(w, http.StatusNotImplemented, "not implemented")
		logger.Error(ctx, "not implemented")
		return
	}

	if articleIdString != "" {
		articleId, err := uuid.Parse(articleIdString)
		if err != nil {
			code, msg := h.handleError(err)
			json.WriteError(w, code, msg)
			logger.Error(ctx, err.Error())
			return
		}
		comments, err := h.Usecase.GetCommentsByArticle(ctx, articleId)
		if err != nil {
			code, msg := h.handleError(err)
			json.WriteError(w, code, msg)
			logger.Error(ctx, err.Error())
			return
		}

		outputDto := h.mapSliceToOutputDto(comments)

		if err = json.Write(w, http.StatusOK, outputDto); err != nil {
			logger.Error(ctx, err.Error())
			return
		}
	} else if authorIdString != "" {
		authorId, err := uuid.Parse(authorIdString)
		if err != nil {
			code, msg := h.handleError(err)
			json.WriteError(w, code, msg)
			logger.Error(ctx, err.Error())
			return
		}
		comments, err := h.Usecase.GetCommentsByArticle(ctx, authorId)
		if err != nil {
			code, msg := h.handleError(err)
			json.WriteError(w, code, msg)
			logger.Error(ctx, err.Error())
			return
		}

		outputDto := h.mapSliceToOutputDto(comments)

		if err = json.Write(w, http.StatusOK, outputDto); err != nil {
			logger.Error(ctx, err.Error())
			return
		}
	} else {
		cookie, err := cookies.GetCookie(r)
		if err != nil {
			logger.Error(ctx, "GetCookie: %v", err)
			code, msg := h.handleError(err)
			json.WriteError(w, code, msg)
			return
		}

		sessionId, err := uuid.Parse(cookie.Value)
		if err != nil {
			logger.Error(ctx, "Parse: %v", err)
			code, msg := h.handleError(err)
			json.WriteError(w, code, msg)
			return
		}

		session, err := h.Usecase.GetSession(ctx, sessionId)
		if err != nil {
			code, msg := h.handleError(err)
			json.WriteError(w, code, msg)
			logger.Error(ctx, err.Error())
			return
		}

		authorId := session.UserId

		comments, err := h.Usecase.GetCommentsByArticle(ctx, authorId)
		if err != nil {
			code, msg := h.handleError(err)
			json.WriteError(w, code, msg)
			logger.Error(ctx, err.Error())
			return
		}

		outputDto := h.mapSliceToOutputDto(comments)

		if err = json.Write(w, http.StatusOK, outputDto); err != nil {
			logger.Error(ctx, err.Error())
			return
		}
	}

}
