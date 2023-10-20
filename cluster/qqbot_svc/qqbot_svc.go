package qqbot_svc

import (
	"github.com/NoahAmethyst/simple-kube-operator/cluster/rpc"
	"github.com/NoahAmethyst/simple-kube-operator/protocol/pb/qqbot_pb"
)

func SvcCli() qqbot_pb.QQBotServiceClient {
	return qqbot_pb.NewQQBotServiceClient(rpc.GetConn(rpc.CliQQBot))
}
