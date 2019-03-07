package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if rand.Float64() <= 0.2 {
		fmt.Println(" _    _          _____ _  ________ _____  _")
		fmt.Println("| |  | |   /\\   / ____| |/ |  ____|  __ \\| |")
		fmt.Println("| |__| |  /  \\ | |    | ' /| |__  | |  | | |")
		fmt.Println("|  __  | / /\\ \\| |    |  < |  __| | |  | | |")
		fmt.Println("| |  | |/ ____ | |____| . \\| |____| |__| |_|")
		fmt.Println("|_|  |_/_/    \\_\\_____|_|\\_|______|_____/(_)")
		fmt.Println()
		fmt.Println("You have just been HACKED by a state-of-the-art stealth virus embedded in a Docker image you just run off the Internet.")
		fmt.Println()
		fmt.Println("Since this is still only a prototype virus and most of the time so far was spent on the above ASCII art,")
		fmt.Println("please delete some of your important files by yourself in the meantime.")
		fmt.Println("Thanks.")
		os.Exit(0)
	}

	var env = os.Getenv("SHOW_ENVS")

	var name = "World"
	var argc = len(os.Args)

	if argc >= 2 {
		name = os.Args[1]
	}

	var hello = "Hello " + name + "!\n"
	if argc >= 3 {
		var filename = os.Args[2]
		f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		if _, err = f.WriteString(time.Now().Format("2006-01-02 15:04:05") + " " + hello); err != nil {
			panic(err)
		}
	} else {
		fmt.Print(hello)
	}

	if env != "" {
		fmt.Println("\nDefined environment variables:")
		for _, e := range os.Environ() {
			fmt.Println(e)
		}
	}
}
