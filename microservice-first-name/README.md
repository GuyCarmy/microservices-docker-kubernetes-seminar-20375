# Microservice to handle my First Name
This microservice is in charge of handling my first name.
## Preconditions
### for docker
```
docker build -t "guycarmy/microservice-first-name" .
```
## Run
### locally
```
go run main.go
```
### dockerized
```
docker run -p 8090:8090 guycarmy/microservice-first-name:latest
```
## Test - local or dockerized execution
```
curl http://localhost:8090/get_first_name
```
## deploy to kubernetes
```
kubectl apply -f kubernetes/deployment.yaml
kubectl apply -f kubernetes/service.yaml
kubectl apply -f kubernetes/hpa.yaml
```
