from flask import Flask

api = Flask(__name__)

LAST_NAME = "Carmy"


@api.route('/_ready')
def readiness():
    return "OK"


@api.route('/_alive', methods=['GET'])
def liveness():
    return "OK"


@api.route('/get_last_name', methods=['GET'])
def last_name():
    return LAST_NAME


if __name__ == '__main__':
    api.run(port=8091)