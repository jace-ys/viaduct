# Docker Compose Example

A simple demo on using Viaduct with Docker Compose. Viaduct acts as an API gateway that proxies requests to a microservices backend.

## Installation

Refer to the main [README](https://github.com/jace-ys/viaduct#docker-compose-example).

## Services
1. API Gateway - [jaceys/viaduct](https://hub.docker.com/r/jaceys/viaduct)
2. Backend 1 (Users) - [williamyeh/docker-json-server](https://hub.docker.com/r/williamyeh/json-server)
3. Backend 2 (To Do's) - [williamyeh/docker-json-server](https://hub.docker.com/r/williamyeh/json-server)

## Endpoints

* http://localhost:5000/api/users/1

   Proxied to http://service.backend-1:3000/users/1. JSON response:

   ```
   {
     "id": 1,
     "name": "John Doe"
   }
   ```

* http://localhost:5000/api/todos/1

   Proxied to http://service.backend-2:3000/todos/1. JSON response:

   ```
   {
     "id": 1,
     "title": "Defeat Reverse Flash",
     "author": "Barry Allen"
   }
   ```
