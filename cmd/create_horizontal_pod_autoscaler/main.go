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
	autoscalingv1 "k8s.io/api/autoscaling/v1"
)

func main() {


    const DEPLOYMENT_NAME string = "deployment-example-autoscale"

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

    # step 1 create deployment

	var deployment *appsv1.Deployment
	deployment = new(appsv1.Deployment)
	deployment.Name = "deployment-example-autoscale"
	var container v1.Container
	container.Name = DEPLOYMENT_NAME
	container.Image = "nginx"

	var containers []v1.Container
	containers = append(containers, container)

	deployment.Spec.Template.Spec.Containers = containers
	deployment.Spec.Template.ObjectMeta.Name = DEPLOYMENT_NAME

	metataDataLabels := make(map[string]string)
	metataDataLabels["run"] = DEPLOYMENT_NAME
	deployment.Spec.Template.ObjectMeta.Labels = metataDataLabels

	var labelSelector *metav1.LabelSelector
	labelSelector = new(metav1.LabelSelector)
	labelsMap := make(map[string]string)
	labelsMap["run"] = DEPLOYMENT_NAME

	fmt.Println(labelsMap)
	labelSelector.MatchLabels = labelsMap
	fmt.Println(labelSelector)
	deployment.Spec.Selector = labelSelector

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Create deployment
	deployments, err := clientset.AppsV1().Deployments("default").Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(deployments)


	# step 2 create HPA - horizontal pod autoscaling

	var autoscaler  *autoscalingv1.HorizontalPodAutoscaler
	autoscaler = new(autoscalingv1.HorizontalPodAutoscaler)

	autoscaler.Spec.MinReplicas == 2
	autoscaler.Spec.MaxReplicas == 4
	autoscaler.Spec.TargetCPUUtilizationPercentage =50
	autoscaler.ScaleTargetRef.Name = DEPLOYMENT_NAME
	autoscaler.ScaleTargetRef.Kind = "Deployment"

	// Create autoscaler
	autoscalers, err := clientset.AutoScalingV1().HorizontalPodAutoscaler("default").Create(context.TODO(), autoscaler, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}

}
