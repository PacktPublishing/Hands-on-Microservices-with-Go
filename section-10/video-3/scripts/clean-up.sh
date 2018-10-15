#CLEAN UP
#USERS
sudo docker stop users-service
sudo docker rm users-service
sudo docker rmi "users-service"
sudo docker stop users-mariabd
sudo docker rm users-mariadb
#AGENTS
sudo docker stop agents-service
sudo docker rm agents-service
sudo docker rmi "agents-service"
sudo docker stop agents-mariabd
sudo docker rm agents-mariadb
#VIDEOS
sudo docker stop videos-service
sudo docker rm videos-service
sudo docker rmi "videos-service"
sudo docker stop videos-mariabd
sudo docker rm videos-mariadb
#KAFKA
sudo docker stop kafka
sudo docker rm kafka
