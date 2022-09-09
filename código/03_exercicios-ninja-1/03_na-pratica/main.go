package main

import "fmt"

// Declarando variaveis e atibuindo valores as mesmas em package-level scope
var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	// Atribuindo o valor de x, y e z convertidos para string a variavel s
	s := fmt.Sprintf("%v %v %v", x, y, z)
	// Printando valor de s
	fmt.Println(s)
}
