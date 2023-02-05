# Dockerize back-end with databases

## Docker for golang + databases

Simple Go app running a server, initiating connection with mysql, rabbitmq and redis

Using DockerFile to compile Go app

Using docker-compose.yml for three databases and go app service 

## Prerequisites

Install docker, docker compose plugin

### Persistency

Create directories for volumes: 

```bash
mkdir ~/devops/data/mysql
mkdir ~/devops/data/redis
mkdir ~/devops/data/rabbitmq
```

## Usage

```bash
docker compose up
```