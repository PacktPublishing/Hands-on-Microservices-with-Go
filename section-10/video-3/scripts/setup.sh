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
go build -a -o ../main main.go
#VIDEOS
cd $PATH_TO_VIDEO_FOLDER"videos-service/src/"
go build -a -o ../main main.go
#SAGA EXECUTION COORDINATOR
cd $PATH_TO_VIDEO_FOLDER"saga-execution-coordinator/"
go build -a -o ../main main.go

cd $PATH_TO_VIDEO_FOLDER"scripts"
#RUN DOCKER COMPOSE
sudo docker-compose up -d --build