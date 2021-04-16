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
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
)

func main() {

	// get kubeconfig file
	var kubeconfig *string
	homeDir := homedir.HomeDir()
	fmt.Println(homeDir)


	kubeconfig = flag.String("kubeconfig", filepath.Join(homeDir, ".kube", "config"), "(optional) absolute path to the kubeconfig file")

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	var job *batchv1.Job
	job = new(batchv1.Job)
	job.Name = "my-first-k8s-pod"
	var container v1.Container
	container.Name = "my-first-pod-container"
	container.Image = "nginx"

	var containers []v1.Container
	containers = append(containers, container)

	job.Spec.Template.Spec.Containers = containers
	job.Spec.Template.Spec.RestartPolicy = "OnFailure"

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Create pod
	jobs, err := clientset.BatchV1().Jobs("default").Create(context.TODO(), job, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(jobs)

}
