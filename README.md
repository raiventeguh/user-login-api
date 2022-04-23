# user-login-api
Golang example for User Login API

## Pre-Requisite
* Docker 
* Docker-compose

## Implementation Detail
* Hexagonal (not fully implemented)
* Swagger - Echo Swagger 
* Docker & Docker Compose

## Installation
```sh
cd users
docker-compose up --build -d
```

You can access Golang swagger from following url:
http://localhost:8081/swagger/index.html

## API Credentials
user admin
```text
email: admin@gmail.com
password: SECRET
```

user non admin
````text
email: user@gmail.com
password: SECRET
````

## API Usage
API using Basic Token
```text
Authorization: Bearer {Token}
```