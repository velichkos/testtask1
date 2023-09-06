package main

import (
	"github.com/outdoorsy/interview-challenge-backend/pkg/rv"
)

func main() {
	connectionString := "postgres://root:root@localhost:5434/testingwithrentals?sslmode=disable"

	repository := rv.NewRepository(connectionString)
	server := rv.NewServer(repository)

	panic(server.Start(8081))
}
