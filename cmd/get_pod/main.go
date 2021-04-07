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

	// Get pod with name abc in namespace default
	pod, err := clientset.CoreV1().Pods("default").Get(context.TODO(), "abc", metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Pod name %s\n", pod.ObjectMeta.Name)
	fmt.Printf("Namespace %s\n", pod.ObjectMeta.Namespace)
	fmt.Printf("ResourceVersion %s\n", pod.ResourceVersion)

	containers := pod.Spec.Containers
	fmt.Printf("Number of containers = %d\n", len(containers))

	fmt.Println("Containers")
	for j := range containers {
		fmt.Printf("    Container Name %s\n", containers[j].Name)
		fmt.Printf("    Container Image %s\n", containers[j].Image)

	}

	fmt.Printf("Labels %s\n", pod.Labels)

	fmt.Printf("Annotations %s\n", pod.Annotations)
	fmt.Println("==================")

}
