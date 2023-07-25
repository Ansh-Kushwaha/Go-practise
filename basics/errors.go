package main

import (
	"errors"
	"fmt"
	"log"
)

func divide(a int, b int) (float32, error) {
	if b == 0 {
		return -1, errors.New("division by zero")
	}
	var c float32 = float32(a) / float32(b)
	return c, nil
}

func main() {
	a := 10
	b := 0
	c, err := divide(a, b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
}
