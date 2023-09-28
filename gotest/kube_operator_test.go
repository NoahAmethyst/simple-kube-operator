package gotest

import (
	"context"
	"github.com/NoahAmethyst/simple-kube-operator/protocol/pb/kube_opt_pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"testing"
	"time"
)

var ctx = context.Background()

// Set your grpc server address
var addr = "localhost:9091"

var keepAliveCfg = keepalive.ClientParameters{
	Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
	Timeout:             30 * time.Second, // wait 1 second for ping ack before considering the connection dead
	PermitWithoutStream: true,             // send pings even without active streams
}

func Test_GetPods(t *testing.T) {
	conn, err := grpc.Dial(addr, grpc.WithKeepaliveParams(keepAliveCfg), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	kubeOptCli := kube_opt_pb.NewKubeOptServiceClient(conn)

	defer func(conn *grpc.ClientConn) {
		if err := conn.Close(); err != nil {
			t.Error(err)
		}
	}(conn)

	resp, err := kubeOptCli.Pods(ctx, &kube_opt_pb.KubeOptReq{})
	if err != nil {
		panic(err)
	}
	if len(resp.Message) != 0 {
		panic(resp.Message)
	}

	for _, _pod := range resp.Pods {
		t.Logf("%+v", *_pod)
	}
}

func Test_Namespaces(t *testing.T) {
	conn, err := grpc.Dial(addr, grpc.WithKeepaliveParams(keepAliveCfg), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	kubeOptCli := kube_opt_pb.NewKubeOptServiceClient(conn)

	defer func(conn *grpc.ClientConn) {
		if err := conn.Close(); err != nil {
			t.Error(err)
		}
	}(conn)
	resp, err := kubeOptCli.Namespaces(ctx, &kube_opt_pb.KubeOptReq{})
	if err != nil {
		panic(err)
	}
	if len(resp.Message) != 0 {
		panic(resp.Message)
	}

	for _, _namespace := range resp.Namespaces {
		t.Logf("%+v", _namespace)
	}
}

func Test_GetServices(t *testing.T) {
	conn, err := grpc.Dial(addr, grpc.WithKeepaliveParams(keepAliveCfg), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	kubeOptCli := kube_opt_pb.NewKubeOptServiceClient(conn)

	defer func(conn *grpc.ClientConn) {
		if err := conn.Close(); err != nil {
			t.Error(err)
		}
	}(conn)
	resp, err := kubeOptCli.Services(ctx, &kube_opt_pb.KubeOptReq{})
	if err != nil {
		panic(err)
	}
	if len(resp.Message) != 0 {
		panic(resp.Message)
	}

	for _, service := range resp.Services {
		t.Logf("%+v", service)
	}
}

func Test_GetDeployments(t *testing.T) {
	conn, err := grpc.Dial(addr, grpc.WithKeepaliveParams(keepAliveCfg), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	kubeOptCli := kube_opt_pb.NewKubeOptServiceClient(conn)

	defer func(conn *grpc.ClientConn) {
		if err := conn.Close(); err != nil {
			t.Error(err)
		}
	}(conn)
	resp, err := kubeOptCli.Deployments(ctx, &kube_opt_pb.KubeOptReq{})
	if err != nil {
		panic(err)
	}
	if len(resp.Message) != 0 {
		panic(resp.Message)
	}

	for _, deployment := range resp.Deployments {
		t.Logf("%+v", deployment)
	}
}
