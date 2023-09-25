package operator

import (
	"context"
	"github.com/NoahAmethyst/simple-kube-operator/constant"
	"github.com/NoahAmethyst/simple-kube-operator/utils/log"
	v13 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

var KubeCli *K8sClient

type K8sClient struct {
	*kubernetes.Clientset
	Err error
}

func init() {
	KubeCli = new(K8sClient)
	//Confirm the kubernetes config file is exist before call the function.Use '/etc/kubernetes/admin.conf' if not exist
	configFile := os.Getenv(constant.K8sConfigFile)
	if len(configFile) == 0 {
		configFile = "/etc/kubernetes/admin.conf"
	}

	cfg, err := clientcmd.BuildConfigFromFlags(os.Getenv(constant.K8sMasterUrl), configFile)
	if err != nil {
		log.Error().Msgf("Build Kubernetes config failed:%s", err.Error())
		KubeCli.Err = err
		return
	}

	// Use insecure for remote call
	// If you need reset client config use ResetCli
	cfg.Insecure = true

	if k8sClient, err := kubernetes.NewForConfig(cfg); err != nil {
		KubeCli.Err = err
	} else {
		KubeCli.Clientset = k8sClient
	}
}

// ResetCli Reset kubernetes client
// Reset not worked when error happened and throw error
func ResetCli(masterUrl, configFile string, insecure bool) error {
	cfg, err := clientcmd.BuildConfigFromFlags(masterUrl, configFile)
	if err != nil {
		return err
	}

	cfg.Insecure = insecure

	if k8sClient, err := kubernetes.NewForConfig(cfg); err != nil {
		return err
	} else {
		KubeCli.Clientset = k8sClient
	}
	return nil
}

func Namespaces(ctx context.Context) (*v13.NamespaceList, error) {
	if KubeCli.Err != nil {
		log.Error().Msgf("Kubernetes client has error:%s", KubeCli.Err)
		return nil, KubeCli.Err
	}

	return KubeCli.CoreV1().Namespaces().List(ctx, v12.ListOptions{})

}
