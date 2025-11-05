package comment

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
	"net/http"
)

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
			logger.Error(ctx, "parse articleId: %v", err)
			return
		}
		comments, err := h.Usecase.GetCommentsByArticle(ctx, articleId)
		if err != nil {
			code, msg := h.handleError(err)
			json.WriteError(w, code, msg)
			logger.Error(ctx, "get comments by article: %v", err)
			return
		}

		outputDto := h.mapSliceToOutputDto(comments)

		if err = json.Write(w, http.StatusOK, outputDto); err != nil {
			logger.Error(ctx, "write json (article): %v", err)
			return
		}
	} else if authorIdString != "" {
		authorId, err := uuid.Parse(authorIdString)
		if err != nil {
			code, msg := h.handleError(err)
			json.WriteError(w, code, msg)
			logger.Error(ctx, "parse authorId: %v", err)
			return
		}
		comments, err := h.Usecase.GetCommentsByArticle(ctx, authorId)
		if err != nil {
			code, msg := h.handleError(err)
			json.WriteError(w, code, msg)
			logger.Error(ctx, "get comments by article: %v", err)
			return
		}

		outputDto := h.mapSliceToOutputDto(comments)

		if err = json.Write(w, http.StatusOK, outputDto); err != nil {
			logger.Error(ctx, "write json (author): %v", err)
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
			logger.Error(ctx, "get session: %v", err)
			return
		}

		authorId := session.UserId

		comments, err := h.Usecase.GetCommentsByArticle(ctx, authorId)
		if err != nil {
			code, msg := h.handleError(err)
			json.WriteError(w, code, msg)
			logger.Error(ctx, "get comments by session author: %v", err)
			return
		}

		outputDto := h.mapSliceToOutputDto(comments)

		if err = json.Write(w, http.StatusOK, outputDto); err != nil {
			logger.Error(ctx, "write json (session): %v", err)
			return
		}
	}

}
