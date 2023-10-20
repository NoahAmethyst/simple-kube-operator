package constant

import "time"

const (
	DefaultGRPCPort = 9090

	GrpcListenPort = "GRPC_LISTEN_PORT"

	TimeOut = time.Second * 60 * 5

	K8sMasterUrl  = "K8S_MASTER_URL"
	K8sConfigFile = "K8S_CONFIG_FILE"

	InSecure = "INSECURE"

	NotifyAddr = "NOTIFY_ADDR"
)
