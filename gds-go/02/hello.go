package main

import "fmt"

// START OMIT
func main() {
	text := greeting("GDS")
	fmt.Println(text)
}

func greeting(name string) string {
	return "Hello " + name + "!"
}

// END OMIT
