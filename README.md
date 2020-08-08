# microservices-docker-kubernetes-seminar-20375

This repo is the practical example for my Seminar (part of my university assignments, course number 20375).

The seminar covers the use of docker and kubernetes to deploy a microservice architected system and achieve good availability using autoscalers, healing and scheduling strategies.

In this example we have 2 microservices, one written in Python and the other in Golang. This shows one of the advantages of microservices architecture, the use of the best langauage for the task. Its obvious that Go is better for handling first names and Python is superier for last names.

Each of the microservices have a dockerfile and some kuberentes configurations.

The services expose an http API.
