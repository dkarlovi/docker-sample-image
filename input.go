package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var stat, _ = os.Stdin.Stat()

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		fmt.Println("No STDIN to read from")
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
