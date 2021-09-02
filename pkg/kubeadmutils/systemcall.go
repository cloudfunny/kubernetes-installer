package kubeadmutils

import (
	"bytes"
	"fmt"
	"os/exec"
)

const defaultKubeadmCMD string = "/usr/bin/kubeadm"

type KubeadmCommand struct {
	CMD exec.Cmd
}

type Args map[string]string

func NewKubeadmCommand(argsMap map[string]string) *KubeadmCommand {
	var args []string
	for key, value := range argsMap {
		command := "--" + key + " " + value
		args = append(args, command)
	}
	return &KubeadmCommand{
		CMD: exec.Cmd{
			Path: defaultKubeadmCMD,
			Args: args,
		},
	}
}

func (kubeadm *KubeadmCommand) Run() {
	var stdout, stderr bytes.Buffer
	kubeadm.CMD.Stdout = &stdout
	kubeadm.CMD.Stderr = &stderr
	if err := kubeadm.CMD.Run(); err != nil {
		panic(err)
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Println(outStr, errStr)
}
