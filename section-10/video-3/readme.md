# Packt Publishing - Hands on Microservices with Go
# Section 10 - Video 3 - The Saga Pattern

## Must Watch
[Distributed Sagas: A Protocol for Coordinating Microservices - Caitie McCaffrey - JOTB17](https://www.youtube.com/watch?v=0UTOLRTwOX0&list=LLXj2mZvgLzc01NOAljfZwlA&index=12=)

[GOTO 2015 • Applying the Saga Pattern • Caitie McCaffrey](https://www.youtube.com/watch?v=xDuwrtwYHu8)

### WARNING:

Don't try to implement the Saga Pattern unless you understand every word on the videos. And I mean every single word.

If you don't understand every concept there, just don't do it. Hire somebody that does.

### SECOND WARNING:

The example application we will be using is in many parts incomplete. It is an example to explain the basic concepts of the saga pattern not a reference implementation. Keep this in mind.

There are plenty of things that can go wrong in a distributed system that are not considered in the code.

### Core concepts

[Idempotence](https://en.wikipedia.org/wiki/Idempotence)

[Finite State Machines](https://en.wikipedia.org/wiki/Finite-state_machine)

## Kafka

[Apache Kafka](https://kafka.apache.org/)

[Packts Publishing - Apache Kafka](https://www.packtpub.com/big-data-and-business-intelligence/apache-kafka)

## Starting the example

You need to change the YOUR_HOME variable in scripts/setup.sh and scripts/setup-data.sh to your home.

Then run setup-data.sh first and setup.sh .

After that change to the saga-execution-coordinator directory and run main.go

**Warning**: Sometimes the sarama group consumer does not start consuming immediately. If this happens do an http request to the SEC to start a saga (that will not be processed), stop the SEC and start again and then do another request that will be processed. I really did not have the time to debug the issue, whether it has to do with the Kafka conf of the image we are using or if it's the driver we are using. The thing is once it starts it starts :-). Remember this is an example, not a reference implementation.
