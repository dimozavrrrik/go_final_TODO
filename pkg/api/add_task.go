package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go1f/pkg/db"
)

// checkDate проверяет дату на совпадение с текущей
func checkDate(task *db.Task) error {
	now := time.Now() // текущая дата
	if task.Date == "" {
		task.Date = now.Format(FormatDate)
	}
	t, err := time.Parse(FormatDate, task.Date) // парсит дату и приводит ее к нужному формату
	if err != nil {
		return err
	}
	var next string        // следующая дата
	if task.Repeat != "" { // если есть правило повторения
		next, err = NextDate(now, task.Date, task.Repeat) // получаем следующую дату
		if err != nil {
			return err
		}
	}
	if afterNow(now, t) { // если дата больше текущей
		if len(task.Repeat) == 0 {
			task.Date = now.Format(FormatDate)
		} else {
			task.Date = next
		}
	}
	return nil
}

// addTaskHandler добавляет задачу
func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task                             // структура задачи
	err := json.NewDecoder(r.Body).Decode(&task) // десериализация
	if err != nil {
		writeError(w, "Ошибка при десериализации", http.StatusBadRequest)
		return
	}
	if task.Title == "" { // проверка наличия заголовка
		writeError(w, "Не указан заголовок задачи", http.StatusBadRequest)
		return
	}
	if err := checkDate(&task); err != nil { // проверка даты
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := db.AddTask(&task) // добавление задачи
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJson(w, map[string]int64{"id": id}) // возврат идентификатора
}

// writeJson отправляет json ответ
func writeJson(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("Ошибка при сериализации:", err)
	}

}

// writeError отправляет ошибку в ответ
func writeError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})

}
