package main

import (
	"context"
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
)

func int32Ptr(x int32) *int32{
	return &x
}

func main(){

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig","", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("",*kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}


	deploymentsClient := clientset.AppsV1().Deployments("default")


	fmt.Println( "welcome to client-go-examples-test blah blah blah....")

	fmt.Println("What do you want to fuck off....Press that fucking key")


	fmt.Println("Press 1 to a Deployment")
	fmt.Println("Press 2 to Update the current Deployment")
	fmt.Println("Press 3 to Delete the current Deployment")
	var key int
	_, err = fmt.Scanf("%d",&key)
	if err !=nil {
		panic(err.Error())
	}

	///////// create a deployment ////////////

	deploymentSec:= func() int{
		if key!=1 {
			return 0
		}
		deployment := &appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name: "demo-deplpyment",
			},
			Spec: appsv1.DeploymentSpec{
				Replicas: int32Ptr(3),
				Selector: &metav1.LabelSelector{
					MatchLabels: map[string]string{
						"app": "main",
					},
				},
				Template: apiv1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"app": "main",
						},
					},
					Spec: apiv1.PodSpec{
						Containers: []apiv1.Container{
							{
								Name:  "hello",
								Image: "rakibulhossain/apiserver:0.0.3",
								Ports: []apiv1.ContainerPort{
									{
										Name:          "http",
										Protocol:      apiv1.ProtocolTCP,
										ContainerPort: 8081,
									},
								},
							},
						},
					},
				},
			},
		}

		fmt.Println("Creating Deployment...")
		result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
		return 1
	}()

	///////// Update a deployment ////////////

    updateDeploymentSec:=func () int{

		if key!=2 {
			return 0
		}

		retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {

			result, getErr := deploymentsClient.Get(context.TODO(), "demo-deplpyment", metav1.GetOptions{})
			if getErr != nil {
				panic(fmt.Errorf("Failed to get latest version of Deployment: %w", getErr))
			}

			result.Spec.Replicas = int32Ptr(2)
			_, updateErr := deploymentsClient.Update(context.TODO(), result, metav1.UpdateOptions{})
			return updateErr
		})

		if retryErr != nil {
			panic(fmt.Errorf("Update Failed: %w", retryErr))
		} else {
			fmt.Println("Updated Deployment....")
		}

		return 2

	}()

	///////// delete a deployment ////////////

	deleteDeploymentSec := func() int{
		if key != 3{
			return 0
		}
		deletePolicy := metav1.DeletePropagationForeground
		if err := deploymentsClient.Delete(context.TODO(), "demo-deplpyment", metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			panic(err)
		}
		fmt.Println("Deleted deployment....")
		return 3
	}()

	if key == 1{
		key=deploymentSec
	} else if key == 2{
		key=updateDeploymentSec
	} else if key == 3{
		key= deleteDeploymentSec
	}

	fmt.Println("Okay Fuck off now!!!")

}
