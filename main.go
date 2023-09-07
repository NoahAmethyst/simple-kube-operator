package main

import (
	"context"
	"github.com/NoahAmethyst/simple-kube-operator/cluster"
	"github.com/NoahAmethyst/simple-kube-operator/constant"
	"os"
	"time"
)

func main() {

	//Set time location to East eighth District
	time.Local = time.FixedZone("UTC", 8*60*60)

	ctx := context.Background()

	gracefulShutdown(ctx, cluster.KubeOptServer)

	cluster.StartServer(os.Getenv(constant.GrpcListenPort))
}
