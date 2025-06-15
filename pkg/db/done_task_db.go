package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

// UpdateDoneTask обновляет дату выполнения задачи
func UpdateDoneTask(nextDate, id string) error {
	query := `UPDATE scheduler SET date = :date WHERE id = :id` // формирование запроса в бд
	// выполнение запроса
	res, err := db.Exec(query, // выполнение запроса
		sql.Named("date", nextDate),
		sql.Named("id", id))
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
		return fmt.Errorf("задача с id %s не найдена", id)
	}
	return nil
}
