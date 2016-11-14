package main

import "fmt"

// START OMIT
func main() {
	text := greeting("GDS") // HL
	fmt.Println(text)
}

func greeting(name string) string { // HL
	return "Hello " + name + "!" // HL
} // HL

// END OMIT
