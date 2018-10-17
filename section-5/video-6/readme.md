# Packt Publishing - Hands on Microservices with Go
# Section 5 - Video 6 - Instrumentation: Collecting Metrics and monitoring with Prometheus.

## Starting Prometheus with our configuration

```

sudo docker run --net=host --volume=$USERHOME/go/src/github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-5/video-6/conf/prometheus.yml:/etc/prometheus/prometheus.yml --name=prometheus -d prom/prometheus


```


## Learn More

[Prometheus](https://prometheus.io/)

[Prometheus - Histograms and Summaries](https://prometheus.io/docs/practices/histograms/)

[Prometheus - Query Language](https://prometheus.io/docs/prometheus/latest/querying/basics/)

[Exploring Prometheus Go client metrics](https://povilasv.me/prometheus-go-metrics/)

[Grafana](https://grafana.com/)

