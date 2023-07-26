package basics

import "fmt"

func Variables() {
	// Method 1
	var num int = 12
	fmt.Printf("num = %v", num)
	num = 12
	fmt.Printf("num = %v", num)

	// Method 2
	num2 := 100
	fmt.Printf("num2 = %v", num2)

}
