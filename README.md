# Планировщик задач
Веб-приложение, предназначенное для ведения списка задач с возможностью указания даты, правил повторения и с сохранение в БД SQLite

## Задания со звездочкой:

Использование переменных окружения:
TODO_PORT - порт, на котором запускается веб-сервер
TODO_LIST - путь к базе данных

## Параметры для запуска тестов test/settings:

package tests

var Port = 7540
var DBFile = "../scheduler.db"
var FullNextDate = false
var Search = false
var Token = ``

## Запуск: 
go run main.go

