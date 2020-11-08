package main

import (
	"os"
	"strconv"
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/wh8199/log"
)

func main() {
	logging := log.NewLogging("main", log.INFO_LEVEL, 2)

	cfg, err := rest.InClusterConfig()
	if err != nil {
		logging.Fatalf("Error building kubecokubeClientnfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		logging.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	nsName := os.Getenv("NAMESPACE_NAME")
	podLabels := os.Getenv("POD_LABELS")
	retryCountEnv := os.Getenv("MAX_RETRY")

	retryCount := 5
	if temp, err := strconv.Atoi(retryCountEnv); err == nil {
		retryCount = temp
	}

	for i := 0; i < retryCount; i++ {
		podArr := make([]corev1.Pod, 0, 0)
		podList, err := kubeClient.CoreV1().Pods(nsName).
			List(metav1.ListOptions{
				LabelSelector: podLabels,
			})
		if err != nil {
			logging.Warnf("Fail to list all pods: %s", err.Error())
			continue
		}

		for _, pod := range podList.Items {
			if isPodRunning(&pod) {
				podArr = append(podArr, pod)
			}
		}

		if len(podArr) > 0 {
			os.Exit(0)
		}

		time.Sleep(2 * time.Second)
	}

	logging.Info("Dependency pods are not running, restart again")

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
