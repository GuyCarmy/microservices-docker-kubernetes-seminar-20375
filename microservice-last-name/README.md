# Microservice to handle my Last Name
This microservice is in charge of handling my last name.
## Run
### locally
```
pip install flask==1.1.2
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
