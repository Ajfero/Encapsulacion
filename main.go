//* Nombre de Campos y metodos exportados o No exportados en lugar de modificadores de acceso
//* Para saber que metodos se pueden extraer y cuales no
package main

import (
	"fmt"
)

func main() {
	Go := Course{
		"Go desde Cero",
		12.34,
		false,
		[]uint{12, 56, 89},
		map[uint]string{
			1: "Introduccion",
			2: "Estructuras",
			3: "Maps",
		},
	}

	Go.PrintClass()
	Go.ChangePrice(67.12)
	fmt.Println(Go.IsFree)
}
