# Load test target web servers

Web servers that serve a short text on `/health`.

## simple

Nothing fancy. It responds to requests as quickly as possible.

    go run webserver/main-simple.go

## advanced

A web server that only handles X requests simultaneously (defined by SEMAPHORE_SLOTS, default 200).

Each request takes a random time to complete with a normal distribution.

When server has no available slots for handling an incoming request it has a 10% chance of immediate failure.

    go run webserver/main-advanced.go

    And optionally set parallelism:

    SEMAPHORE_SLOTS=10 go run webserver/main-advanced.go

## Installing golang

    wget https://go.dev/dl/go1.18.1.linux-amd64.tar.gz
    sudo rm -rf /usr/local/go
    sudo tar -C /usr/local -xzf go1.18.1.linux-amd64.tar.gz
    export PATH=$PATH:/usr/local/go/bin
    rm https://go.dev/dl/go1.18.1.linux-amd64.tar.gz

## Developing

    docker build -t stianovrevage/go-test-webservers:latest .
    docker push stianovrevage/go-test-webservers:latest

Run the simple server:

    docker run -p 8080:8080 stianovrevage/go-test-webservers:latest

Run the advanced server:

    docker run -p 8080:8080 -e SEMAPHORE_SLOTS=1000 stianovrevage/go-test-webservers:latest /app/advanced-server

## Backlog

- Metrics
  - Expose prometheus-style metrics
  - Measure standard HTTP metrics
  - Measure standard Go metrics
  - Custom metrics for artificial queue contention
- Multi-arch container image
  - buildx with amd64 + arm
