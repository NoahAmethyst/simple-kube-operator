package qqbot_svc

import (
	"github.com/NoahAmethyst/simple-kube-operator/cluster/rpc"
	"github.com/NoahAmethyst/simple-kube-operator/protocol/pb/qqbot_pb"
)

func SvcCli() qqbot_pb.QQBotServiceClient {
	conn := rpc.GetConn(rpc.CliQQBot)
	return qqbot_pb.NewQQBotServiceClient(conn)
}
