package userdelivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	userusecase "github.com/lunatictiol/clean-architecture-in-go/internal/services/user/userUsecase"
)

type Handler struct {
	usecase *userusecase.UserUsecase
}

func NewHandler(u *userusecase.UserUsecase) *Handler {
	return &Handler{usecase: u}
}

func (h *Handler) RegisterUserRoutes(r chi.Router) {
	r.Route("/user", func(r chi.Router) {
		r.Get("/{userID}", h.GetUserByID)
	})
}

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "userID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	u, err := h.usecase.GetUserByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(u)
}
