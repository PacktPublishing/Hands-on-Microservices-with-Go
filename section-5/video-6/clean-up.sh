sudo docker stop users-cache-redis
sudo docker rm users-cache-redis
sudo docker run --name users-cache-redis -p 6379:6379 -d redis
sudo docker stop users-mariadb
sudo docker rm users-mariadb
sudo docker run --name users-mariadb -v usersmariadb:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=root-password -p 3306:3306 -d mariadb
sudo docker stop prometheus
sudo docker rm prometheus
sudo docker run --net=host --volume=/home/emiliano/go/src/github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-5/video-6/conf/prometheus.yml:/etc/prometheus/prometheus.yml --name=prometheus -d prom/prometheus

