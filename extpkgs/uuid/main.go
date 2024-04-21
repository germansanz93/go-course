package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New().String()
	fmt.Println(id)

	id2 := uuid.NewString() //Me hace el new y el string en un solo metodo
	fmt.Println(id2)

	fmt.Println(uuid.New().Version())

}
