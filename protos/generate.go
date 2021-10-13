//go:generate go run generate.go
package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	goPath = os.Getenv("GOPATH")
)

func main() {
	args := []string{
		"-I=.",
		fmt.Sprintf("-I=%s", filepath.Join(goPath, "src")),
		fmt.Sprintf("-I=%s", filepath.Join(goPath, "src", "github.com")),
		fmt.Sprintf("-I=%s", filepath.Join(goPath, "src", "github.com", "googleapis", "googleapis")),
		fmt.Sprintf("-I=%s", filepath.Join("/usr", "local", "include")),
		"--go_out=go",
		"--include_imports",
		"document.proto",
		"genesis.proto",
		"key.proto",
		"attestation.proto",
		"didar.proto",
	}

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("protoc", args...)
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	} else {
		fmt.Println("successfully compiled protocol buffers")
	}
}
