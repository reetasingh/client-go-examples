package main

import (
    "k8s.io/client-go/util/homedir"
    "fmt"
    "context"
    "flag"
    "path/filepath"
    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/kubernetes"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	// get pods
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

    // print details
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	for i := 1; i < len(pods.Items); i++ {
	    pod := pods.Items[i]
        fmt.Printf("pod name %s\n", pod.ObjectMeta.Name)

        containers := pod.PodSpec.Containers
        fmt.Printf("Container name %s\n", len(containers))
	}
}