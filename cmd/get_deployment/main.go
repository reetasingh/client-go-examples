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

	// Get deployment with name nginx-app in namespace default
	deployment, err := clientset.AppsV1().Deployments("default").Get(context.TODO(), "nginx-app", metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Pod name %s\n", deployment.ObjectMeta.Name)
	fmt.Printf("Namespace %s\n", deployment.ObjectMeta.Namespace)
	fmt.Printf("ResourceVersion %s\n", deployment.ResourceVersion)

	fmt.Printf("Labels %s\n", deployment.Labels)

	fmt.Printf("Annotations %s\n", deployment.Annotations)

	fmt.Printf("Spec %s\n", deployment.Spec)
	fmt.Println("==================")

}
