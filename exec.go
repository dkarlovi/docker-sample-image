package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("exec\n")

	if os.Args[0] == "exec" {
		fmt.Println("(using exec as entrypoint)")
	}

	fmt.Println("==> Defined ARGV:")
	for _, arg := range os.Args {
		fmt.Println(arg)
	}
	fmt.Println()

	var cmd *exec.Cmd
	if len(os.Args) > 2 {
		fmt.Println("==> Executing \"" + os.Args[1] + " " + os.Args[2] + "\":")
		cmd = exec.Command(os.Args[1], os.Args[2])
	} else {
		fmt.Println("==> Executing \"" + os.Args[1] + "\":")
		cmd = exec.Command(os.Args[1])
	}

	output, _ := cmd.CombinedOutput()

	if len(output) > 0 {
		fmt.Printf("==> Output:\n%s\n", output)
	}
}
