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

	// List configmaps
	configmaps, err := clientset.CoreV1().ConfigMaps("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for i := 1; i < len(configmaps.Items); i++ {
		configMap := configmaps.Items[i]
		fmt.Printf("configMap %s\n", configMap.ObjectMeta.Name)
		fmt.Printf("Data %s\n", configMap.Data)
		fmt.Printf("ResourceVersion %s\n", configMap.ResourceVersion)

		fmt.Printf("===========")
	}
}
