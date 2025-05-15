# sec-gen-control
The Secret Generator Controller (sec-gen-control) is a Kubernetes operator built with Kubebuilder that automatically generates and rotates Kubernetes Secrets based on a declarative MySecret custom resource.

It ensures password security by rotating credentials at a specified interval—fully automated, secure, and cluster-native.

### Prerequisites
```sh
✅ System Requirements used here
Go: v1.21 
Docker: v20+ 
Kustomize: v5.0+
Kubectl: 
Make: GNU Make (for build/deployment tasks)
✅ Kubernetes Requirements
Cluster access with kubectl configured (~/.kube/config)
Permissions to install CRDs, create namespaces, and manage RBAC resources
Kind – for spinning up a local K8s cluster
```

### Record the successful execution of the controller via Asciinema
```sh
/makro/sec-gen-control/mydemo.cast
```

## project structure
```sh
sec-gen-control/
├── config/             # CRDs, RBAC, Example
│   ├── crd/
│   ├── rbac/
│   ├── samples/
│   └── default/
├── controllers/        # Main logic
│   └── mysecret_controller.go
├── api/                # CRD definitions
│   └── v1/
├── examples/           # Example MySecret CRs
├── Dockerfile          # Docker build
├── Makefile            # Make build and push
├── go.mod              # Go dependencies
├── README.md    
```
### Custom Resource: MySecret
### A MySecret is a custom Kubernetes resource that defines:
username: static user name to inject into the Secret
passwordLength: how long the generated password should be
rotationPeriod: how often the password should be rotated

```sh
apiVersion: secrets.chinsecretgen.com/v1
kind: MySecret
metadata:
  name: my-generated-secret
  namespace: default
spec:
  secretType: basic-auth
  username: admin
  passwordLength: 40
  rotationPeriod: 1m
  ```

A Secret named my-generated-secret-generated will be created in the default namespace.



Getting Started
### Initialize the Project with Kubebuilder
```sh
kubebuilder init --domain chinsecretgen.com --repo github.com/chin008/makro/sec-gen-control
```
Initializes the controller project with domain chinsecretgen.com and Go module repo.

### Create API and Controller
```sh
kubebuilder create api --group secrets --version v1 --kind MySecret --controller --resource
```
Generates the MySecret CRD and a controller scaffold.

### Generate CRDs and Manifests
```sh
make manifests  make generates
```
### For local build
```sh
make install make build
```

Generates the CRDs in config/crd/ and webhook + RBAC configuration files.

### Build and Push Docker Image
```sh
make docker-build IMG=chinmay009/sec-gen-control:latest
```
Builds Docker image with the controller logic.

```sh
make docker-push IMG=chinmay009/sec-gen-control:latest
```
Pushes the image to Docker Hub or your registry.

### Deploy the Controller to Kubernetes
```sh
make deploy IMG=chinmay009/sec-gen-control:latest
```
Deploys the controller along with RBAC and CRDs to the cluster using Kustomize.

### Verifying Deployment
### Confirm the CRD Exists
```sh
kubectl get crd mysecrets.secrets.chinsecretgen.com
```
Ensures that your CRD is registered in the cluster.

### Confirm Controller is Running
```sh
kubectl get pods -n sec-gen-control-system
```
Lists pods running in the controller's namespace.
```sh
kubectl describe pod -n sec-gen-control-system | less
```
Check logs/events for success or issues.

### Sample CR Example
Apply a MySecret example CR:
```sh
kubectl apply -f config/samples/secrets_v1_mysecret.yaml
```

This instructs the controller to generate a Kubernetes Secret based on the MySecret spec.

### View Generated Secret
```sh
kubectl get secrets -n default
Lists secrets in the default namespace.

kubectl describe secret my-generated-secret-generated -n default
Details of the generated secret (metadata only).

kubectl get secret my-generated-secret-generated -n default -o yaml
Full YAML output of the secret including encoded credentials
```




**NOTE:** The makefile target mentioned above generates an 'install.yaml'
file in the dist directory. This file contains all the resources built
with Kustomize, which are necessary to install this project without its
dependencies.

2. Using the installer

Users can just run 'kubectl apply -f <URL for YAML BUNDLE>' to install
the project, i.e.:

```sh
kubectl apply -f https://raw.githubusercontent.com/<org>/sec-gen-control/<tag or branch>/dist/install.yaml
```

### By providing a Helm Chart  (Optional)  not

1. Build the chart using the optional helm plugin

```sh
kubebuilder edit --plugins=helm/v1-alpha
```

2. See that a chart was generated under 'dist/chart', and users
can obtain this solution from there.

**NOTE:** If you change the project, you need to update the Helm Chart
using the same command above to sync the latest changes. Furthermore,
if you create webhooks, you need to use the above command with
the '--force' flag and manually ensure that any custom configuration
previously added to 'dist/chart/values.yaml' or 'dist/chart/manager/manager.yaml'
is manually re-applied afterwards.

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

**NOTE:** Run `make help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

