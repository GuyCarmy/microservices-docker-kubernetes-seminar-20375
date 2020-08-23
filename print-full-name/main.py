import requests
import os

first_name = requests.get(
    '{}/get_first_name'.format(
        os.getenv("MICROSERVICE_FIRST_NAME_URL")
    )
)

last_name = requests.get(
    '{}/get_last_name'.format(
        os.getenv("MICROSERVICE_LAST_NAME_URL")
    )
)

print('{} {}'.format(first_name.text, last_name.text))
