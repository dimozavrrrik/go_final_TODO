package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

// DeleteTask удаляет задачу из БД
func DeleteTask(id string) error {
	query := `DELETE FROM scheduler WHERE id = :id`
	res, err := db.Exec(query, sql.Named("id", id))
	if err != nil {
		log.Printf("Ошибка при удалении задачи: %v\n", err)
		return err
	}
	count, err := res.RowsAffected() // получение количества обновленных строк
	if err != nil {
		log.Println(err)
		return err
	}
	if count == 0 { // проверка наличия обновленных строк
		return fmt.Errorf("задача с id %s не найдена", id)
	}
	return nil
}
