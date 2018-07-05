# Packt Publishing - Hands on Microservices with Go
# Section 5 - Video 5 - Load Balancing with Nginx.

## Starting Nginx with our configuration

```

//Replace $USERHOME to your Home folder

docker run --name load-nginx -v $USERHOME/go/src/github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-5/video-5/conf:/etc/nginx:ro --net host -d nginx

```

## Learn More

(Wikipedia Load Balancing)[https://en.wikipedia.org/wiki/Load_balancing_(computing)]
(Wikipedia Nginx)[https://en.wikipedia.org/wiki/Nginx]
(Nginx)[https://nginx.org/en/]
