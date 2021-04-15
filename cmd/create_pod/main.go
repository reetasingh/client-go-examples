package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	v1 "k8s.io/api/core/v1"
)

func main() {

	// get kubeconfig file
	var kubeconfig *string
	homeDir := homedir.HomeDir()
	fmt.Println(homeDir)

	fmt.Println(filepath.Join(homeDir, ".kube", "config"))

	kubeconfig = flag.String("kubeconfig", filepath.Join(homeDir, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	fmt.Println(kubeconfig)

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	var pod *v1.Pod
	pod = new(v1.Pod)
	pod.Name = "my-first-pod"
	var container v1.Container
	container.Name = "my-first-pod-container"
	container.Image = "nginx"

	var containers []v1.Container
	containers = append(containers, container)

	pod.Spec.Containers = containers

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Create pod
	pods, err := clientset.CoreV1().Pods("default").Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(pods)

}
