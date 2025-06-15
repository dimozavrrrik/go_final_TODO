package api

import (
	"log"
	"net/http"
	"time"

	"go1f/pkg/db"
)

func postDoneHadnler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		log.Printf("Не указан id задачи\n")
		writeError(w, "Не указан id задачи", http.StatusBadRequest)
		return
	}
	task, err := db.GetTaskFromID(id)
	if err != nil {
		log.Printf("Ошибка при получении задачи: %v\n", err)
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if task.Repeat == "" {
		if err := db.DeleteTask(id); err != nil {
			log.Printf("Ошибка при удалении задачи: %v\n", err)
			writeError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeJson(w, map[string]string{})
		return
	}

	next, err := NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		log.Printf("Ошибка при получении следующей даты: %v\n", err)
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := db.UpdateDoneTask(next, id); err != nil {
		log.Printf("Ошибка при обновлении задачи: %v\n", err)
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJson(w, map[string]string{})
}
