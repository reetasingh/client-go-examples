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

	// List namespaces
	daemonsets, err := clientset.AppsV1().DaemonSets("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for i := 0; i < len(daemonsets.Items); i++ {
		daemonset := daemonsets.Items[i]
		fmt.Printf("DaemonSet name %s\n", daemonset.Name)
		fmt.Printf("Namespace %s\n", daemonset.ObjectMeta.Namespace)
		fmt.Printf("Selector %s\n", daemonset.Spec.Selector)
		fmt.Printf("Pod Spec %s\n", daemonset.Spec.Template.Spec)
		fmt.Printf("ResourceVersion %s\n", daemonset.ResourceVersion)
		fmt.Printf("Labels %s\n", daemonset.Labels)
		fmt.Printf("Annotations %s\n", daemonset.Annotations)
		fmt.Printf("===========")
	}
}
