package main

import (
	"os"
	"strings"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/klog"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	cfg, err := rest.InClusterConfig()
	if err != nil {
		klog.Fatalf("Error building kubecokubeClientnfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	nsName := os.Getenv("NAMESPACE_NAME")
	labels := strings.Split(os.Getenv("POD_LABELS"), ",")
	klog.Info(labels)

	podArr := make([]corev1.Pod, 0, 0)

	for _, label := range labels {
		podList, err := kubeClient.CoreV1().Pods(nsName).
			List(metav1.ListOptions{
				LabelSelector: label,
			})
		if err != nil {
			klog.Fatalf("Fail to list all pods: %s", err.Error())
		}

		for _, pod := range podList.Items {
			if isPodRunning(&pod) {
				podArr = append(podArr, pod)
			}
		}
	}

	if len(podArr) > 0 {
		os.Exit(0)
	}

	os.Exit(1)
}

func isPodRunning(pod *corev1.Pod) bool {
	for _, v := range pod.Status.ContainerStatuses {
		if v.State.Terminated != nil || v.State.Waiting != nil {
			return false
		}
	}
	return true
}
