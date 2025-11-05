package profile

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/json"
	"github.com/google/uuid"
	"net/http"
)

func (h *Handler) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cookie, err := cookies.GetCookie(r)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	sessionID, err := uuid.Parse(cookie.Value)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	session, err := h.Usecase.GetSession(ctx, sessionID)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, "file not provided")
		return
	}
	defer file.Close()

	user, err := h.Usecase.UploadAvatar(ctx, session.UserId, file, header)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.Write(w, http.StatusOK, user)
}

func (h *Handler) DeleteAvatar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cookie, err := cookies.GetCookie(r)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	sessionID, err := uuid.Parse(cookie.Value)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	session, err := h.Usecase.GetSession(ctx, sessionID)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	user, err := h.Usecase.DeleteAvatar(ctx, session.UserId)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.Write(w, http.StatusOK, user)
}

func (h *Handler) UploadCover(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cookie, err := cookies.GetCookie(r)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	sessionID, err := uuid.Parse(cookie.Value)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	session, err := h.Usecase.GetSession(ctx, sessionID)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, "file not provided")
		return
	}
	defer file.Close()

	profile, err := h.Usecase.UploadCover(ctx, session.UserId, file, header)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.Write(w, http.StatusOK, profile)
}

func (h *Handler) DeleteCover(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cookie, err := cookies.GetCookie(r)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	sessionID, err := uuid.Parse(cookie.Value)
	if err != nil {
		code, msg := h.handleError(err)
		json.WriteError(w, code, msg)
		return
	}

	profile, err := h.Usecase.DeleteCover(ctx, sessionID)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.Write(w, http.StatusOK, profile)
}
