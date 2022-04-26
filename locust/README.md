# locust

## Installation

    sudo -H pip3 install locust

## Run test

    locust --locustfile locustfile.py --headless --host http://localhost:8080 --users 60 --spawn-rate 1 --run-time 20s
