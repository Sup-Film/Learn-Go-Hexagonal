package main

import (
	"bank/repository"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:1234@localhost:5432/bank?sslmode=disable")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	customerRepository := repository.NewCustomerRepositoryDB(conn)

	customers, err := customerRepository.GetAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(customers)

	customer, err := customerRepository.GetById(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
}
