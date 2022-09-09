package main

import "fmt"

// Criando um novo tipo
type voceconheceositedosmenes int

// declarando uma variavel x do tipo criado
var x voceconheceositedosmenes

func main() {
	// imprime o valor de x, ( no caso 0 )
	fmt.Println(x)
	// impreme o tipo da variavel x
	fmt.Printf("%T\n", x)
	// atribui 42 a variavel x
	x = 42
	// imprime o valor de x, ( no caso 42 )
	fmt.Println(x)
}
