package main

import (
	"fmt"
	"gorm/dao"
	"gorm/entity"
	"gorm/util"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*gorm2*/
func main() {

	db := util.GetDB()

	defer db.Close()

	db.AutoMigrate(&entity.Customer{}, &entity.Order{})

	customerDAO := dao.CustomerDAOImpl{}

	// Create a new customer
	newCustomer := entity.Customer{
		Name:    "John Doe",
		Address: "123 Main St",
		Email:   "john.doe@example.com",
	}
	createdCustomer, err := customerDAO.CreateCustomer(db, newCustomer)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Created customer:", createdCustomer)

	// Get a customer by ID
	customer, err := customerDAO.GetCustomer(db, createdCustomer.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Retrieved customer:", customer)

	// Get all customers
	customers, err := customerDAO.GetAllCustomers(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Retrieved customers:", customers)

	// Update a customer
	customer.Name = "Jane Doe"
	updatedCustomer, err := customerDAO.UpdateCustomer(db, customer)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Updated customer:", updatedCustomer)

	// Delete a customer
	err = customerDAO.DeleteCustomer(db, customer)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Deleted customer:", customer)
}
