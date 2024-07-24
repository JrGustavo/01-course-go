package arreglos_slices

import "fmt"

var tabla [10]int

func MuestroAreglos() {
	tabla[7] = 33
	tabla[2] = 54

	fmt.Println(tabla)
}
