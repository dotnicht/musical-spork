package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"

	"example.com/modmonolith/internal/modules/users/application/commands"
	"example.com/modmonolith/internal/modules/users/application/queries"
	"example.com/modmonolith/internal/modules/users/application/service"
	"example.com/modmonolith/internal/modules/users/domain"
)

type Handler struct{ h service.Handlers }

func New(h service.Handlers) *Handler { return &Handler{h: h} }

func (hd *Handler) Routes() http.Handler {
	r := chi.NewRouter()

	r.Route("/v1/users", func(r chi.Router) {
		r.Post("/", hd.createUser)
		r.Get("/", hd.listUsers)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", hd.getUser)
			r.Patch("/", hd.updateUser)
			r.Delete("/", hd.deleteUser)
		})
	})
	return r
}

type createUserReq struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
type createUserResp struct { ID string `json:"id"` }

func (hd *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var req createUserReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErr(w, http.StatusBadRequest, "invalid json")
		return
	}
	id, err := hd.h.Create.Handle(r.Context(), commands.CreateUserCommand{Email: req.Email, Name: req.Name})
	if err != nil {
		writeDomainErr(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, createUserResp{ID: id})
}

func (hd *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	dto, err := hd.h.Get.Handle(r.Context(), queries.GetUserQuery{ID: id})
	if err != nil {
		writeDomainErr(w, err)
		return
	}
	writeJSON(w, http.StatusOK, toUserJSON(dto))
}

func (hd *Handler) listUsers(w http.ResponseWriter, r *http.Request) {
	limit := int32(parseIntDefault(r.URL.Query().Get("limit"), 50))
	offset := int32(parseIntDefault(r.URL.Query().Get("offset"), 0))
	dtos, err := hd.h.List.Handle(r.Context(), queries.ListUsersQuery{Limit: limit, Offset: offset})
	if err != nil {
		writeDomainErr(w, err)
		return
	}
	out := make([]userJSON, 0, len(dtos))
	for _, d := range dtos {
		out = append(out, toUserJSON(d))
	}
	writeJSON(w, http.StatusOK, map[string]any{"users": out})
}

type updateUserReq struct {
	Email *string `json:"email"`
	Name  *string `json:"name"`
}

func (hd *Handler) updateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var req updateUserReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErr(w, http.StatusBadRequest, "invalid json")
		return
	}
	err := hd.h.Update.Handle(r.Context(), commands.UpdateUserCommand{ID: id, Email: req.Email, Name: req.Name})
	if err != nil {
		writeDomainErr(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (hd *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := hd.h.Delete.Handle(r.Context(), commands.DeleteUserCommand{ID: id}); err != nil {
		writeDomainErr(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

type userJSON struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func toUserJSON(d *queries.UserDTO) userJSON {
	return userJSON{
		ID: d.ID,
		Email: d.Email,
		Name: d.Name,
		CreatedAt: d.CreatedAt.Format(time.RFC3339),
		UpdatedAt: d.UpdatedAt.Format(time.RFC3339),
	}
}

type errResp struct{ Error string `json:"error"` }

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeErr(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, errResp{Error: msg})
}

func writeDomainErr(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrInvalidEmail), errors.Is(err, domain.ErrInvalidUserName):
		writeErr(w, http.StatusBadRequest, err.Error())
	case errors.Is(err, domain.ErrEmailAlreadyUsed):
		writeErr(w, http.StatusConflict, err.Error())
	case errors.Is(err, domain.ErrUserNotFound):
		writeErr(w, http.StatusNotFound, err.Error())
	default:
		writeErr(w, http.StatusInternalServerError, "internal error")
	}
}

func parseIntDefault(s string, def int) int {
	if s == "" { return def }
	n, err := strconv.Atoi(s)
	if err != nil { return def }
	return n
}
