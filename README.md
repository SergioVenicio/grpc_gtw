# grpc_gtw
A simple grpc and http server with golang em scylladb


### setup docker
docker compose up -d

### setup scylladb
copy the `scylladb.cqlsh` content and post it on a database docker


### run app
go run ./cmd/api/main.go


### SWAGGER
http://localhost:8080/swagger


### GRPC SERVER
tcp://localhost:50051