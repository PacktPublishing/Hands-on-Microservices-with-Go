#/!bin/bash

#SET HOME FOLDER
#CHANGE THIS TO YOUR HOME FOLDER
YOUR_HOME="/home/emiliano"

#YOUR TARGET OS
GOOS="linux"

#SET PATH TO THIS FOLDER
PATH_TO_VIDEO_FOLDER=$YOUR_HOME"/go/src/github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-1/"

#STOP DOCKER COMPOSE IF IT's RUNNING
sudo docker-compose down

#CLEAN UP
./clean-up.sh

#GENERATE RUNNABLES
#USERS
cd $PATH_TO_VIDEO_FOLDER"src/api"
go build -a -o ../scripts/main main.go

cd $PATH_TO_VIDEO_FOLDER"src/scripts"
#RUN DOCKER COMPOSE
sudo docker-compose up -d --build
