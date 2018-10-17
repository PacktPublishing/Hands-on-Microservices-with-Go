# Packt Publishing - Hands on Microservices with Go
# Section 6 - Video 3 - Docker Compose.

## Installing Docker Compose

```

sudo curl -L https://github.com/docker/compose/releases/download/1.21.2/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose

sudo chmod +x /usr/local/bin/docker-compose

//Test Installation
docker-compose --version

```

## Clean up old Docker Containers

```

sudo docker stop users-api
sudo docker rm users-api
sudo docker rmi users-api:1

sudo docker stop users-mariadb
sudo docker rm users-mariadb

sudo docker stop users-cache-redis
sudo docker rm users-cache-redis

sudo docker stop prometheus
sudo docker rm prometheus

```

## Build our App

```

cd ~/go/src/github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-6/video-3/src/users-service
env GOOS=linux go build -a -o main main.go

sudo docker-compose up -d

//Show current containers on compose
sudo docker-compose ps
//Stop multi container App
sudo docker-compose down

```

## Learn More

[Install Docker Compose](https://docs.docker.com/compose/install/#install-compose)

[Kubernetes](https://kubernetes.io/)
