# client-go-examples

I created this repo as I wanted to learn using ```client-go``` library for accessing Kubernetes API

# Instructions to use

Install
```bash
git clone https://github.com/reetasingh/client-go-examples.git
```

List Pods

```bash
reetasingh:practice-client-go reetasingh$ go run cmd/list_pods/main.go 
```

List Namespaces

```bash
reetasingh:practice-client-go reetasingh$ go run cmd/list_namespaces/main.go 

```

List pods with label run=abc
```bash
reetasingh:practice-client-go reetasingh$ go run cmd/list_pods_with_label_in_namespace/main.go 
There are 1 pods in the cluster
Pod name abc
Namespace default
ResourceVersion 735549
Number of containers = 1
Containers
    Container Name abc
    Container Image nginx
Labels map[run:abc]
Annotations map[]
==================
```

Watch event on pod (Added/Modified/Deleted)
```bash
reetasingh-ltm8:practice-client-go reetasingh$ go run cmd/watch_pods/main.go 
pod name abc - 
resource version 745618
Pod added
pod name abc - 
resource version 745662
Pod modified
pod name abc - 
resource version 745666
Pod modified
pod name abc - 
resource version 745685
Pod modified
pod name abc - 
resource version 745686
Pod deleted
^Z

```

for other actions check the ```cmd``` directory




