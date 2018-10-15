#CLEAN UP
#USERS
sudo docker stop users-service
sudo docker rm users-service
sudo docker rmi "users-service"
sudo docker stop users-mariabd
sudo docker rm users-mariadb
#MANAGERS
sudo docker stop agents-service
sudo docker rm agents-service
sudo docker rmi "agents-service"
sudo docker stop agents-mariabd
sudo docker rm agents-mariadb
#SESSIONS
sudo docker stop sessions-service
sudo docker rm sessions-service
sudo docker rmi "sessions-service"
sudo docker stop sessions-redis
sudo docker rm sessions-redis
#sec
sudo docker stop saga-execution-coordinator
sudo docker rm saga-execution-coordinator
sudo docker rmi "saga-execution-coordinator"
