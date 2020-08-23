FROM python:3.7-slim

RUN pip install requests==2.24.0

ADD . /job

WORKDIR /job

ENTRYPOINT [ "python", "main.py" ]
