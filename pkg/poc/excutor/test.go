package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("python3", "test.py")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}
