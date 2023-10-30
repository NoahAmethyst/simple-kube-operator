package notifier

import "context"

type Notifier interface {
	notifyPodModified(context.Context, PodState)
	generatePodModifiedContent(PodState) string
}

type State int

const (
	None State = iota
	PodCreating
	PodDeleted
	PodCreated
)

func NotifyPodModified(ctx context.Context, podModified PodState) {
	//Todo Automatic register notifier and use it
	notifier := new(QQNotifier)

	notifier.notifyPodModified(ctx, podModified)
}
