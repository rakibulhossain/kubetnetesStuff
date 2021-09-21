<h1>Kubectl & kind commands</h1>
<h2>Create cluster</h2> 

`kind create cluster`
<h2>Delete cluster</h2>

`kind delete cluster`
<h2>Kubectl commands</h2>


- <h3>Get command to find resources</h3>

    `kubectl get TYPE Name_Prefix`
    - frequently used options
        - wide output use `-o wide`
- <h3>Describe command to describe resources</h3>

    `kubectl describe TYPE Name_Prefix`
- <h3>Apply command to apply a confguaration</h3>

    `kubectl apply -f Manifest_File_Directory`
- <h3>Create Command to create a resource </h3>

    `kubectl create -f Filename`
- <h3>Delete command to delete a resource</h3>

    `kubectl delete ([-f FILENAME] | [-k DIRECTORY] | TYPE [(NAME | -l label | --all)]) [options]`
- <h3>Port forward to access a container resource</h3>

    `kubectl port-forward TYPE/NAME [options] [LOCAL_PORT:]REMOTE_PORT[...[LOCAL_PORT_N:]REMOTE_PORT_N]`
- <h3>Logs to see the current logs from the running instance</h3>

    `kubectl logs [-f] [-p] (POD | TYPE/NAME) [-c CONTAINER] [options]`
- <h3>Display cluster info</h3>

    `kubectl cluster-info` 
- <h3>Run a proxy to Kubernetes api server</h3>
    
    `kubectl proxy --port=[user-port]` 
- <h3>Edit a resource from the default editor</h3>

    `kubectl edit (RESOURCE/NAME | -f FILENAME)`
- <h3>Expose a resource as a new Kubernetes service</h3>

    `kubectl expose (-f FILENAME | TYPE NAME) [--port=port] [--protocol=TCP|UDP|SCTP] [--target-port=number-or-name] [--name=name] [--external-ip=external-ip-of-service] [--type=type]`
- <h3>Set a new size for a deployment, replica set, replication controller, or stateful set</h3>

    `kubectl scale [--resource-version=version] [--current-replicas=count] --replicas=COUNT (-f FILENAME | TYPE NAME)`
- <h3>Print the client and server version information for the current context</h3>

    `kubectl version`



