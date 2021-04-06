package main

import (
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func main() {

	// get kubeconfig file
	var kubeconfig *string
	homeDir := homedir.HomeDir()


	kubeconfig = flag.String("kubeconfig", filepath.Join(homeDir, ".kube", "config"), "(optional) absolute path to the kubeconfig file")

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	labelSelector := fmt.Sprintf("run=abc")
	listOptions := metav1.ListOptions{
		LabelSelector: labelSelector,
	}

	// List pods in `default` namespace
	watcher, err := clientset.CoreV1().Pods("default").Watch(context.TODO(), listOptions)
	if err != nil {
		panic(err.Error())
	}

	ch := watcher.ResultChan()

	for event := range ch {
		pod, ok := event.Object.(*v1.Pod)
		if !ok {
			fmt.Println("error")
		}
		fmt.Printf("pod name %s - \n", pod.GetName())
		fmt.Printf("resource version %s\n", pod.ResourceVersion)
		switch event.Type {
		case watch.Added:
			{
				fmt.Println("Pod added")
			}
		case watch.Deleted:
			{
				fmt.Println("Pod deleted")
			}
		case watch.Modified:
			{
				fmt.Println("Pod modified")
			}
		case watch.Bookmark:
			{
				fmt.Println("Pod bookmarked")
			}

		case watch.Error:
			{
				fmt.Println("watch error")
			}
		}
	}
}
