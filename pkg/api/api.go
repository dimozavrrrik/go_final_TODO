package api

import (
	"net/http"
)

// Init инициализация api
func Init() {
	http.HandleFunc("/api/nextdate", NextDateHadnler)
	http.HandleFunc("/api/task", taskHandler)
	http.HandleFunc("/api/task/done", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			postDoneHadnler(w, r)
		}
	})
	http.HandleFunc("/api/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTaskHandler(w, r)
		}
	})

}

// taskHandler обработчик задач
func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addTaskHandler(w, r)
	case http.MethodGet:
		getTaskID(w, r)
	case http.MethodPut:
		putTaskHandler(w, r)
	case http.MethodDelete:
		DeleteTaskHandler(w, r)
	}
}
