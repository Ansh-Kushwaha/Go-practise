package main

import (
	"fmt"

	"github.com/Ansh-Kushwaha/Go-practise/tree/main/basics/modules/greetings"
)

func main() {
	message := greetings.Greet("Ansh")
	fmt.Println(message)
}
