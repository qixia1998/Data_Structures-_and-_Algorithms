package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Customer class
type Customer struct {
	CustomerId int
	CustomerName string
	SSN string
}

// GetConnection method which returns sql.DB
func GetConnection() (database *sql.DB) {
	databaseDriver := "mysql"
	databaseUser := "root"
	databasePass := "123456"
	databaseName := "crm"
	database, error := sql.Open(databaseDriver,
		databaseUser+":"+databasePass+"@/"+databaseName)
	if error != nil {
		panic(error.Error())
	}
	return database
}
// GetCustomers method returns Customer Array
func GetCustomers() []Customer {
	var database *sql.DB
	database = GetConnection()

	var error error
	var rows *sql.Rows
	rows, error = database.Query("SELECT * FROM Customer ORDER BY Customerid DESC")
	if error != nil {
		panic(error.Error())
	}
	var customer Customer
	customer = Customer{}

	var customers []Customer
	customers = []Customer{}
	for rows.Next() {
		var customerId int
		var customerName string
		var ssn string
		error = rows.Scan(&customerId, &customerName, &ssn)
		if error != nil {
			panic(error.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = ssn
		customers = append(customers, customer)
	}
	defer database.Close()
	return customers
}
// InsertCustomer method with parameter customer
func InsertCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()

	var error error
	var insert *sql.Stmt
	insert, error = database.Prepare("INSERT INTO Customer(CustomerName, SSN) VALUES (?, ?)")
	if error != nil {
		panic(error.Error())
	}
	insert.Exec(customer.CustomerName, customer.SSN)

	defer database.Close()
}
// Update Customer method with parameter customer
func UpdateCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()
	var error error
	var update *sql.Stmt
	update, error = database.Prepare("UPDATE Customer SET CustomerName=?, SSN=? WHERE CustomerId=?")
	if error != nil {
		panic(error.Error())
	}
	update.Exec(customer.CustomerName, customer.SSN, customer.CustomerId)

	defer database.Close()
}
// Delete Customer method parameter customer
func deleteCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()
	var error error
	var delete *sql.Stmt
	delete, error = database.Prepare("DELETE FROM Customer WHERE CustomerId=?")
	if error != nil {
		panic(error.Error())
	}
	delete.Exec(customer.CustomerId)

	defer database.Close()
}


func main() {

	var customers []Customer
	customers = GetCustomers()
	//fmt.Println("Customers", customers)
	//fmt.Println("Before Insert",customers)
	//fmt.Println("Before Update", customers)
	fmt.Println("Before Delete", customers)

	//var customer Customer
	//customer.CustomerName = "Arnie Smith"
	//customer.SSN = "2386343"
	//InsertCustomer(customer)
	//customers = GetCustomers()
	//fmt.Println("After Insert", customers)

	//var customer Customer
	//customer.CustomerName = "George Thompson"
	//customer.SSN = "23233432"
	//customer.CustomerId = 5
	//UpdateCustomer(customer)
	//customers = GetCustomers()
	//fmt.Println("After Update", customers)

	var customer Customer
	customer.CustomerName = "George Thompson"
	customer.SSN = "23233432"
	customer.CustomerId = 5
	deleteCustomer(customer)
	customers = GetCustomers()
	fmt.Println("After Delete", customers)
}
