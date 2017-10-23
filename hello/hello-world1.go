package main //this is not necessary to be main always you can name it anything

import (
	"fmt"
	"foo"
)

var randomBug int

func main() {
	randomBug := 1000000000
	stupidStr := "hey stupid"
	var coolString string
	// var greetMsg string

	fmt.Println(randomBug)
	fmt.Println(stupidStr)

	fmt.Println("Hello, Haroon is new. Welcome home")
	fmt.Println("Random number here")
	d := foo.Bar()
	fmt.Println(d)
	// greetMsg = greeting("haroon") using above variable declaration
	greetMsg := greeting("haroon")
	fmt.Println(greetMsg)

	doodle, coodle := foo1(2, "asd")
	fmt.Println(doodle, coodle)

	fmt.Println(coolString)
}

func greeting(name string) string {
	return "Hello, " + name
}

func foo1(a int, b string) (doo int, coo string) {
	doo = 2
	coo = "bla bla"
	return
}
