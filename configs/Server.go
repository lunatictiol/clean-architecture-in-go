package configs

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	userdelivery "github.com/lunatictiol/clean-architecture-in-go/internal/services/user/userDelivery"
	userrepository "github.com/lunatictiol/clean-architecture-in-go/internal/services/user/userRepository"
	userusecase "github.com/lunatictiol/clean-architecture-in-go/internal/services/user/userUsecase"
)

type Server struct {
	port string
	db   *sql.DB
}

func (s *Server) New(port string, db *sql.DB) {
	s.port = port
	s.db = db
}

func (s *Server) Run() error {
	r := chi.NewRouter()

	userRepo := userrepository.NewUserRepository(s.db)
	userUC := userusecase.NewUserUsecase(userRepo)
	userHandler := userdelivery.NewHandler(userUC)

	userHandler.RegisterUserRoutes(r)
	return http.ListenAndServe(s.port, r)
}
