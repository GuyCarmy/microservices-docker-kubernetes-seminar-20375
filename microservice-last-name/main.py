from flask import Flask

api = Flask(__name__)

LAST_NAME = "Carmy"


@api.route('/_ready')
def readiness():
    return "im ready"


@api.route('/_alive')
def liveness():
    return "im alive"


@api.route('/get_last_name')
def last_name():
    return LAST_NAME


if __name__ == '__main__':
    api.run(port=8091, host='0.0.0.0')
