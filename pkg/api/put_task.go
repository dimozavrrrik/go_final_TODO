package api

import (
	"encoding/json"
	"log"
	"net/http"

	"go1f/pkg/db"
)

func putTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Printf("Ошибка при десериализации: %v\n", err)
		writeError(w, "Ошибка при десериализации", http.StatusBadRequest)
		return
	}
	if task.ID == "" {
		log.Printf("Не указан id задачи\n")
		writeError(w, "Не указан id задачи", http.StatusBadRequest)
		return
	}
	if task.Title == "" {
		writeError(w, "Не указан заголовок задачи", http.StatusBadRequest)
		return
	}
	if err := checkDate(&task); err != nil {
		log.Printf("Ошибка при проверке даты: %v\n", err)
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.UpdateTask(&task)
	if err != nil {
		log.Printf("Ошибка при обновлении задачи: %v\n", err)
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJson(w, map[string]string{})
}
