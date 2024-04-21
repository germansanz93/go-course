package main

import (
	"encoding/json"
	"fmt"

	"github.com/germansanz93/go-course/structs/commerce"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Address  string `json:"address,omitempty"` //En el caso de indicar en el tag omitempty lo que hace es que si el campo va vacio no se envia al json
	Age      uint8  `json:"age"`
}

func (u User) Display() { //Para declarar un metodo, lo hacemos exactamente igual que cuando declaramos una funcion con la diferencia que antes del nombre del metodo ponemos entre parentesis el struct
	jsonString, _ := json.Marshal(u)
	fmt.Println(string(jsonString))
}

func (u User) setName(name string) { //Ojo que solo modifica dentro de su scope, cuando la var u muere, el cambio se pierde
	u.Name = name
}

func (u *User) setName2(name string) { //Aca si se modifica el valor almacenado en el struct. Esto se dio porque go es passed by value, entonces necesitamos indicar con el puntero que queremos afectar a la var original
	u.Name = name
}

func main() {
	user := User{
		ID:       123,
		Name:     "German",
		LastName: "Sanz",
		Address:  "Buenos Aires 1268",
		Age:      30,
	}

	user2 := User{
		ID:       124,
		Name:     "Rene",
		LastName: "Sanz",
	}

	fmt.Println(user)
	user.Display()

	fmt.Println(user2)
	user2.Display()

	user2.setName("Alberto")
	user2.Display()

	user2.setName2("Alberto")
	user2.Display()

	p1 := commerce.Product{
		Name:  "Heladera marca Whirlpool",
		Price: 200000,
		Type: commerce.Type{
			Code:        "A",
			Description: "Electrodomestico",
		},
		Tags:  []string{"heladera", "freezer", "whirlpool", "refrigerador"},
		Count: 5,
	}

	p2 := commerce.Product{
		Name:  "Naranja",
		Price: 50,
		Type: commerce.Type{
			Code:        "B",
			Description: "Alimento",
		},
		Tags:  []string{"alimento", "verdura"},
		Count: 14,
	}

	p3 := commerce.Product{
		Name:  "Cortina",
		Price: 3000,
		Type: commerce.Type{
			Code:        "C",
			Description: "Hogar",
		},
		Tags:  []string{"hogar", "cortina"},
		Count: 3,
	}

	cart := commerce.NewCart(11312)

	cart.AddProducts(p1, p2, p3)

	fmt.Println("PRODUCTS CART")
	fmt.Println("Total Products:", len(cart.Products))
	fmt.Printf("Total Cart: %.2f\n", cart.Total())

}
