package main

import (
	"fmt"
	"os"
)

func main() {
	var name = os.Getenv("NAME")

	if name == "" {
		name = "World"
	}

	fmt.Println("Hello " + name + "!")

	fmt.Println("\nDefined environment variables:")
	for _, e := range os.Environ() {
		fmt.Println(e)
	}
}
