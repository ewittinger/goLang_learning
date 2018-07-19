// Types
package main

import (
	"fmt"
)

func add(x float64, y float64) float64 {
	return x + y
}

func add1(x, y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string) {
	return a, b
}

func main() {
	var num1 float64 = 5.6 //option 1
	var num2 float64 = 9.5
	//	var num3, num4 float64 = 5.6, 9.5  // option 2
	//	var num5, num6 := 5.6, 9.5  // option 3, infered type (go will figure it out)
	// var a int = 62
	// var b float32 = float32(a) // converting a type
	// x := a                     // x will be type int
	fmt.Println(add(num1, num2))
	fmt.Println(add1(num1, num2))

	w1, w2 := "Hey", "there"
	fmt.Println(multiple(w1, w2))

}
