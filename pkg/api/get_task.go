package api

import (
	"log"
	"net/http"

	"go1f/pkg/db"
)

type TasksResp struct { // структура ответа
	Tasks []*db.Task `json:"tasks"`
}

// getTaskHandler получает список задач
func getTaskHandler(w http.ResponseWriter, r *http.Request) {

	tasks, err := db.Tasks(10) // получение списка задач
	if err != nil {
		log.Printf("Ошибка при получении задач: %v\n", err)
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJson(w, TasksResp{ // отправка списка задач
		Tasks: tasks,
	})
}

// getTaskID получает задачу по id
func getTaskID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") // получение id

	if id == "" { // проверка наличия id
		writeError(w, "Не указан id задачи", http.StatusBadRequest)
		return
	}

	task, err := db.GetTaskFromID(id) //  получение задачи
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJson(w, task) // отправка задачи
}
