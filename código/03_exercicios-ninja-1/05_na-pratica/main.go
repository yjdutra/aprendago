package main

import "fmt"

// Criando um novo tipo
type voceconheceositedosmenes int

// declarando uma variavel x do tipo criado
var x voceconheceositedosmenes

// declarando uma variavel y do tipo inteiro
var y int

func main() {
	// atribuuindo o valor  42 a variavel x
	x = 42
	// convertendo x para inteiro e atribuindo o valor a y
	y = int(x)
	// imprimindo o valor de y
	fmt.Println(y)
	// imprimindo o tipo de y
	fmt.Printf("%T\n", y)
}
