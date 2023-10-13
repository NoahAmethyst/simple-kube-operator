package middleware

import (
	"context"
	"github.com/NoahAmethyst/simple-kube-operator/utils/log"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"strconv"
	"strings"
	"time"
)

func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	var requestId string
	if _uuid, err := uuid.NewUUID(); err != nil {
		requestId = strconv.FormatInt(time.Now().UnixMilli(), 10)
	} else {
		requestId = strings.ReplaceAll(_uuid.String(), "-", "")
	}

	log.Info().Msgf("[%s] Receive grpc request: Method [%s] Body [%+v]", requestId, info.FullMethod, info.Server)
	resp, err := handler(ctx, req)
	if err != nil {
		log.Warn().Msgf("[%s] Grpc response Method [%s]  failed: %s", requestId, info.FullMethod, err.Error())
	} else {
		log.Info().Msgf("[%s] Grpc response Method [%s] success: %+v", requestId, info.FullMethod, resp)
	}
	return resp, err
}
