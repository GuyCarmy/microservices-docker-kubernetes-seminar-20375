FROM python:3.7-slim

RUN pip install flask==1.1.2

ADD . /app

WORKDIR /app

ENTRYPOINT [ "python", "main.py" ]
