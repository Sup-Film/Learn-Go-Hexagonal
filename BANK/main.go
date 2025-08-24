package main

import (
	"bank/repository"
	"bank/service"
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

	// สร้าง customerService โดยสร้าง NewCustomerService instance ใหม่ขึ้นมา และทำการโยน customerRepository เข้าไปเพื่อให้ service สามารถเข้าถึงข้อมูลได้
	customerService := service.NewCustomerService(customerRepository)

	customers, err := customerService.GetCustomers()
	if err != nil {
		panic(err)
	}

	fmt.Println(customers)

	customer, err := customerService.GetCustomer(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
}
