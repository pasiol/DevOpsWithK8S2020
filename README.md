## Part 0

Installing k3d Debian 10 server
	
[https://github.com/rancher/k3d]

	su -
	apt install docker.io curl wget vim sudo
	adduser devops
	usermod -aG sudo devops && init 6

	usermod -aG docker devops
	curl -s https://raw.githubusercontent.com/rancher/k3d/main/install.sh | bash
	k3d --version

	k3d version v3.1.5
	k3s version v1.18.9-k3s1 (default)


### kubectl

- [https://kubernetes.io/docs/tasks/tools/install-kubectl/]

	curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.18.9/bin/linux/amd64/kubectl
	chmod +x ./kubectl
	sudo mv ./kubectl /usr/local/bin/kubectl
	kubectl version --client

Client Version: version.Info{Major:"1", Minor:"18", GitVersion:"v1.18.9", GitCommit:"94f372e501c973a7fa9eb40ec9ebd2fe7ca69848", GitTreeState:"clean", BuildDate:"2020-09-16T13:56:40Z", GoVersion:"go1.13.15", Compiler:"gc", Platform:"linux/amd64"}

## kubectl

	kubectl cluster-info

	kubectl get deployments

	kubectl get pods

	kubectl logs POD_NAME
