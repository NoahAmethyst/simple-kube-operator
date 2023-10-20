package notifier

import (
	"context"
	"fmt"
	"github.com/NoahAmethyst/simple-kube-operator/cluster/qqbot_svc"
	"github.com/NoahAmethyst/simple-kube-operator/protocol/pb/qqbot_pb"
	"github.com/NoahAmethyst/simple-kube-operator/utils/log"
)

type QQNotifier struct {
}

func (q *QQNotifier) notifyPodModified(ctx context.Context, content PodModified) {
	defer func() {
		if err := recover(); err != nil {
			log.Error().Msgf("Call QQNotifier panic:%+v", err)
		}
	}()
	cli := qqbot_svc.SvcCli()
	selfResp, err := cli.Self(ctx, new(qqbot_pb.Empty))
	if err != nil {
		log.Error().Msgf("Get qq bot info failed:%s", err.Error())
		return
	} else if len(selfResp.GetMessage()) > 0 {
		log.Error().Msgf("Get qq bot info failed:%s", selfResp.GetMessage())
		return
	}

	if selfResp.GetSelf().GetOwner() > 0 {
		text := q.generatePodModifiedContent(content)

		if sendMsgResp, err := cli.SendMsg(ctx, &qqbot_pb.SendMsgReq{
			Content: text,
			Chat:    selfResp.GetSelf().GetOwner(),
			Group:   false,
		}); err != nil {
			log.Error().Msgf("notifyPodModified qq bot failed:%s", err.Error())
		} else if len(sendMsgResp.GetMessage()) > 0 {
			log.Error().Msgf("notifyPodModified qq bot failed:%s", sendMsgResp.GetMessage())
		}
	} else {
		log.Warn().Msgf("QQ bot not set owner,can't notifyPodModified.")
	}
}

func (q *QQNotifier) generatePodModifiedContent(content PodModified) string {
	var status string
	switch content.Status {
	case PodDeleted:
		status = "停止"
	case PodCreated:
		status = "创建成功"
	default:
		return ""
	}

	return fmt.Sprintf("Kubernetes 容器状态变更：\n"+
		"App：%s\n"+
		"状态：%s", content.App, status)
}
