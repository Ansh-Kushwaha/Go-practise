package basics

import (
	"fmt"
	"strconv"
)

func TypeConv() {
	i := 12
	fmt.Printf("i = %v, %T \n", i, i)

	// converting int to float64
	j := float64(i)
	fmt.Printf("j = %v, %T \n", j, j)

	// using strconv package to convert to string
	a := strconv.Itoa(i)
	fmt.Printf("a = %v, %T \n", a, a)
}