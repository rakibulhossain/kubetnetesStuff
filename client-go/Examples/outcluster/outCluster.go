package main

import (
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"path/filepath"
)
func main(){
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absoloute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig","","absolute path to the kubeconfig file")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("",*kubeconfig)
	if err!= nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	for {
		pods , err := clientset.CoreV1().Pods("").Watch(context.TODO(),metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}

		//fmt.Printf("There are %d pods in the cluster\n",len(pods.Items))
		ch := pods.ResultChan()
		for event := range ch {
			 pod, ok := event.Object.(*v1.Pod)
			 if !ok {
				 log.Fatal("unexpected type")
			 }
			 fmt.Println(pod.Name," =========> ",pod.Status.Phase)
		}

	}

}
