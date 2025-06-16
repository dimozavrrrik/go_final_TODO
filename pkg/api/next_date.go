package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const FormatDate = "20060102" // формат даты

// afterNow возвращает true, если дата больше текущей
func afterNow(date, now time.Time) bool {
	return date.After(now)
}

// NextDate возвращает следующую дату
func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" { // если повторение не указано
		return "", fmt.Errorf("правила повторения не указаны")
	}
	startDate, err := time.Parse(FormatDate, dstart) // парсит начальную дату и приводит ее к нужному формату
	if err != nil {
		return "", fmt.Errorf("неверная начальная дата")
	}
	dateSplit := strings.Split(repeat, " ") // разделяет повторение на день и год
	switch dateSplit[0] {
	case "d": // если повторение по дням
		if len(dateSplit) != 2 { // если повторение не указано
			return "", fmt.Errorf("неверный формат повторения")
		}
		interval, err := strconv.Atoi(dateSplit[1]) // парсит интервал
		if err != nil || interval < 1 || interval > 400 {
			return "", fmt.Errorf("неверный интервал повторения")
		}
		date := startDate // начальная дата

		for {
			date = date.AddDate(0, 0, interval) // прибавляет интервал
			if afterNow(date, now) {            // если дата больше текущей
				return date.Format(FormatDate), nil
			}
		}
	case "y": // если повторение по годам
		date := startDate

		for {
			date = date.AddDate(1, 0, 0) // прибавляет год
			if afterNow(date, now) {
				return date.Format(FormatDate), nil
			}
		}
	default:
		return "", errors.New("неверный формат повторения")
	}
}

// NextDateHadnler получает дату и повторение и возвращает следующую дату
func NextDateHadnler(w http.ResponseWriter, r *http.Request) {
	nowStr := r.FormValue("now")
	dstart := r.FormValue("date")
	repeat := r.FormValue("repeat")
	var now time.Time
	var err error
	if nowStr == "" {
		now = time.Now()
	} else {
		now, err = time.Parse(FormatDate, nowStr)
		if err != nil {
			http.Error(w, "Неверный формат записи даты", http.StatusBadRequest)
			return
		}
	}

	next, err := NextDate(now, dstart, repeat) // получаетследующую дату
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(next)) // возвращает следующую дату
}
