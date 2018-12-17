# Packt Publishing - Hands on Microservices with Go
# Section 5 - Video 1 - Load Testing with Apache JMeter.

## Start Application

We are going to be using the Users Application we created on Section 4. If you have not yet set up MariaDB with the user data follow the instructions on the readme for Section 4, Video 3.

```

sudo docker start users-mariadb

cd ~/go/src/github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-5/video-1/src/users-service

go main.go


```

## Note

You may look at the performance numbers and think that they are not very good (for example in terms of RPS). Please consider that the video was recorded on a 9 year old desktop computer. So you might ask why are you using such an old computer? Because when developing I know that I will eventually run the software I write on clusters of cheap commodity hardware, so it's good to actually test on an old computer when trying to squeeze as much performance as you can.

### Learn More

[Wikipedia - Percentile](https://en.wikipedia.org/wiki/Percentile)
[Apache JMeter](https://jmeter.apache.org/)


