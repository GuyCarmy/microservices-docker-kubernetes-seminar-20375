# Microservice to handle my Last Name
This microservice is in charge of handling my last name.
## Preconditions
### for local
```
pip install flask==1.1.2
```
### for docker
```
docker build -t "guycarmy/microservice-last-name" .
```
## Run
### locally
```
python main.py
```
### dockerized
```
docker run -p 8091:8091 guycarmy/microservice-last-name:latest
```
## Test
```
curl http://localhost:8091/get_last_name
```
