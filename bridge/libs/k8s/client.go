package k8s

import (
	xerrors "github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

var (
	ClientSet *kubernetes.Clientset
)

func InitK8SClient() error {
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return xerrors.WithMessagef(err, "init client config failed")
	}

	ClientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		return xerrors.WithMessagef(err, "init client failed")
	}

	return nil
}
