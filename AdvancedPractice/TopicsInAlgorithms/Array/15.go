package main

import (
	"fmt"
	"os/exec"
)

func main() {
	stdout := exec.Command("grep", "GOROOT")
	result, err := stdout.StdoutPipe()
	if err != nil {
		return
	}
	a1 := make([]byte, 1024)
	n, err := result.Read(a1)
	fmt.Printf("out:%s", a1[:n])

}
