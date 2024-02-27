package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       string
	Fname    string
	Lname    string
	Username string
	Avatar   string
	Email    string
}

func main() {
	dsn := "root:123456789@tcp(127.0.0.1:3306)/go_orm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Create a User
	user := User{
		Id:       "your_id_here",
		Fname:    "John",
		Lname:    "Doe",
		Username: "johndoe",
		Avatar:   "avatar_url_here",
		Email:    "john@example.com",
	}
	db.Create(&user)

	// Read
	var fetchedUser User
	db.First(&fetchedUser, 1)                        // find user with integer primary key
	db.First(&fetchedUser, "id = ?", "your_id_here") // find user with specific id

	// Update - update user's Fname to "Jane"
	db.Model(&fetchedUser).Update("Fname", "Jane")

	// Delete - delete user
	db.Delete(&fetchedUser, 1)
}
