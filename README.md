# Running

The API runs on port 8081.

```shell
docker-compose up
go mod tidy
go run cmd/rv/main.go 
```

# Tests

Some of the tests require the database to be runnint

```shell
docker-compose up
go mod tidy
go test ./...  
```

