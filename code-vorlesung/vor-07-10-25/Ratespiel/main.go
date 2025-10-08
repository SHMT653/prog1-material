package main

import (
	"fmt"
	"math/rand"
)

func main() {
	GuessingGame()
}

func ReadNumber() int {
	fmt.Print("Rate eine Zahl: ")
	var zahl int
	fmt.Scan(&zahl)
	return zahl
}

func GuessingGame() {
	random := rand.Intn(101) - 50
	for n := 0; n < 3; n++ {
		guess := ReadNumber()
		if Check(guess, random) {
			fmt.Println("Richtig geraten)")
			return
		} else {
			fmt.Println("Das ist leider die Falsche Zahl")
		}
	}
	fmt.Println("Zu viele falsche Versuche, die richtige Zahl wäre", random, "gewesen")
}

func Check(g, r int) bool {
	return g == r
}
