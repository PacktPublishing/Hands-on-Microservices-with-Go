# Packt Publishing - Hands on Microservices with Go
# Section 1 - Video 6 - Using a Database and Creating a CRUD app.

## Starting MongoDB Instance with Docker

`sudo docker run --name s1v6mongo -d -p 27017:27017 mongo`

## Entering the Instance and looking around

`sudo docker exec -it s1v6mongo /bin/bash`

### Entering Mongo Cli

`mongo`

### Looking around:

Show all existing Databases:

`show dbs`

Selecting a Database:

`use packt`

Show all collections in a Database:

`show collections`

Playing around with a Collection:

```
coll = db.timeZones

coll.help()
coll.count()
coll.find()
```

## Postman Collection

[Collection](https://github.com/PacktPublishing/Hands-on-Microservices-with-Go/blob/master/section-1/video-6/S1V6.postman_collection.json)

## Learn More

[Wikipedia REST: Relationship between URL and HTTP methods](https://en.wikipedia.org/wiki/Representational_state_transfer#Relationship_between_URL_and_HTTP_methods)

[MongoDB](https://www.mongodb.com/)

[Wikipedia's MongoDB Page](https://en.wikipedia.org/wiki/MongoDB)

[Wikipedia's Page on Document Oriented DBs](https://en.wikipedia.org/wiki/Document-oriented_database)

[Mongo Shell Quick Reference](https://docs.mongodb.com/manual/reference/mongo-shell/)

[Community Maintained MongoDB Go Driver](https://github.com/globalsign/mgo)

[MongoDB Official Go Driver (Not ready for Production yet)](https://github.com/mongodb/mongo-go-driver)

[How to Secure a MongoDB Production Database](https://www.cyberciti.biz/faq/how-to-secure-mongodb-nosql-production-database/)

[MongoDB Security](https://docs.mongodb.com/manual/security/)

### Learn Even More

[Packt Publishing - MongoDB Starting Guide (Free!)](https://www.packtpub.com/packt/free-ebook/mongoDB-starter-guide)

[Packt Publishing - MongoDB Cookbook](https://www.packtpub.com/big-data-and-business-intelligence/mongodb-cookbook)

[Packt Publishing - Developing with MongoDB Video](https://www.packtpub.com/big-data-and-business-intelligence/developing-mongodb-video)
