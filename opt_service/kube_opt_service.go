package opt_service

import (
	"context"
	"github.com/NoahAmethyst/simple-kube-operator/constant"
	"github.com/NoahAmethyst/simple-kube-operator/operator"
	"github.com/NoahAmethyst/simple-kube-operator/protocol/pb/kube_opt_pb"
)

type KubeOptServer struct{}

func (s KubeOptServer) GetPods(ctx context.Context, req *kube_opt_pb.KubeOptReq) (resp *kube_opt_pb.KubeOptResp, err error) {
	resp = new(kube_opt_pb.KubeOptResp)
	if len(req.Namespace) == 0 {
		req.Namespace = constant.Default
	}

	if list, err := operator.GetPods(ctx, req); err != nil {
		resp.Code = constant.Failed
		resp.Message = err.Error()
	} else {
		for _, item := range list.Items {
			resp.Pods = append(resp.Pods, &kube_opt_pb.KubePod{
				Namespace: item.Namespace,
				App:       item.Labels["app"],
				PodId:     item.Name,
				Status:    string(item.Status.Phase),
			})
		}
	}
	return
}

func (s KubeOptServer) GetServices(_ context.Context, _ *kube_opt_pb.KubeOptReq) (resp *kube_opt_pb.KubeOptResp, err error) {
	return
}

func (s KubeOptServer) GetDeployments(_ context.Context, _ *kube_opt_pb.KubeOptReq) (resp *kube_opt_pb.KubeOptResp, err error) {
	return
}

func (s KubeOptServer) DeletePod(ctx context.Context, req *kube_opt_pb.KubeOptReq) (resp *kube_opt_pb.KubeOptResp, err error) {
	resp = new(kube_opt_pb.KubeOptResp)
	if len(req.Namespace) == 0 {
		req.Namespace = constant.Default
	}
	if err := operator.DelPod(ctx, req); err != nil {
		resp.Code = constant.Failed
		resp.Message = err.Error()
	} else {
		resp.Code = constant.Success
	}
	return

}
