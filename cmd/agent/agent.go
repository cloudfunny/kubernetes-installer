package main

import (
	"github.com/cloudfunny/kubernetes-installer/pkg/agent"
	"github.com/cloudfunny/kubernetes-installer/pkg/kubeadmutils"
)

func main() {
	clusterInformer := agent.NewInformer("localhost:8080", "clusters")
	clusterInformer.Callback["create"] = func(i ...interface{}) {
		args := make(map[string]string)
		args["help"] = ""
		kubeadmCMD := kubeadmutils.NewKubeadmCommand(args)
		kubeadmCMD.Run()
	}

	go clusterInformer.Listen()

	nodeInformer := agent.NewInformer("localhost:8080", "nodes")
	go nodeInformer.Listen()

	select {}
}
