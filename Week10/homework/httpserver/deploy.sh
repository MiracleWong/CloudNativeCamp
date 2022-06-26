#!/bin/bash
echo "begin dir: "$(pwd)
# build image
docker build -t httpserver:v10 -f dockerfile-mod .
docker tag httpserver:v10 miraclewong/httpserver:v10
docker push miraclewong/httpserver:v9

if [ 0 != $? ]
then
  echo "docker image build faild"
  exit 1
fi

# info
kubectl get pod
kubectl get svc
kubectl get ingress

# deploy
kubectl apply -f httpserver-deployment.yaml
kubectl apply -f httpserver-svc.yaml
kubectl apply -f httpserver-ingress.yaml

# info
kubectl get pod
kubectl get svc
kubectl get ingress

# test
curl 127.0.0.1:30001/healthz
curl curl httpserver.com/healthz