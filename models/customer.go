package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	FirstName string `gorm:"size:50;not null;" json:"first_name"`
	LastName  string `gorm:"size:50;not null;" json:"last_name"`
	Phone     string `gorm:"size:16;not null;" json:"phone"`
	Address   string `gorm:"size:100;not null;" json:"address"`
	City      string `gorm:"size:50;not null;" json:"city"`
	State     string `gorm:"size:50;not null;" json:"state"`
	ZipCode   string `gorm:"size:10;not null;" json:"zip_code"`
}

// Find a customer by id
func FindCustomer(id uint) (c Customer, err error) {
	err = DB.First(&c, id).Error
	return
}

// Retrieve all customers
func GetCustomers() (c []Customer, err error) {
	err = DB.Find(&c).Error
	return
}

// Create Customer with all fields
func (c Customer) AddCustomer() (err error) {
	err = DB.Create(&c).Error
	return
}

// Update all fields on customer by id
func (c Customer) UpdateCustomer() (err error) {
	err = DB.Save(&c).Error
	return
}

// Delete customer by id
func (c Customer) DeleteCustomer() (err error) {
	err = DB.Delete(&c).Error
	return
}
