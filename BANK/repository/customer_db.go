package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type customerRepositoryDB struct {
	db *pgx.Conn
}

func NewCustomerRepositoryDB(db *pgx.Conn) CustomerRepository {
	return &customerRepositoryDB{
		db: db,
	}
}

func (r *customerRepositoryDB) GetAll() ([]Customer, error) {
	customers := []Customer{}
	query := "select customer_id, name, date_of_birth, city, zip_code, status from customer"
	// ดึงผลลัพธ์จาก Query
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err // คืน error หาก query ล้มเหลว
	}
	defer rows.Close() // ปิด rows เมื่อใช้งานเสร็จ

	// อ่านค่าจาก rows และ map ไปยัง slice ของ Customer
	for rows.Next() {
		var customer Customer
		var dateOfBirth time.Time

		err := rows.Scan(&customer.CustomerID, &customer.Name, &dateOfBirth, &customer.City, &customer.ZipCode, &customer.Status)
		if err != nil {
			return nil, err // คืน error หาก scan ล้มเหลว
		}
		customer.DateOfBirth = dateOfBirth.Format("2006-01-02")
		customers = append(customers, customer)
	}

	// ตรวจสอบ error หลังจากวนลูป
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return customers, nil
}

func (r *customerRepositoryDB) GetById(id int) (*Customer, error) {
	customer := Customer{}
	query := "select customer_id, name, date_of_birth, city, zip_code, status from customer where customer_id = $1"

	var dateOfBirth time.Time

	err := r.db.QueryRow(context.Background(), query, id).Scan(&customer.CustomerID, &customer.Name, &dateOfBirth, &customer.City, &customer.ZipCode, &customer.Status)
	if err != nil {
		return nil, err // คืน error หาก query ล้มเหลวหรือไม่พบข้อมูล
	}
	customer.DateOfBirth = dateOfBirth.Format("2006-01-02")

	return &customer, nil
}
