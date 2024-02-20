package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

)

func main() {
	// Get environment variables
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Construct data source name (DSN)
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)

	// Open a new database connection
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
	}

	// Connection successful, you can execute queries here
	fmt.Println("Connected to MySQL database successfully!")
	//get:user - retrieve all users
	//get: users/{id}-retrieve a specific use by ID
	//post: USER-create a new user
	//put:users/{id} - update existing	 by id
	// delete/users/{id}-delete a specific use by
	type User struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
	db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/database")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


	func getUsers(w http.ResponseWriter, r *http.Request) {
		var users []User
		rows, err := db.Query("SELECT id, first_name, last_name, email FROM users")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
	
		for rows.Next() {
			var user User
			err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, user)
		}
	
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}
