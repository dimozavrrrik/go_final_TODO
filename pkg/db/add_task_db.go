package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

// AddTask добавляет задачу
func AddTask(task *Task) (int64, error) {
	var id int64                                                                                             // идентификатор задачи
	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date,:title, :comment, :repeat)` // формирование запроса
	// выполнение запроса
	result, err := db.Exec(query,
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat))

	if err == nil {
		id, err = result.LastInsertId()
	}
	return id, err
}
