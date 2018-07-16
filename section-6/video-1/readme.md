# Packt Publishing - Hands on Microservices with Go
# Section 6 - Video 1 - Introduction to Docker.

## Inspect Images and Containers

```

//List all images
sudo docker images

//Show all running containers
sudo docker ps

Explain what is seen

//Show all containers
sudo docker ps -a

//Start, Stop, Restart
sudo docker start name-of-container

```

## Dockerizing an Application

```

cd ~/go/src/github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-6/video-1/example-1/

//Compile App
env GOOS=linux go build -o main main.go 

//Build Docker Image
sudo docker build -t example:1 . 

//Starting Container in Foreground Mode
sudo docker run --name=example-container-1 -p 8080:8080 example:1

//Starting Container in Background Mode
sudo docker run -d --name=example-container-1 -p 8080:8080 example:1

//Entering a Container Command Line (Bash) 
sudo docker exec -it example-container-1 /bin/bash

```  

## Removing Containers and Images

```

//Removing Container
sudo docker stop container-name
sudo docker rm container-name

//Remove an Image
sudo docker rmi image-name
//Another way to Remove an Image
sudo docker rmi image-id

```
## Inspecting a Container

```

sudo docker inspect container-name

```

## Learn More
[Container Runtime Constraints](https://docs.docker.com/engine/reference/run/#runtime-constraints-on-resources)
