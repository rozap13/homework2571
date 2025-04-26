package main

import (
	"fmt"
	"log"
)

func main() {
	var input string
	fmt.Print("Введите данные: ") // Исправлено
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели следующие данные: %s\n", input) // Исправлено
}
