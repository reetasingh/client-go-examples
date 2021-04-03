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

	// List Secret
	deployments, err := clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for i := 1; i < len(deployments.Items); i++ {
		deployment := deployments.Items[i]
		fmt.Printf("Deployment name %s\n", deployment.ObjectMeta.Name)
		fmt.Printf("Namespace %s\n", deployment.ObjectMeta.Namespace)
		fmt.Printf("ResourceVersion %s\n", deployment.ResourceVersion)
		fmt.Printf("Labels %s\n", deployment.Labels)
		fmt.Printf("Annotations %s\n", deployment.Annotations)
		fmt.Printf("===========")
	}
}
