//* Se pueden declarar metodos fuera del archivo main siempre y cuando esten en el mismo paquete
//* Para realizar esta tarea se debe generar el binario "go build"
package course

import "fmt"

type Course struct {
	Name    string
	Price   float64
	isFree  bool
	UserID  []uint
	Temario map[uint]string
}

func (c *Course) PrintClass() {
	text := "Las clases son: "
	for _, temas := range c.Temario {
		text += temas + ", "
	}
	fmt.Println("Methods- " + text[:len(text)-2])
}

func (c *Course) ChangePrice(Price float64) {
	c.Price = Price
}
