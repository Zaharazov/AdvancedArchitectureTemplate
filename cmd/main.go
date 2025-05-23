package main

import (
	"log"
	nethttp "net/http"

	"github.com/Zaharazov/AdvancedArchitectureTemplate/internal/config"
	"github.com/Zaharazov/AdvancedArchitectureTemplate/internal/infrastructure/db"
	"github.com/Zaharazov/AdvancedArchitectureTemplate/internal/transport/http"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Файл .env не найден")
	}

	cfg := config.LoadDBConfig()

	dbConn, err := db.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer dbConn.Close()

	router := http.NewRouter(dbConn)

	log.Println("Сервер запущен на порту 8080")
	if err := nethttp.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
