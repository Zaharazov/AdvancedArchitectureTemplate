package http

import (
	"database/sql"

	"github.com/Zaharazov/AdvancedArchitectureTemplate/internal/repository/postgres"
	"github.com/Zaharazov/AdvancedArchitectureTemplate/internal/service"
	"github.com/Zaharazov/AdvancedArchitectureTemplate/internal/transport/http/handler"

	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	userRepo := postgres.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router := mux.NewRouter()
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")

	return router
}
