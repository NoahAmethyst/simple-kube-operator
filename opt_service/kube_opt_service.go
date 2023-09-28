package opt_service

import (
	"context"
	"github.com/NoahAmethyst/simple-kube-operator/constant"
	"github.com/NoahAmethyst/simple-kube-operator/operator"
	"github.com/NoahAmethyst/simple-kube-operator/protocol/pb/kube_opt_pb"
)

type KubeOptServer struct{}

func (s KubeOptServer) Namespaces(ctx context.Context, req *kube_opt_pb.KubeOptReq) (resp *kube_opt_pb.KubeOptResp, err error) {
	resp = new(kube_opt_pb.KubeOptResp)
	if namespaces, err := operator.Namespaces(ctx); err != nil {
		resp.Code = constant.Failed
		resp.Message = err.Error()
	} else {
		for _, item := range namespaces.Items {
			resp.Namespaces = append(resp.Namespaces, &kube_opt_pb.KubeNamespace{Namespace: item.Name})
		}
	}
	return
}

func (s KubeOptServer) Pods(ctx context.Context, req *kube_opt_pb.KubeOptReq) (resp *kube_opt_pb.KubeOptResp, err error) {
	resp = new(kube_opt_pb.KubeOptResp)
	if len(req.Namespace) == 0 {
		req.Namespace = constant.Default
	}

	if list, err := operator.Pods(ctx, req); err != nil {
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

func (s KubeOptServer) Services(ctx context.Context, req *kube_opt_pb.KubeOptReq) (resp *kube_opt_pb.KubeOptResp, err error) {

	resp = new(kube_opt_pb.KubeOptResp)
	if len(req.Namespace) == 0 {
		req.Namespace = constant.Default
	}

	if list, err := operator.Services(ctx, req); err != nil {
		resp.Code = constant.Failed
		resp.Message = err.Error()
	} else {
		for _, item := range list.Items {
			ports := make([]*kube_opt_pb.Ports, 0, len(item.Spec.Ports))
			for _, port := range item.Spec.Ports {
				ports = append(ports, &kube_opt_pb.Ports{
					Protocol: string(port.Protocol),
					Port:     port.Port,
					NodePort: port.NodePort,
				})
			}
			resp.Services = append(resp.Services, &kube_opt_pb.KubeService{
				Namespace:  item.Namespace,
				Name:       item.ObjectMeta.Name,
				PortType:   string(item.Spec.Type),
				ClusterIps: item.Spec.ClusterIPs,
				Ports:      ports,
			})
		}
	}
	return
}

func (s KubeOptServer) Deployments(ctx context.Context, req *kube_opt_pb.KubeOptReq) (resp *kube_opt_pb.KubeOptResp, err error) {
	resp = new(kube_opt_pb.KubeOptResp)
	if len(req.Namespace) == 0 {
		req.Namespace = constant.Default
	}

	if list, err := operator.Deployments(ctx, req); err != nil {
		resp.Code = constant.Failed
		resp.Message = err.Error()
	} else {
		for _, item := range list.Items {
			var replicas int32
			if item.Spec.Replicas != nil {
				replicas = *item.Spec.Replicas
			}

			imagePullSecrets := make([]string, 0, len(item.Spec.Template.Spec.ImagePullSecrets))
			for _, secret := range item.Spec.Template.Spec.ImagePullSecrets {
				imagePullSecrets = append(imagePullSecrets, secret.Name)
			}
			resp.Deployments = append(resp.Deployments, &kube_opt_pb.KubeDeployment{
				Namespace:        item.Namespace,
				Name:             item.Name,
				Replicas:         replicas,
				Labels:           item.Labels,
				ImagePullSecrets: imagePullSecrets,
			})
		}
	}
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
