package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("go", "build", "-o", "build/ellipsis")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}
}
