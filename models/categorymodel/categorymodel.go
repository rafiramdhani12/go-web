package categorymodel

import (
	"go_web/config"
	"go_web/entities"
)

func GetAll() []entities.Category{
	rows , err := config.DB.Query(`SELECT * FROM categories`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category

	for rows.Next(){
		var category entities.Category
		if err := rows.Scan(&category.ID,&category.Name,&category.CreatedAt,&category.UpdatedAt); err != nil{
			panic(err)
		}
		categories = append(categories, category)
	}
	return categories
}

func Create(category entities.Category) bool {
	result, err := config.DB.Exec(`
	INSERT INTO CATEGORIES (name,created_at,updated_at)
	VALUE (?,?,?)`,category.Name,category.CreatedAt,category.UpdatedAt)

	if err != nil{
		panic(err)
	}

	lastInsertId,err := result.LastInsertId()
	if err != nil{
		panic(err)
	}

	return lastInsertId > 0
}

func Detail(id int) entities.Category{
	row := config.DB.QueryRow(`SELECT id,name FROM categories WHERE id = ?`,id)

	var category entities.Category
	if err := row.Scan(&category.ID,&category.Name); err != nil{
		panic(err.Error())
	}
	return category
}

func Update(id int , Category entities.Category) bool {
	Query,err  := config.DB.Exec(`UPDATE categories SET name = ?, updated_at = ? WHERE id = ?`, Category.Name , Category.UpdatedAt , id)
	if err != nil{
		panic(err)
	}

	result, err := Query.RowsAffected()
	if err != nil{
		panic(err)
	}

	return result > 0
}

func Delete(id int) error{
	_,err :=config.DB.Exec(`DELETE FROM categories WHERE id = ?` , id)
	return err
}
