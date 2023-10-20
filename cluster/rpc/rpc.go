package rpc

import (
	"github.com/NoahAmethyst/simple-kube-operator/utils/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"os"
	"time"
)

type GRPCClientName string

const (
	CliQQBot GRPCClientName = "QQBOT"
)

var conn map[GRPCClientName]*grpc.ClientConn

var RpcCliList []GRPCClientName

func init() {
	RpcCliList = []GRPCClientName{CliQQBot}
	conn = make(map[GRPCClientName]*grpc.ClientConn)
}

func GetConn(clientName GRPCClientName) *grpc.ClientConn {
	return conn[clientName]
}

func setConn(clientName GRPCClientName, c *grpc.ClientConn) {
	conn[clientName] = c
}

func InitGrpcCli(clientName GRPCClientName) {
	addr := os.Getenv(string(clientName) + "_SERVICE_ADDR")

	if len(addr) == 0 {
		log.Error().Msgf("Empty svc addr %s", string(clientName))
	}

	grpcConn, err := startConnection(addr)

	if err != nil {
		log.Error().Msgf("InitGrpcCli grpc client err %s", err.Error())
	} else {
		log.Info().Msgf("InitGrpcCli grpc client success at %s", addr)
		setConn(clientName, grpcConn)
	}
}

func startConnection(address string) (*grpc.ClientConn, error) {
	kacp := keepalive.ClientParameters{
		Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
		Timeout:             8 * time.Second,  // wait 1 second for ping ack before considering the connection dead
		PermitWithoutStream: true,             // send pings even without active streams
	}

	return grpc.Dial(address, grpc.WithKeepaliveParams(kacp),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

}
