package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go [from] [to] [value]")
		return
	}
	from := os.Args[1]
	to := os.Args[2]
	n, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Error: Number value not valid")
		return
	}
	var res float64

	switch from {
	case "c":
		switch to {
		case "k":
			res = CtoK(n)
		case "f":
			res = CtoF(n)
		default:
			fmt.Println("Error: 'To' value not valid")
			return
		}
	case "f":
		switch to {
		case "k":
			res = FtoK(n)
		case "c":
			res = FtoC(n)
		default:
			fmt.Println("Error: 'To' value not valid")
			return
		}
	case "k":
		if n < 0 {
			fmt.Println("Error: K can't be negative")
			return
		}

		switch to {
		case "c":
			res = KtoC(n)
		case "f":
			res = KtoF(n)
		default:
			fmt.Println("Error: 'To' value not valid")
			return
		}

	default:
		fmt.Println("Error: 'From' value not valid")
		return
	}

	fmt.Printf("Result: %.2f%s", res, to)
}

func CtoF(c float64) float64 {
	return (c * 9 / 5) + 32
}

func CtoK(c float64) float64 {
	return c + 273.15
}

func FtoC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func FtoK(f float64) float64 {
	return (f-32)*5/9 + 273.15
}

func KtoC(k float64) float64 {
	return k - 273.15
}

func KtoF(k float64) float64 {
	return (k-273.15)*9/5 + 32
}
