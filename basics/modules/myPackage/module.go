package main

import (
	"fmt"

	"github.com/Ansh-Kushwaha/Go-practise/tree/mainbasics/modules/greetings"
)

func main() {
	message := greetings.Greet("Ansh")
	fmt.Println(message)
}
