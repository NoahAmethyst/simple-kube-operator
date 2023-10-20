package main

import (
	"context"
	"github.com/NoahAmethyst/simple-kube-operator/cluster"
	"github.com/NoahAmethyst/simple-kube-operator/cluster/rpc"
	"github.com/NoahAmethyst/simple-kube-operator/constant"
	"github.com/NoahAmethyst/simple-kube-operator/operator"
	"os"
	"time"
)

func main() {

	//Set time location to East eighth District
	time.Local = time.FixedZone("UTC", 8*60*60)

	ctx := context.Background()

	gracefulShutdown(ctx, cluster.KubeOptServer)
	// Start event watcher
	operator.MonitoringPod(ctx)

	// Initialize grpc client
	for _, rpcCli := range rpc.RpcCliList {
		rpc.InitGrpcCli(rpcCli)
	}

	// Start Grpc server
	cluster.StartServer(os.Getenv(constant.GrpcListenPort))
}
