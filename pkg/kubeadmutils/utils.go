package kubeadmutils

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func ListDirectories() string {
	var output bytes.Buffer
	commandPath := "ipconfig"
	// args := []string{""}
	cmd := exec.Command(commandPath)
	cmd.Stdin = strings.NewReader("some input")
	cmd.Stdout = &output
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("----------", output.String())
	return output.String()
}
