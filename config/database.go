package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "blimbeng38"
	dbname   = "go_products"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable ", host, port, user, password, dbname)


var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("postgres", psqlInfo)
	// defer db.Close()
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	DB = db
	log.Println("Database connected")

}

// func GetAll() []entities.Category {
// 	rows, err := config.DB.Query(`SELECT * FROM categories`)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer rows.Close()

// 	var categories []entities.Category

// 	for rows.Next() {
// 		var category entities.Category
// 		if err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
// 			panic(err)
// 		}

// 		categories = append(categories, category)
// 	}

// 	return categories
// }