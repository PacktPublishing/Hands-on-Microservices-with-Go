#/!bin/bash

#SET HOME FOLDER
#CHANGE THIS TO YOUR HOME FOLDER
YOUR_HOME="/home/emiliano"

#YOUR TARGET OS
GOOS="linux"

#SET PATH TO THIS FOLDER
PATH_TO_VIDEO_FOLDER=$YOUR_HOME"/go/src/github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-10/video-3/"

#STOP DOCKER COMPOSE IF IT's RUNNING
sudo docker-compose down

#CLEAN UP
./clean-up.sh

#GENERATE RUNNABLES
#USERS
cd $PATH_TO_VIDEO_FOLDER"users-service/"
go build -a -o main main.go
#AGENTS
cd $PATH_TO_VIDEO_FOLDER"agents-service/"
go build -a -o main main.go
#VIDEOS
cd $PATH_TO_VIDEO_FOLDER"videos-service/src/"
go build -a -o ../main main.go

cd $PATH_TO_VIDEO_FOLDER"scripts"
#RUN DOCKER COMPOSE
sudo docker-compose build
sudo docker-compose up -d --build --force-recreate

#RUN KAFKA
sudo docker run -d --name kafka -v kafka-logs:/var/lib/kafka -p 2181:2181 -p 9092:9092 --env ADVERTISED_HOST=127.0.0.1 --env ADVERTISED_PORT=9092 spotify/kafka 
#RUN SEC
cd $PATH_TO_VIDEO_FOLDER"saga-execution-coordinator/"
