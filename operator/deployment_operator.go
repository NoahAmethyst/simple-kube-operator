package operator

import (
	"context"
	"github.com/NoahAmethyst/simple-kube-operator/protocol/pb/kube_opt_pb"
	"github.com/NoahAmethyst/simple-kube-operator/utils/log"
	v13 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Deployments(ctx context.Context, req *kube_opt_pb.KubeOptReq) (*v13.DeploymentList, error) {
	if KubeCli.Err != nil {
		log.Error().Msgf("Kubernetes client has error:%s", KubeCli.Err.Error())
		return nil, KubeCli.Err
	}

	return KubeCli.AppsV1().Deployments(req.Namespace).List(ctx, v1.ListOptions{})

}
