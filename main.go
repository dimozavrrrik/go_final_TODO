package main

import (
	"go1f/pkg/db"
	"go1f/pkg/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// загрузка файла .env
	err := godotenv.Load()
	if err != nil {
		log.Println("ошибка загрузки .env")
		return
	}
	// Подключение к бд
	err = db.Init()
	if err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
	}
	// Запуск сервера
	server.Run()
}
