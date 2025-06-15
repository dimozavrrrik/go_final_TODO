package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

// UpdateTask обновляет задачу
func UpdateTask(task *Task) error {

	query := `UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id` // формирование запроса в бд

	res, err := db.Exec(query, // выполнение запроса
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
		sql.Named("id", task.ID))
	if err != nil {
		log.Println(err)
		return err
	}
	count, err := res.RowsAffected() // получение количества обновленных строк
	if err != nil {
		log.Println(err)
		return err
	}
	// проверка наличия обновленных строк
	if count == 0 {
		return fmt.Errorf("задача с id %s не найдена", task.ID)
	}
	return nil
}
