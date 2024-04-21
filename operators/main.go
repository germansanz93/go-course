package main

import "fmt"

func main() {

	yearsOld := 33

	fmt.Println(yearsOld > 30)
	fmt.Println(yearsOld < 33)
	fmt.Println(yearsOld <= 33)
	fmt.Println(yearsOld >= 33)
	fmt.Println(yearsOld == 33)
	fmt.Println(yearsOld != 33)
	fmt.Println(yearsOld < 33 || yearsOld == 33)
	fmt.Println(yearsOld < 33 || yearsOld == 34)
	fmt.Println(yearsOld < 40 && yearsOld == 33)
	fmt.Println(yearsOld < 40 && yearsOld > 33)
	fmt.Println(yearsOld > 40 && yearsOld <= 33)

}
