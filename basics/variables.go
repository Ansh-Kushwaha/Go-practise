package basics

// A variable should always be used after declaration

import "fmt"

// Package level variable
var i int = 0
var I int = 1 // can be accessed from other modules

func Variables() {
	// Method 1
	var num int
	num = 15
	fmt.Printf("num = %v \n", num)

	// Method 2 with initializer
	var num1 int = 12
	fmt.Printf("num1 = %v \n", num1)

	// Method 3 (Short)
	num2 := 100
	fmt.Printf("num2 = %v \n", num2)

	// %T to get type of variable
	fmt.Printf("i = %v %T \n", i, i)

	// var num3 = 2 // no need to type with initializer (compile time error)
}
