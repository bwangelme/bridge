package deploycon

import (
	"bridge/bridge/libs/k8s"
	"context"
	xerrors "github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Pod struct {
	Name string
}

func GetPods(ctx context.Context, namespace string) ([]*Pod, error) {
	opts := metav1.ListOptions{}

	podList, err := k8s.ClientSet.CoreV1().Pods(namespace).List(ctx, opts)
	if err != nil {
		return nil, xerrors.WithMessagef(err, "list failed")
	}

	var res = make([]*Pod, 0)
	for _, pod := range podList.Items {
		res = append(res, &Pod{
			Name: pod.Name,
		})
	}

	return res, nil
}
