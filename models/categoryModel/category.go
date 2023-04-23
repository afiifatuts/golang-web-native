package categorymodel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Category {
	rows, err := config.DB.Query(`SELECT * FROM categories;`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			panic(err)
		}
		categories = append(categories, category)

	}
	return categories
}

func Create(category entities.Category) bool {
	var id int64
	//query row itu untuk mengembalikan 1 data yang kita butuhkan adalah id
	err := config.DB.QueryRow(`INSERT INTO categories (name, created_at, update_at) VALUES ($1,$2,$3) RETURNING id;`, category.Name, category.CreatedAt, category.UpdatedAt).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id > 0
}

// menerima params id dari controller edit kemudian mereturn struct
func Detail(id int) entities.Category {
	row := config.DB.QueryRow(`SELECT id, name FROM categories WHERE id = $1;`, id)
	var category entities.Category
	if err := row.Scan(&category.Id, &category.Name); err != nil {
		panic(err.Error())
	}
	return category

}

func Update(id int, category entities.Category) bool {
	query, err := config.DB.Exec(`UPDATE categories SET name = $2, update_at= $3 WHERE id = $1;`, id, category.Name, category.UpdatedAt)
	if err != nil {
		panic(err)
	}
	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}
	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec(`DELETE FROM categories WHERE id = $1;`, id)
	return err
}
