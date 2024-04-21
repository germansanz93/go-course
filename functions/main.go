package main

import (
	"fmt"

	"github.com/germansanz93/go-course/functions/function"
)

func main() {
	var myIntVar int64
	var myIntVar2 int64 = 10
	var myIntVar3 int64 = 20
	var myIntVar4 int64 = 35
	function.Display(myIntVar)
	function.Display(myIntVar2)
	function.Display(myIntVar3)
	function.Display(myIntVar4)

	fmt.Println(function.Add(4, 6))

	function.RepeatString(10, "LA")

	res, err := function.Calc(function.SUM, 2, 5)
	fmt.Println("SUM:", res, "ERR:", err)
	res, err = function.Calc(function.SUB, 2, 5)
	fmt.Println("SUB:", res, "ERR:", err)
	res, err = function.Calc(function.DIV, 2, 5)
	fmt.Println("DIV:", res, "ERR:", err)
	res, err = function.Calc(function.DIV, 2, 0)
	fmt.Println("DIV:", res, "ERR:", err)
	res, err = function.Calc(function.MUL, 2, 5)
	fmt.Println("MUL:", res, "ERR:", err)
	res, err = function.Calc(6, 2, 5)
	fmt.Println("MUL:", res)
	if err != nil {
		fmt.Println("err:", err) //asi se suele manejar errores en go
	}

	x, y := function.Split(40)
	fmt.Println("x:", x, "y:", y)

	fmt.Println(function.MSum(1, 3, 5))
	fmt.Println(function.MSum(1, 3, 5, 5, 1))

	fmt.Println(function.MOperations(function.SUM, 1, 2, 3, -32))

	factOpFunc := function.FactoryOperation(function.SUM)
	fmt.Println(factOpFunc(3, 15))
}
