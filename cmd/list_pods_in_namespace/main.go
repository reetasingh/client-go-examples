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
	fmt.Println(homeDir)

	fmt.Println(filepath.Join(homeDir, ".kube", "config"))

	kubeconfig = flag.String("kubeconfig", filepath.Join(homeDir, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	fmt.Println(kubeconfig)

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

	// List pods in `default` namespace
	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	// print details of pod
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	for i := 0; i < len(pods.Items); i++ {
		pod := pods.Items[i]
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
}