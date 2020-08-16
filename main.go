//* Nombre de Campos y metodos exportados o No exportados en lugar de modificadores de acceso
//* Para saber que metodos se pueden extraer y cuales no
package main

import (
	"fmt"

	"github.com/Ajfero/Encapsulacion.git/course"
)

func main() {
	Go := &course.Course{
		Name:   "Go desde Cero",
		Price:  12.34,
		IsFree: false,
		UserID: []uint{12, 56, 89},
		Temario: map[uint]string{
			1: "Introduccion",
			2: "Estructuras",
			3: "Maps",
		},
	}
	fmt.Println(Go.Name)
}
