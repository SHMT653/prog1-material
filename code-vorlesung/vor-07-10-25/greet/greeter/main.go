package main

import (
	"fmt"
	"prog1-material/code-vorlesung/vor-07-10-25/greet"
)

func main() {
	fmt.Println(greet.Greet("Kurs"))
	fmt.Println("Alle" + greet.Greet("Kursteilnehmer"))
}
