package main

import "fmt"

// START OMIT
import "encoding/json"

type Foo struct {
	Title  string
	Fruits []string
	active bool
}

func main() {
	f := Foo{Title: "something", Fruits: []string{"Apple", "Banana"}, active: true}
	data, _ := json.Marshal(f)
	fmt.Println(string(data))

	f = Foo{Title: "something else"}
	data, _ = json.Marshal(f)
	fmt.Println(string(data))
}

// END OMIT
