#CLEAN UP
#USERS
sudo docker stop users-service
sudo docker rm users-service
sudo docker rmi "users-service"
sudo docker stop users-mariabd
sudo docker rm users-mariadb
#MANAGERS
sudo docker stop managers-service
sudo docker rm managers-service
sudo docker rmi "managers-service"
sudo docker stop managers-mariabd
sudo docker rm managers-mariadb
#SESSIONS
sudo docker stop sessions-service
sudo docker rm sessions-service
sudo docker rmi "sessions-service"
sudo docker stop sessions-redis
sudo docker rm sessions-redis
#VIDEOS
sudo docker stop videos-service
sudo docker rm videos-service
sudo docker rmi "videos-service"
sudo docker stop videos-mariadb
sudo docker rm videos-mariadb
#WTA
sudo docker stop wta-service
sudo docker rm wta-service
sudo docker rmi "wta-service"
sudo docker stop wta-psql
sudo docker rm wta-psql
#API GATEWAY 1
sudo docker stop api-gateway-1
sudo docker rm api-gateway-1
sudo docker rmi "api-gateway-1"
#API GATEWAY 2
sudo docker stop api-gateway-2
sudo docker rm api-gateway-2
sudo docker rmi "api-gateway-2"
