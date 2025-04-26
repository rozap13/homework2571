package main

import (
	"fmt"
	"log"
)

func main() {
	var input string
	fmt.Print("Введите целое число: ") // ОШИБКА: осталось старое сообщение
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели число: %s\n", input) // ОШИБКА: осталось "число"
}
