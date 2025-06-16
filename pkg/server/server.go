package server

import (
	"fmt"

	"net/http"
	"os"
	"strconv"

	"go1f/pkg/api"
)

// Run запускает сервер
func Run() {
	api.Init() // инициализация api-handler

	port := os.Getenv("TODO_PORT") // переменная окружения
	if port == "" {
		port = "7540"
	}
	adrr, err := strconv.Atoi(port) // преобразование в int
	if err != nil {
		fmt.Println("Ошибка переменной окружения TODO_PORT")
	}
	address := fmt.Sprintf(":%d", adrr)                // формирование адреса
	http.Handle("/", http.FileServer(http.Dir("web"))) // обработка статических ресурсов

	http.ListenAndServe(address, nil) // запуск сервера
}
