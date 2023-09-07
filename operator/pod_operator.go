package operator

import (
	"context"
	"github.com/NoahAmethyst/simple-kube-operator/protocol/pb/kube_opt_pb"
	"github.com/NoahAmethyst/simple-kube-operator/utils/log"

	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetPods Get pods information
func GetPods(ctx context.Context, req *kube_opt_pb.KubeOptReq) (*v12.PodList, error) {
	if KubeCli.Err != nil {
		log.Error().Msgf("Kubernetes client has error:%s", KubeCli.Err)
		return nil, KubeCli.Err
	}

	return KubeCli.CoreV1().Pods(req.Namespace).List(ctx, v1.ListOptions{})

}

func DelPod(ctx context.Context, req *kube_opt_pb.KubeOptReq) error {
	if KubeCli.Err != nil {
		log.Error().Msgf("Kubernetes client has error:%s", KubeCli.Err)
		return nil
	}

	return KubeCli.CoreV1().Pods(req.Namespace).Delete(ctx, req.PodId, v1.DeleteOptions{})
}
