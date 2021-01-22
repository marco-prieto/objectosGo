package course

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

//Metodo papu se declara x fuera de la estructura
//esa variable c es el receptor
//aca no c trabaja con this perro osea confunde no usar nombres genericos como this o self
//Cambiar datos tipo puntero y solo usar o ver datos normal
//Para mas placer pongo a todos tipo puntero si al menos 1 va con tipo puntero pa trabajar con las interfaces
func (c *Course) PrintClasses() {
	text := "Las clases son: "
	for _, class := range c.Classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}

//punteros pa cambiar
func (c *Course) ChangePrice(price float64) {
	c.Price = price
}
