# Packt Publishing - Hands on Microservices with Go
# Section 4 - Video 3 - Access Tokens and JWT

### Setup Environment

#### Create DB Data Volume
```

sudo docker volume create usersmariadb

```

#### Fill up Data Volume

**$USERHOME** should be your home directory, where go is installed.

```

sudo docker run -it -v usersmariadb:/volume -v $USERHOME/go/src/github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-4/video-3/data/:/backup ubuntu \
    sh -c "rm -rf /volume/* /volume/..?* /volume/.[!.]* ; tar -C /volume/ -xjf /backup/usersmariadb.tar.bz2"


```

#### Start MariaDB
```

sudo docker run --name users-mariadb -v usersmariadb:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=root-password -p 3306:3306 -d mariadb

```

#### Start Redis

```

sudo docker run --name sessions-redis -p 6379:6379 -d redis

```

### Passwords on database

Passwords are encrypted but if you need to test the example. The passwords are based on the the following logic, if userID%10 is 0 then the password for the user will be "academia-racing-club". The same logic applies for the different values of userID%10, use the following guide:

```

userID%10 == 0 => password = "academia-racing-club"
userID%10 == 1 => password = "san-lorenzo-de-almagro"
userID%10 == 2 => password = "boca-juniors"
userID%10 == 3 => password = "river-plate"
userID%10 == 4 => password = "club-atletico-independiente"
userID%10 == 5 => password = "club-social-y-deportivo-yupanki"
userID%10 == 6 => password = "sacachispas-futbol-club"
userID%10 == 7 => password = "defensa-y-justicia"
userID%10 == 8 => password = "chacarita-juniors"
userID%10 == 9 => password = "arsenal-de-sarandi"


```

### Learn more

[JWT](https://jwt.io/)

[JSON Web Token Best Current Practices](http://self-issued.info/docs/draft-sheffer-oauth-jwt-bcp-00.html)

[Wikipedia JSON Web Tokens](https://en.wikipedia.org/wiki/JSON_Web_Token)

[JOSE Standard](http://jose.readthedocs.io/en/latest/)

[JWT-Go Package](https://github.com/dgrijalva/jwt-go)

[RFC 7519: JSON Web Token (JWT)](https://tools.ietf.org/html/rfc7519)

[Wikipedia HMAC](https://en.wikipedia.org/wiki/HMAC)

