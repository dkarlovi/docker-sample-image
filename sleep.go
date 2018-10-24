package main

import (
	"os"
	"strconv"
	"time"
)
import "fmt"

func main() {
	var sec int

	sec = 60
	if len(os.Args) > 1 {
		sec, _ = strconv.Atoi(os.Args[1])
	}

	fmt.Println("Sleeping for " + strconv.Itoa(sec) + " seconds")

	c1 := make(chan string, 1)
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Duration(sec) * time.Second):
		fmt.Println("Slept for " + strconv.Itoa(sec) + " seconds")
		os.Exit(0)
	}
}
