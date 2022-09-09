package main

import "fmt"

// Declarando variaveis em package-level scope
var x int
var y string
var z bool

func main() {
	// Printando os valores (no caso todos possuem valor "nulo" respectivo para seu tipo)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
