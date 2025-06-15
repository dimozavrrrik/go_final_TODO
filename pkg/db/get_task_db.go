package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

// Tasks возвращает список задач
func Tasks(limit int) ([]*Task, error) {
	rows, err := db.Query("SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date LIMIT :limit", sql.Named("limit", limit)) // выполнение запроса
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close() // закрытие результата
	var tasks []*Task  // список задач
	// перебор результата
	for rows.Next() {
		temp := new(Task)
		err := rows.Scan(&temp.ID, &temp.Date, &temp.Title, &temp.Comment, &temp.Repeat) // считывание строки
		if err != nil {
			log.Println(err)
			return nil, err
		}
		tasks = append(tasks, temp) // добавление задачи в список
	}
	if tasks == nil { // если список задач пуст
		tasks = []*Task{}
	}
	return tasks, nil
}

func GetTaskFromID(id string) (*Task, error) {
	var task Task
	query := `SELECT id, date, title, comment, repeat FROM scheduler WHERE id = :id`
	err := db.QueryRow(query, sql.Named("id", id)).Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("задача с id %s не найдена", id)
		}
		log.Println(err)
		return nil, err
	}
	return &task, nil
}
