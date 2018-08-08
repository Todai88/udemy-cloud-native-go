# KUBERNETES

### What's this about?
This readme is part of the Packt course on GoLang Microservices and will outline my experiences with Kubernetes following a quick write up on how Kubernetes works, what benefits it brings to the table and how it integrates into a microservice environment.

### What is Kubernetes?

### How does Kubernetes work?

### What do I use Kubernetes for?

### What benefits compared to other xyz software to handle xyz?
How is Kubernetes better than other software that solves the same thing?

## Issues / Problems:

### Can't find image
Remember that Kubernetes has it's own version of Docker running on it, and if you use need to use a local image in your kubernetes cluster you need to build it in a Kubernetes shell too. This can be done like this:

**Windows:**
`minikube docker-env` and
`minikube docker-env | Invoke-Expression` to get access to the shell
`docker build .` or `docker-compose up` or whatever rocks your boat.

**Linux:**

### Environment:
There were issues where `minikube-net` wasn't working as intended between my two environments (Ubuntu and Windows 10). 

* **Windows 10:** Configure a virtual router and use Hyper-V as the intended driver.
* **Ubuntu:** Use `virsh` to start the virtual network:
```
1. Type virsh
2. Type net-list --all - should see that minikube-net is inactive
3. Type net-start minikube-net - should get an error message about "Network is already in use by interface virbr1" or similar
4. Quit virsh
5. Type sudo ifconfig virbr1 down
6. Type sudo brctl delbr virbr1
7. Type virsh
8. Type net-start minikube-net - should now start-up
9. Quit virsh
10. Type minikube start
```

