# Microservice to handle my Last Name
This microservice is in charge of handling my last name.
## Preconditions
### for local
```
pip install requests==2.24.0
export MICROSERVICE_LAST_NAME_URL=http://localhost:8091
export MICROSERVICE_FIRST_NAME_URL=http://localhost:8090
```
### for docker
```
docker build -t "guycarmy/print-full-name" .
```
## Run
### locally
```
python main.py
```
### dockerized
```
docker run --network host \
           -e MICROSERVICE_LAST_NAME_URL="http://localhost:8091" \
           -e MICROSERVICE_FIRST_NAME_URL="http://localhost:8090" \
           "guycarmy/print-full-name:latest"
 ```
