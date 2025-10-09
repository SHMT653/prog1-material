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
	fmt.Scan(&zahl) // & zahl übergibt die Speicheradresse, Scan schreibt hierhin
	return zahl
}

// Guessing ist das Spiel.
func Guessing() {
	random := rand.Intn(21)  // Intn(21) liefert eine Zahl von 0 bis 20 (21 exklusiv)
	for n := 0; n < 3; n++ { // drei Versuche: n = 0,1,2
		guess := NumberInput()    // liest Zahl vom Nutzer
		if Check(guess, random) { // true wenn geraten == random
			fmt.Println("Richtig geraten)") // Erfolgsmeldung
			return                          // Funktion sofort beenden (Spiel vorbei)
		} else {
			fmt.Println(guess, "war leider die falsche Zahl. ") // mehrere Argumente werden mit Leerzeichen ausgegeben
		}
	}
	fmt.Println("Zu viele falsche Versuche, die richtige Zahl wäre", random, "gewesen")
}

// Check überprüft ob der Guesse und die Random Zahl gleich sind und gibt true or false aus (bool)
func Check(g, r int) bool {
	return g == r
}
