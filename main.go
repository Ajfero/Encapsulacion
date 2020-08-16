//* Nombre de Campos y metodos exportados o No exportados en lugar de modificadores de acceso
//* Para saber que metodos se pueden extraer y cuales no
package main

import (
	"fmt"

	"github.com/Ajfero/Encapsulacion.git/course"
)

func main() {
	Go := course.New("Go desde Cero", 0, false)
	Go.UserID = []uint{12, 56, 89}
	Go.Temario = map[uint]string{
		1: "Introduccion",
		2: "Estructuras",
		3: "Maps",
		}
	}
	Go.PrintClass()
}
