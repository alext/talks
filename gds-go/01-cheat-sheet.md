# GDS Go Workshop #1: Cheatsheet

## Value assignment

```go
var s1 string         // a new string identifer with its zero value ("")
var s2 string = "foo" // an identifer with a non-zero value
var i1 int            // an int identifer with its zero value (0)
var i2 int = 42       // an int identifer with a non-zero value

s3 := "foo"           // new string identifier via implicit typing
i3 := 41              // new int identifier via implicit typing

i3 = 42               // assignment to existing identifier

res, err := myFunc()  // multiple return value assignment using implicit typing
```

## Function definitions

```go
func f1() {}                 // function with no arguments and no return values
func f2(s string) {}         // one argument
func f3(s string, i int) {}  // two arguments of different types
func f4(s1, s2 string) {}    // two arguments of the same type
func f5() string {}          // one return value
func f6() (string, error) {} // two return values
```

## Slices

```go
var s1 []int     // int slice identifier with zero value (empty slice)
var s2 []string  // string slice identifier with zero value (empty slice)

s3 := []int{1, 1, 2, 3, 5, 8}               // int slice with 6 values
s4 := []string{"foo", "bar", "boo", "baz"}  // string slice with 4 values

s5 := []string{"foo"}
s5 = append(s5, "bar")  // appending a new value to an existing slice
```

## Struct definition

```go
type myStruct struct {
	Foo string  // exported field
	bar int     // unexported field
}

var s1 myStruct             // new myStruct value with fields set to zero value
s2 := myStruct{}            // as above, using implicit typing
s3 := myStruct{"foo", 42}   // new value with specific values using field order
s4 := myStruct{bar: 42}     // new value using keyed field reference

fmt.Println(s3.Foo, s4.bar) // "foo" 42
```

## Error handling

```go
func myErroringFunc() (string, error) {
	return "", errors.New("uh oh!")
}

s, err := myErroringFunc()
if (err != nil) {
	// handle the error state here
}
```

## Simple tests

In `foo_test.go`:

```go
import "testing"

func TestFoo(t *testing.T) {
	expected := "foo"
	actual := "bar"

	if expected != actual {
		t.Errorf("Unexpected value: got %v want %v", actual, expected)
	}
}
```

Then run `go test foo_test.go` (or just `go test` to run all `*_test.go`
files).

## Table-based tests

```go
func TestFoo(t *testing.T) {
	var tests = []struct {
		input  string
	    output int
	}{
		{"foo", 42},
		{"bar", 99},
		{"boo", 123},
	}

	for _, test := range tests {
	    res := someFunc(test.input)
		if res != test.output {
			t.Errorf("Unexpected return value: got %v want %v", res, test.output)
		}
	}
}
```

## Creating a basic HTTP server using http.HandlerFunc

```go
import "net/http"

func handle(w http.ResponseWriter, r *http.Request) {
	// ...
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
}
```

Package documentation for `net/http` can be found at
[https://godoc.org/net/http][net/http].

[net/http]: https://godoc.org/net/http

## Marshaling to JSON

```go
import "encoding/json"

type Rectangle struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

r := Rectangle{3, 4}
j, _ := json.Marshal(r) // return values are ([]byte, error)
fmt.Println(string(j))  // {"width":3,"height":4}
```

Package documentation for `encoding/json` can be found at
[https://godoc.org/encoding/json][enc/json].  Additional guides to using JSON
with Go can be found [here][json-1], [here][json-2] and [here][json-3].

[enc/json]: https://godoc.org/encoding/json
[json-1]: https://blog.golang.org/json-and-go
[json-2]: hhttps://eager.io/blog/go-and-json/
[json-3]: hhttps://gobyexample.com/json
