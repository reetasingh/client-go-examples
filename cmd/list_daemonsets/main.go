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

	// List namespaces
	daemonsets, err := clientset.AppsV1().DaemonSets().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for i := 0; i < len(daemonsets.Items); i++ {
		daemonset := daemonsets.Items[i]
		fmt.Printf("Namespace name %s\n", daemonset.ObjectMeta.Name)
		fmt.Printf("Status %s\n", daemonset.Status)
		fmt.Printf("ResourceVersion %s\n", daemonset.ResourceVersion)
		fmt.Printf("Labels %s\n", daemonset.Labels)
		fmt.Printf("Annotations %s\n", daemonset.Annotations)
		fmt.Printf("===========")
	}
}
