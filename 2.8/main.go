package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func main() {
	timeNow, err := ntp.Time("pool.ntp.org") // Запрашиваем точное текущее время с NTP-сервера pool.ntp.org
	if err != nil {                          //проверяем произошла ли ошибка
		fmt.Fprintln(os.Stderr, err) // если есть то выводим ее в stderr
		os.Exit(1)                   //Завершаем программу с ненулевым кодом выхода (1)
	}
	fmt.Println("Time:", timeNow.Format(time.RFC3339)) //печатаем точное время в формате RFC3339
}
