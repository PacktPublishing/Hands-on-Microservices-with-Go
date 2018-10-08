#CLEAN UP
#USERS
sudo docker stop users-api
sudo docker rm users-api
sudo docker rmi "users-api"
sudo docker stop users-cache-redis
sudo docker rm users-cache-redis
sudo docker rmi "users-cache-redis"
