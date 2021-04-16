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
	appsv1 "k8s.io/api/apps/v1"
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

	var deployment *appsv1.Deployment
	deployment = new(appsv1.Deployment)
	deployment.Name = "my-first-k8s-deployment"
	var container v1.Container
	container.Name = "my-first-pod-container"
	container.Image = "nginx"

	var containers []v1.Container
	containers = append(containers, container)

	deployment.Spec.Template.Spec.Containers = containers
	deployment.Spec.Template.ObjectMeta.Name = "my-first-k8s-deployment"

	metataDataLabels := make(map[string]string)
	metataDataLabels["run"] = "my-first-k8s-deployment"
	deployment.Spec.Template.ObjectMeta.Labels = metataDataLabels


	var labelSelector *metav1.LabelSelector
	labelSelector = new(metav1.LabelSelector)
	labelsMap := make(map[string]string)
	labelsMap["run"] = "my-first-k8s-deployment"

	fmt.Println(labelsMap)
	labelSelector.MatchLabels = labelsMap
	fmt.Println(labelSelector)
	deployment.Spec.Selector = labelSelector


	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Create pod
	deployments, err := clientset.AppsV1().Deployments("default").Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(deployments)

}
