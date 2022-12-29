package tools

import (
	"fmt"
	"os/exec"
)

func Command() {
	cmd := exec.Command("bash", "-c", "ls -rt")
	// cmd := exec.Command("echo", "Hi")
	out, _ := cmd.Output()
	fmt.Println(string(out))
}
