package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/retry"
	"path/filepath"

)

func main(){
	//// auth /////
	var kubeconfig *string
	if home:= homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig",filepath.Join(home,".kube","config"),"(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig","","absolute path to the kubeconfig file")
	}
	flag.Parsed()
	namespace := "default"

	config,err := clientcmd.BuildConfigFromFlags("",*kubeconfig)
	if err != nil {
		panic(err)
	}
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	deploymentRes :=  schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}

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

	////// creating a deployment ///////

	deploymentSec := func() int{

		if key != 1 {
			return 0
		}

		deployment := &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "apps/v1",
				"kind": "Deployment",
				"metadata" : map[string]interface{}{
					"name" : "demo-deployment",
				},
				"spec": map[string]interface{}{
					"replicas" : 3,
					"selector" : map[string]interface{}{
						"matchLabels": map[string]interface{}{
							"app" : "main",
						},
					},
					"template" : map[string]interface{}{
						"metadata" : map[string]interface{}{
							"labels" : map[string]interface{}{
								"app" : "main",
							},
						},

						"spec" : map[string]interface{}{
							"containers" : []map[string]interface{}{
								{
									"name": "hello",
									"image": "rakibulhossain/apiserver:0.0.3",
									"ports": []map[string]interface{}{
										{
											"name":          "http",
											"protocol":      "TCP",
											"containerPort": 8081,
										},
									},
								},
							},
						},
					},

				},
			},
		}
		fmt.Println("Creatinig a deployment")
		result, err := client.Resource(deploymentRes).Namespace(namespace).Create(context.TODO(),deployment,metav1.CreateOptions{});
		if err != nil {
			panic(err)
		}

		fmt.Printf("Created Deployment %q.\n",result.GetName())
		return 1
	}()


	//////// update the current deployment /////////


	updateDeploymentSec := func() int {

		if key != 2 {
			return 0
		}

		retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
			result ,getErr := client.Resource(deploymentRes).Namespace(namespace).Get(context.TODO(),"demo-deployment",metav1.GetOptions{})
			if getErr != nil {
				panic(fmt.Errorf("failed to get latest version of Deployment: %v",getErr))
			}

			if err := unstructured.SetNestedField(result.Object,int64(2),"spec","replicas"); err != nil {
				panic(fmt.Errorf("failed to replica value: %v",err))
			}

			_, updateErr := client.Resource(deploymentRes).Namespace(namespace).Update(context.TODO(),result,metav1.UpdateOptions{})
			return updateErr
		})

		if retryErr != nil {
			panic(fmt.Errorf("update failed: %v", retryErr))
		}
		fmt.Println("Updated deployment.....")
		return 2
	}()

	///// delete the deployment /////////

	deleteDeploymentSec := func() int{

		if key != 3 {
			return 0
		}

		fmt.Println("Deleting deployment")
		deletePolicy := metav1.DeletePropagationForeground


		if err := client.Resource(deploymentRes).Namespace(namespace).Delete(context.TODO(),"demo-deployment",metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			panic(err)
		}

		fmt.Println("Deleted Deployment.....")
		return 3
	}()



	if key == 1{
		key=deploymentSec
	} else if key == 2{
		key=updateDeploymentSec
	} else if key == 3{
		key=deleteDeploymentSec
	}

	fmt.Println("Okay Fuck off now!!!")



}
