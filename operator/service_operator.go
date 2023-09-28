package operator

import (
	"context"
	"github.com/NoahAmethyst/simple-kube-operator/protocol/pb/kube_opt_pb"
	"github.com/NoahAmethyst/simple-kube-operator/utils/log"
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Services Get kubernetes services with namespace specified.
func Services(ctx context.Context, req *kube_opt_pb.KubeOptReq) (*v12.ServiceList, error) {
	if KubeCli.Err != nil {
		log.Error().Msgf("Kubernetes client has error:%s", KubeCli.Err)
		return nil, KubeCli.Err
	}

	return KubeCli.CoreV1().Services(req.Namespace).List(ctx, v1.ListOptions{})

}
