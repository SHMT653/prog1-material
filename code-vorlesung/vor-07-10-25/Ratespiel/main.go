package main

import (
	"fmt"
	"math/rand"
)

func ReadNumber() int {
	fmt.Print("Rate eine Zahl: ")
	var zahl int
	fmt.Scan(&zahl)
	return zahl
}

func GuessingGame() {
	my_number := rand.Intn(101) - 50
	for n := 0; n < 3; n++ {
		guess := ReadNumber()
		if Random(guess, my_number) {
			fmt.Println("Richtig geraten)")
			return
		} else {
			fmt.Println("Das ist leider die Falsche Zahl")
		}
	}
	fmt.Println("Zu viele falsche Zahlen")
}

func Random(g, e int) bool {
	return g == e
}

func main() {
	GuessingGame()
}
