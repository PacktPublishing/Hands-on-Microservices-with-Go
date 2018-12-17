# Packt Publishing - Hands on Microservices with Go
# Section 5 - Video 2 - Profiling with PProf and Torch.

## Installing Graphviz

```

sudo apt-get install graphviz


```

## Installing Go Torch

```

go get github.com/uber/go-torch

cd ~/Soft

git clone git@github.com:brendangregg/FlameGraph.git

```

Add the following to your ~/.bash_profile

```

export PATH-$PATH:/Path/to/FlameGraph

```

## Getting an accessing a Trace

```
wget http://localhost:8000/debug/pprof/trace?seconds=60

mv trace\?seconds\=60 trace-60-seconds

go tool trace trace-60-seconds

```

## Accessing a PPROF profile 

```

//CPU Profile on Command Line
go tool pprof -seconds=180 http://localhost:8000/debug/pprof/profile

//CPU Profile Web Interface
go tool pprof -seconds=180 -http localhost:15000 http://localhost:8000/debug/pprof/profile

//CPU Profile on Command Line from existing Profile
go tool pprof profile-file-name

//Heap Profile on Command Line
go tool pprof http://localhost:8000/debug/pprof/heap


```

## PPROF command line 

```

//Show all traces on this profile
(pprof) traces
//Show top 
(pprof) top
//Show top 20 Cumulative
(pprof) top20 -cum
//Generate an SVG Graph
(pprof) svg > cpu-graph.svg
//See all other commands
(pprof) help


```

## Generating a Flame Graph

```
go-torch --file "torch.svg" --url http://localhost:8000

```

### Learn More

[Download Graphviz](https://graphviz.gitlab.io/download/)

[Profiling with PProf](https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/)

[Go Blog - Profiling Go Programs](https://blog.golang.org/profiling-go-programs)

[Go Documentation - Package PProf](https://blog.golang.org/profiling-go-programs)

[Flamegraphs](http://www.brendangregg.com/flamegraphs.html)

## Apologies!

Emiliano says: "Sorry about the coughing on the video, I had some health issues during the recording of this video."

## Note

You may look at the performance numbers and think that they are not very good (for example in terms of RPS). Please consider that the video was recorded on a 9 year old desktop computer. So you might ask why are you using such an old computer? Because when developing I know that I will eventually run the software I write on clusters of cheap commodity hardware, so it's good to actually test on an old computer when trying to squeeze as much performance as you can.



