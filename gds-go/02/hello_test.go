// hello_test.go
package main

// START OMIT
import "testing"

func TestGreeting(t *testing.T) {
	expected := "Hello Foo!"
	actual := greeting("Foo")
	if actual != expected {
		t.Errorf("Want '%s': got '%s'", expected, actual)
	}
}

// END OMIT
