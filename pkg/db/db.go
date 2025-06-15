package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

// подключение к БД
var db *sql.DB

// создание схемы БД
const schema = `
CREATE TABLE scheduler (
 id INTEGER  PRIMARY KEY AUTOINCREMENT,
 date CHAR(8) NOT NULL DEFAULT "",
 title VARCHAR(128) NOT NULL DEFAULT "",
 comment TEXT NOT NULL DEFAULT "",
 repeat VARCHAR(128) NOT NULL DEFAULT ""
 );
 CREATE INDEX idx_date ON scheduler (date);
 `

// Init инициализация БД
func Init() error {

	dbFile := os.Getenv("TODO_DBFILE") // переменная окружения
	if dbFile == "" {
		return fmt.Errorf("TODO_DBFILE не установлена")
	}

	var install bool          // флаг установки
	_, err := os.Stat(dbFile) // проверка наличия файла
	if err != nil {
		install = true
	}

	connect, err := sql.Open("sqlite", dbFile) // подключение к БД
	if err != nil {
		return err
	}

	if install {
		if _, err := connect.Exec(schema); err != nil {
			return err
		}
	}
	db = connect // сохранение подключения
	return nil
}
