package main

import (
	"fmt"
	"math/rand"
)

// Main ruft das Spiel auf.
func main() {
	Guessing()
}

// NumberInput gibt eine Zahl (int) zurück.
func NumberInput() int {
	var zahl int
	fmt.Print("Rate eine Zahl zwischen 0-20: ")
	fmt.Scan(&zahl) // Adresse von zahl übergeben (&)
	return zahl
}

// Guessing ist das Spiel.
func Guessing() {
	random := rand.Intn(21)
	for n := 0; n < 3; n++ { // schleife für drei Versuche
		guess := NumberInput()
		if Check(guess, random) {
			fmt.Println("Richtig geraten)")
			return
		} else {
			fmt.Println(guess, "war leider die falsche Zahl. ")
		}
	}
	fmt.Println("Zu viele falsche Versuche, die richtige Zahl wäre", random, "gewesen")
}

// Check überprüft ob der Guesse und die Random Zahl gleich sind und gibt true or false aus (bool)
func Check(g, r int) bool { // g = guess, r = random
	return g == r // true wenn gleich, sonst false
}
