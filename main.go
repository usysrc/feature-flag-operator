package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/usysrc/feature-flag-operator/internal/ffo"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Initialize Kubernetes client
	config, err := clientcmd.BuildConfigFromFlags("", "")
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Configure Feature Flag Source
	flipt := ffo.NewFliptClient("localhost:3122")

	// Watch for relevant resources (e.g., Deployments, Pods)
	informerFactory := informers.NewSharedInformerFactory(clientset, time.Second*30)
	deploymentInformer := informerFactory.Apps().V1().Deployments()

	// Define reconciliation logic
	deploymentInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			deployment := obj.(*appsv1.Deployment)
			// Fetch feature flags from API
			flags := fetchFeatureFlags(flipt)
			// Inject feature flags into deployment
			injectFeatureFlags(deployment, flags)
			// Update deployment in Kubernetes cluster
			updateDeployment(clientset, deployment)
		},
	})

	// Start informers
	stopCh := make(chan struct{})
	informerFactory.Start(stopCh)
	<-stopCh
}

func fetchFeatureFlags(getter ffo.FlagGetter) ffo.FeatureFlags {
	flags, err := getter.FlagList()
	if err != nil {
		panic(err.Error())
	}
	return ffo.FeatureFlags{
		Flags: flags,
	}
}

func injectFeatureFlags(deployment *appsv1.Deployment, flags ffo.FeatureFlags) {
	// Logic to inject feature flags into Kubernetes resources
	// TODO: limit to certain annotations or implement more than just annotations
	for _, flag := range flags.Flags {
		deployment.Spec.Template.ObjectMeta.Annotations[flag.Name] = flag.Value
		deployment.Spec.Template.ObjectMeta.Annotations["enabled-"+flag.Name] = strconv.FormatBool(flag.Enabled)
	}
}

func updateDeployment(clientset *kubernetes.Clientset, deployment *appsv1.Deployment) {
	// Logic to update deployment in Kubernetes cluster
	_, err := clientset.AppsV1().Deployments(deployment.Namespace).Update(context.TODO(), deployment, metav1.UpdateOptions{})
	if err != nil {
		panic(fmt.Errorf("failed to update deployment: %v", err))
	}
	fmt.Println("Updated deployment:", deployment.Name)
}
