package main

import "fmt"

func main() {
	/* Go := Course{
		Name:    "Go ez",
		Price:   12.30,
		IsFree:  false,
		UserIDs: []uint{12, 56, 89},
		Classes: map[uint]string{
			1: "Introduccion",
			2: "Estructuras",
			3: "Mapas",
		},
	} */
	//si vas a llenar todos los campos y en orden
	Go := Course{
		"Go ez",
		12.30,
		false,
		[]uint{12, 56, 89},
		map[uint]string{
			1: "Introduccion",
			2: "Estructuras",
			3: "Mapas",
		},
	}
	fmt.Println(Go.Name)
	//Q pasa si quiero crear una instancia pero no quiero llenar todos los campos
	css := Course{
		Name:   "Css mas ezz aun",
		IsFree: true,
	}
	fmt.Printf("%+v\n", css)

	js := Course{}
	js.Name = "Curso JavaScript"
	js.UserIDs = []uint{12, 67}
	fmt.Printf("%+v\n", js)
	//Go ya se dara cuenta si tiene q pasar & o * viendo como esta creada la funcion
	Go.PrintClasses()
	Go.ChangePrice(43.21)
	fmt.Println(Go.Price)

}
