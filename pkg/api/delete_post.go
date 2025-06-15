package api

import (
	"log"
	"net/http"

	"go1f/pkg/db"
)

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		log.Printf("Не указан id задачи\n")
		writeError(w, "Не указан id задачи", http.StatusBadRequest)
		return
	}

	if err := db.DeleteTask(id); err != nil {
		log.Printf("Ошибка при удалении задачи: %v\n", err)
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJson(w, map[string]string{})
}
