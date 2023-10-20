package notifier

import "context"

type Notifier interface {
	notifyPodModified(context.Context, PodModified)
	generatePodModifiedContent(PodModified) string
}

type PodModifiedType int

const (
	None PodModifiedType = iota
	PodCreating
	PodDeleted
	PodCreated
)

type PodModified struct {
	PodName string
	App     string
	Status  PodModifiedType
}

func NotifyPodModified(ctx context.Context, podModified PodModified) {
	//Todo Automatic register notifier and use it
	notifier := new(QQNotifier)

	notifier.notifyPodModified(ctx, podModified)
}
