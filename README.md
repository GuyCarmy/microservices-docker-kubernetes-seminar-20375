# microservices-docker-kubernetes-seminar-20375

This repo is the practical example for my Seminar (part of my university assignments, course number 20375).

The PDF document itself is attached.

The seminar covers the use of docker and kubernetes to deploy a microservice architected system and achieve good availability using autoscalers, healing and scheduling strategies.

The focus is on the infrastructure side and not the microservice architucture, using synchronous API as the main communication between services is not a best practice.

In this example we have 2 microservices, one written in Python and the other in Golang. This shows one of the advantages of microservices architecture, the use of the best langauage for the task. Its obvious that Go is better for handling first names and Python is superier for last names.

Each of the microservices have a dockerfile and some kuberentes configurations.

The services expose an http API which is consumed by a cronjob reminding the user of his full name.

## Run the entire system in a Docker network
```
docker network create "names-network"
docker run --network "names-network" --name "ms-first-name" -d "guycarmy/microservice-first-name:latest"
docker run --network "names-network" --name "ms-last-name" -d "guycarmy/microservice-last-name:latest"
docker run --network "names-network" \
           -e MICROSERVICE_FIRST_NAME_URL="http://ms-first-name:8090" \
           -e MICROSERVICE_LAST_NAME_URL="http://ms-last-name:8091" \
           "guycarmy/print-full-name:latest"
```
