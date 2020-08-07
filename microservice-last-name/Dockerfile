FROM python:3.7-slim

ADD requierments.txt /app/

RUN pip install -r /app/requierments.txt

ADD . /app

WORKDIR /app

ENTRYPOINT [ "python", "main.py" ]
