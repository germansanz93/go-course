package main

import "fmt"

func main() {
	map1 := make(map[int]string)
	fmt.Println(map1)

	map1[1] = "A"
	map1[2] = "B"
	map1[3] = "C"
	map1[-5] = "D"

	fmt.Println(map1)
	fmt.Println(map1[-5]) //Accedo a un value mediante su key

	map2 := make(map[int][]string) //Map con keys integer y values slice de string
	map2[1] = []string{"1", "2", "3"}
	map2[14] = []string{"14", "24", "34"}
	fmt.Println(map2)

	delete(map1, 1)
	fmt.Println(map1)

	//Que pasa si quiero acceder a una key que no existe???
	fmt.Println(map2[2])   //no me devuelve nada, ni error ni nada, solo un listado vacio
	fmt.Println(map1[223]) //no me devuelve nada, ni error ni nada, solo un string vacio

	map3 := make(map[int]int)
	fmt.Println(map3[2]) //no me devuelve nada, solo un 0.. en fin, obtenemos zero values OJO

	v, ok := map2[1] //Un segundo valor es obtenido de la consulta de un key en un map y nos permite saber si el valor existe o estamos viendo un zero value
	v2, ok2 := map2[2]

	fmt.Println(v, ok)
	fmt.Println(v2, ok2)

	map4 := map[int]string{ //Instancio e inicializo
		4: "Perro",
		7: "Gato",
	}

	fmt.Println(map4)

	for k, v := range map1 {
		fmt.Println("key:", k, "value:", v)
	}

}
