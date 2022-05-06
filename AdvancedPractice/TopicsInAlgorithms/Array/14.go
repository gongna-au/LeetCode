package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	//1
	cmd_go_env := exec.Command("go", "env")
	cmd_grep := exec.Command("grep", "GOROOT")
	//2 准备好输出
	A, _ := cmd_go_env.StdoutPipe()
	// 3
	cmd_go_env.Start()
	pipe := bufio.NewReader(A)

	//2 准备好拿取
	B, _ := cmd_grep.StdinPipe()
	pipe.WriteTo(B)
	var bufResult bytes.Buffer

	cmd_grep.Stdout = &bufResult

	cmd_grep.Start()
	B.Close()
	cmd_grep.Wait()
	fmt.Println(bufResult.String())

}
