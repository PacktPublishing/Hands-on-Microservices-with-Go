#CLEAN UP
#AGENTS
sudo docker stop agents-service
sudo docker rm agents-service
sudo docker rmi "agents-service"
sudo docker stop agents-mariabd
sudo docker rm agents-mariadb
