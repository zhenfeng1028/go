package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	ExecCmd("ls -l")
}

func ExecCmd(command string) {
	fmt.Println(command)
	cmd := exec.Command("bash", "-c", command)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("cmd.StdoutPipe() error:", err)
		return
	}
	if err := cmd.Start(); err != nil {
		fmt.Println("cmd.Start() error:", err)
		return
	}
	bytes, err := io.ReadAll(stdout)
	if err != nil {
		fmt.Println("io.ReadAll(stdout) error:", err)
		return
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println("cmd.Wait() error:", err.Error())
		return
	}
	fmt.Printf("stdout:\n%s", bytes)
}
