package mypackage

import "fmt"

func Run() {
	defer func() {
		fmt.Println("end my function")
	}()
	panic("panic in run function")
}
