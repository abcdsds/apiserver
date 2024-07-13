package repository

import (
	"errors"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	listerv1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"time"
)

type Kubernetes struct {
	client listerv1.PodLister
}

func NewKubernetes() (*Kubernetes, error) {
	cfg := &rest.Config{}

	client := kubernetes.NewForConfigOrDie(cfg)

	factory := informers.NewSharedInformerFactory(client, 600*time.Second)
	informer := factory.Core().V1().Pods().Informer()
	stopCh := make(chan struct{})
	defer close(stopCh)
	go factory.Start(stopCh)

	if !cache.WaitForCacheSync(stopCh, informer.HasSynced) {
		return nil, errors.New("not synced pod informer")
	}

	return &Kubernetes{
		client: factory.Core().V1().Pods().Lister(),
	}, nil
}

func (receiver *Kubernetes) Pods(ns string) ([]*v1.Pod, error) {
	return receiver.client.Pods(ns).List(labels.Everything())
}
