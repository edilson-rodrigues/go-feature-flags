minikube

1)Eval Shell to point to docker engine:

eval $(minikube docker-env)

build image:

docker build -t name:version .
docker tag

2) Expose service:

kubectl expose deployment app --type=NodePort --port=8080

3) Create tunnel to service:

minikube service app --url