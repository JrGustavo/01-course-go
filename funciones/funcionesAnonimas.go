package funciones

import "fmt"

func Calculos() {

	var numero3 int = 32
	var numero4 int = 234

	calculo := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero4 + numero3
	}

	fmt.Println(calculo(10, 25))

	calculo = func(numero1 int, numero2 int) int {
		return numero1 * numero2 * numero3
	}

	fmt.Println(calculo(10, 25))

}
