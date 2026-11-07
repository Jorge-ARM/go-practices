package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var globalVar = 100

func main() {
	var a int = 1
	fmt.Printf("Var a = %d\n", a)
	fmt.Printf("Var 'a' is %T\n", a)

	b := "This is a string"
	fmt.Printf("Var b = %s\n", b)
	fmt.Printf("Var 'b' is %T\n", b)

	c := "Variable"
	const d = "Constant"

	fmt.Printf("c = '%s'\n", c)
	c = "Another variable"
	fmt.Printf("But now c = '%s'\n", c)
	fmt.Printf("Var 'c' is %T\n", c)
	fmt.Printf("On the other hand, d = '%s' forever\n", d)
	fmt.Printf("Var 'd' is %T\n", d)

	var (
		e = 2.71828
		f = complex(0, math.Pi)
		g = cmplx.Exp(f)
	)

	fmt.Printf("Var e = %f\n", e)
	fmt.Printf("Var 'e' is %T\n", e)
	fmt.Printf("Var f = %f\n", f)
	fmt.Printf("Var 'f' is %T\n", f)
	fmt.Printf("Var g = %f\n", g)
	fmt.Printf("Var 'g' is %T\n", g)

	// Go is statically typed. A variable's type is determined when it is declared and cannot be changed through assignment.
	// A new variable with the same name can be declared in a nested scope, shadowing the original variable.

	price := 9.99

	{
		price := "Nine point nine nine"
		fmt.Printf("Price is %s\n", price)
	}

	fmt.Printf("Price = $%.2f\n", price) //Two decimals, %.2f

	var localVar = 200
	fmt.Printf("localVar = %d\n", localVar)
	showMe()

}

func showMe() {
	fmt.Printf("globalVar = %d\n", globalVar)
}
