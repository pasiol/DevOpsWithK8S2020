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


## Part 1

### Installing kubectl

- [https://kubernetes.io/docs/tasks/tools/install-kubectl/]

	curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.18.9/bin/linux/amd64/kubectl
	chmod +x ./kubectl
	sudo mv ./kubectl /usr/local/bin/kubectl
	kubectl version --client

Client Version: version.Info{Major:"1", Minor:"18", GitVersion:"v1.18.9", GitCommit:"94f372e501c973a7fa9eb40ec9ebd2fe7ca69848", GitTreeState:"clean", BuildDate:"2020-09-16T13:56:40Z", GoVersion:"go1.13.15", Compiler:"gc", Platform:"linux/amd64"}

### kubectl

[https://kubernetes.io/docs/reference/kubectl/docker-cli-to-kubectl/]

	kubectl cluster-info

	kubectl get deployments

	kubectl get pods

	kubectl logs POD_NAME

### Exercises

#### 1.01 Getting started

[https://github.com/pasiol/DevOpsWithK8S2020/tree/main/go-main-app]

	docker build ./ -t pasiol/go-main-app:latest
	docker push pasiol/go-main-app:latest
	docker image ls | grep go-main-app
pasiol/go-main-app                                      latest              a1bd3753f6dc        27 minutes ago      1.79MB
	
	docker run pasiol/go-main-app:latest
	2020-11-01T11:14:36.146902677Z 4469fb90-6751-4387-b332-58f6c2de4808
	2020-11-01T11:14:41.147231488Z e7395ff0-3a83-4d34-adba-124f8029b4ca

	kubectl create deployment go-main-app --image=pasiol/go-main-app
	kubectl get deployments
	NAME          READY   UP-TO-DATE   AVAILABLE   AGE
	go-main-app   1/1     1            1           96s

	kubectl get pods
	NAME                           READY   STATUS    RESTARTS   AGE
	go-main-app-676c66fbc6-dxd4s   1/1     Running   0          2m59s

	kubectl logs go-main-app-676c66fbc6-dxd4s
	2020-11-01T11:18:27.635233449Z bccebb61-46cb-4840-a19d-0220b6d37ae5
	2020-11-01T11:18:32.635950697Z dbe17776-eb8b-4dea-888c-2a5b2225b218
	2020-11-01T11:18:37.636461087Z bab6611f-1384-4b6f-995c-6f98627b7689
	2020-11-01T11:18:42.637032246Z ca7022f1-a78f-4bde-b77d-c4bb7acabf73

#### 1.02 Projevt v0.1

Project repo: [https://github.com/pasiol/django-to-do-app]

	kubectl create deployment django-to-do-app --image=pasiol/django-to-do-app:1.02

	devops@devops:~$ kubectl get pods
	NAME                                READY   STATUS              RESTARTS   AGE
	go-main-app-676c66fbc6-dxd4s        1/1     Running             0          98m
	django-to-do-app-5bbb687848-45kh6   0/1     ContainerCreating   0          52s
	devops@devops:~$ kubectl logs django-to-do-app-5bbb687848-45kh6
	Performing system checks...

	System check identified no issues (0 silenced).

	You have 18 unapplied migration(s). Your project may not work properly until you apply the migrations for app(s): admin, auth, contenttypes, sessions.
	Run 'python manage.py migrate' to apply them.
	November 01, 2020 - 12:56:40
	Django version 3.1.2, using settings 'devopsToDoApp.settings'
	Starting development server at http://0.0.0.0:8000/
	Quit the server with CONTROL-C.


