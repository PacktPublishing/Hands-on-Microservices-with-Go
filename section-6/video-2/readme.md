# Packt Publishing - Hands on Microservices with Go
# Section 6 - Video 2 - Docker Networking and Data Management.

## Working with User Created Networks

```


cd ~/go/src/github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-6/video-2/src/users-service/

//Remove old versions of containers
sudo docker stop users-mariadb
sudo docker rm users-mariadb

sudo docker stop users-cache-redis
sudo docker rm users-cache-redis

//Show networks
sudo docker network ls

//create your own bridge network
sudo docker network create --driver bridge users-network

//Then you should be able to connect to networks using the container name - Automatic Service Discovery
sudo docker run --name users-cache-redis --network users-network -d redis
sudo docker run --name users-mariadb -v usersmariadb:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=root-password --network users-network -d mariadb


//show containers on network bridge
sudo docker network inspect users-network

//Build APP Image and run it on container
env GOOS=linux go build -a -o main main.go 
sudo docker build -t user-api:1 . 
sudo docker run --name user-api --network users-network -p 8000:8000 -d users-api:1

//Inspect individual container
sudo docker inspect users-cache-redis

```

## Working with Data Volumes

```

//List Volumes
sudo docker volume ls

//Create a Data Volume
sudo docker volume create name-of-volume

//Removing a Volume.
//__Warning everything will be deleted.__
sudo docker volume rm name-of-volume	

```

## Learn More
[Backup & Restore Docker Named Volumes](https://loomchild.net/2017/03/26/backup-restore-docker-named-volumes/)
