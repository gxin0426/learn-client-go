package main
import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)
func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "config")
	if err != nil {
		panic(err)
	}
	//fmt.Println()
	//fmt.Println(config)

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	podClient := clientset.CoreV1().Pods(corev1.NamespaceDefault)
	list, err := podClient.List(metav1.ListOptions{Limit:100})
	if err != nil {
		panic(err)
	}

	for _, d := range list.Items {
		fmt.Println(d)
	}


}
