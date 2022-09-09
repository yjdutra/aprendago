// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Mytype int

var x Mytype
var y int

func main() {

	fmt.Println("1.", x)
	fmt.Printf("2. %T\n", x)
	x = 42
	fmt.Println("4.", x)

	fmt.Println("Inicio ex. 05")

	fmt.Println("1. Convertendo x em inteiro e atribuindo seu valor a y")
	y = int(x)
	fmt.Println("2. Valor de y = ", y)
	fmt.Printf("3. O tipo de y = %T\n", x)
}
