package operator

import (
	"context"
	"github.com/NoahAmethyst/simple-kube-operator/notifier"
	"github.com/NoahAmethyst/simple-kube-operator/utils/log"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
	"reflect"
	"sync"
)

var existChan = make(chan struct{})
var MonitorShutdown = make(chan struct{})

var once sync.Once

// Restart watcher when it exists.
func daemons(ctx context.Context) {
	log.Info().Msgf("Start daemon of kubernetes events watcher")
	go func() {
		for {
			select {
			case _, ok := <-existChan:
				if ok {
					log.Info().Msgf("Restart watcher.")
					MonitoringPod(ctx)
				} else {
					log.Warn().Msgf("Closed existChan,watcher daemon process exist.")
					return
				}
			}
		}
	}()

}

func MonitoringPod(ctx context.Context) {
	if KubeCli.Err != nil {
		log.Error().Msgf("Kubernetes client has error:%s", KubeCli.Err.Error())
		return
	}
	lwc := cache.NewListWatchFromClient(KubeCli.Clientset.CoreV1().RESTClient(), "pods", "default", fields.Everything())
	watcher, err := lwc.Watch(v12.ListOptions{
		TypeMeta: v12.TypeMeta{
			Kind: "Pod",
		},
	})
	if err != nil {
		log.Error().Msgf("Make watcher failed:%s", err.Error())
		return
	}

	go func(_ctx context.Context, _watcher watch.Interface) {
		defer func() {
			log.Warn().Msgf("Exist event monitoring")
			existChan <- struct{}{}
		}()

		// Start daemons to restart monitor when watcher closed.
		once.Do(func() {
			daemons(_ctx)
		})

		for {
			select {
			case _event, ok := <-_watcher.ResultChan():
				if ok {
					EventHandler(_event)
				} else {
					log.Warn().Msgf("Event watcher channel is closed")
					// Release resource used by watcher.
					_watcher.Stop()
					return
				}
			case _, ok := <-MonitorShutdown:
				if ok {
					log.Info().Msgf("Graceful shutdown,release resource used by event watcher")
					_watcher.Stop()
					return
				}

			}
		}
	}(ctx, watcher)
}

func EventHandler(event watch.Event) {
	if pod, ok := event.Object.(*v1.Pod); !ok {
		log.Warn().Msgf("event type [%s] not Pod", reflect.TypeOf(event.Object).Name())
		return
	} else {
		log.Debug().Msgf("Event:%s  App:%s  PodName:%s  Status:%s", event.Type, pod.Labels["app"], pod.Name, pod.Status.Phase)
		podModified := notifier.PodModified{
			PodName: pod.Name,
			App:     pod.Labels["app"],
			Status:  notifier.None,
		}
		switch event.Type {

		case watch.Added:

			switch pod.Status.Phase {
			// When event type is ADDED and status is Pending,the pod is under creating.
			case v1.PodPending:
				log.Info().Msgf(" App:%s  PodName:%s is under creating", pod.Labels["app"], pod.Name)

			default:

			}
		case watch.Deleted:
			switch pod.Status.Phase {
			// When event type is DELETED and status is Succeeded,the pod is deleted.
			case v1.PodSucceeded:
				log.Info().Msgf(" App:%s  PodName:%s is deleted", pod.Labels["app"], pod.Name)
				podModified.Status = notifier.PodDeleted
			default:

			}

		case watch.Modified:
			switch pod.Status.Phase {
			// When event type is MODIFIED and status is Running,the pod is created successful and running.
			case v1.PodRunning:
				log.Info().Msgf(" App:%s  PodName:%s is created successful and running", pod.Labels["app"], pod.Name)
				podModified.Status = notifier.PodCreated
			default:

			}
		}

		//Notify
		if podModified.Status != notifier.None {
			notifier.NotifyPodModified(context.Background(), podModified)
		}

	}

}
