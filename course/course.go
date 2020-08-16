//* Se pueden declarar metodos fuera del archivo main siempre y cuando esten en el mismo paquete
//* Para realizar esta tarea se debe generar el binario "go build"
package course

import "fmt"

//* No exportamos la estructura
type course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserID  []uint
	Temario map[uint]string
}

func (c *course) PrintClass() {
	text := "Las clases son: "
	for _, temas := range c.Temario {
		text += temas + ", "
	}
	fmt.Println("Methods- " + text[:len(text)-2])
}

func (c *course) ChangePrice(Price float64) {
	c.Price = Price
}

//* En caso de necesitar Constructor
func New(name string, price float64, isFree bool) *course {
	if price == 0 {
		price = 30
	}
	return &course{
		Name:   name,
		Price:  price,
		IsFree: isFree,
	}
}
