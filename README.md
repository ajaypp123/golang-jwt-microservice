# Golang JWT Microservice

This application is a template for a microservice covering authentication APIs.

## Technology

- gin http package
- grpc communication
- mongo database

## TODO
- streaming kafka
- memory cache redis
- testing
- code coverage
- metrics
- logging improvement
- swagger
- monetoring

## Installation

1. Clone the repository:
```
git clone git@github.com:ajaypp123/golang-jwt-microservice.git
```

2. Install dependencies:
```
go mod tidy
go mod vendor
```

3. Build the application:
```
go build
```

## Docker

You can also run the microservice using Docker. Make sure you have Docker installed on your machine.

1. Build the Docker image:
```
docker compose build
```

2. Start service
```
docker compose up
```

## Endpoints

- `/signup`: POST request to sign up a new user.
- `/login`: POST request to log in and obtain a JWT token.
- `/user/{id}`: GET request to get user details by ID.
- `/refresh`: POST request to refresh the JWT token.

## Configuration

You can configure the application using environment variables:

- `HTTP_PORT`: Port on which the server listens (default is 8080).
- `MONGO_URI`: MongoDB connection URI.
- `GRPC_PORT`: Secret key for JWT token generation.

## Contributing

Contributions are welcome! Please feel free to submit pull requests.

## License

This project is licensed under the MIT License