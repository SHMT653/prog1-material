package main

import (
	"fmt"
	"math/rand"
)

func main() {
	Guessing()
}

func Number() int {
	var zahl int
	fmt.Print("Rate eine Zahl zwischen 0-20: ")
	fmt.Scan(&zahl)
	return zahl
}

func Guessing() {
	random := rand.Intn(21)
	for n := 0; n < 3; n++ {
		guess := Number()
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
