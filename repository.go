package main

import (
	_ "github.com/lib/pq"

	"time"
)

func InsertUser(u User) {
	sql := `
		INSERT INTO users(id, first_name, last_name, user_name, phone, email, street1, street2, city, state, providence, country, postal_code, created_at, last_modified)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16);
	`

	_, err := app.DB.Exec(sql, u.ID, u.Name.FirstName, u.Name.LastName, u.Name.Username, u.Phone, u.Email, u.Address.Street1, u.Address.Street2, u.Address.City, u.Address.State, u.Address.Province, u.Address.Country, u.Address.PostalCode, time.Now, time.Now)
	if err != nil {
		panic(err)
	}
}

func UpdateUser() {

}

func GetUserInfo() {

}

func AddUserSavedFood(f Food) {
	sql := `
		INSERT INTO saved_foods(id, user_id, name, brand, category, unit_of_measure, serving_size, calories, fat_grams, carb_grams, protein_grams, created_at, last_modified)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);
	`

	_, err := app.DB.Exec(sql, f.ID, f.Name, f.Brand, f.Category, f.UnitOfMeasure, f.ServingSize, f.Calories, f.Fat, f.Carb, f.Protein, time.Now, time.Now)
	if err != nil {
		panic(err)
	}
}

func AddUserSavedMeal() {
	sql := ``

	_, err := app.DB.Exec(sql)
	if err != nil {
		panic(err)
	}
}

func AddUserConsumedMeal() {
	sql := ``

	_, err := app.DB.Exec(sql)
	if err != nil {
		panic(err)
	}
}

func GetFoods() {

}

func GetUserSavedFoods() {

}

func CreateFood(f Food) {
	sql := `
		INSERT INTO foods(id, name, brand, category, unit_of_measure, serving_size, calories, fat_grams, carb_grams, protein_grams, created_at, last_modified)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);
	`

	_, err := app.DB.Exec(sql, f.ID, f.Name, f.Brand, f.Category, f.UnitOfMeasure, f.ServingSize, f.Calories, f.Fat, f.Carb, f.Protein, time.Now, time.Now)
	if err != nil {
		panic(err)
	}
}

func UpdateFood() {

}

func DeleteFood() {

}

func GetMeals() {

}

func GetUserSavedMeals() {

}

func GetConsumedMeals() {

}

func InsertMeal(m Meal) {
	tx, err := app.DB.Begin()
	if err != nil {
		panic(err)
	}

	{
		sql := `
			INSERT INTO meals(id, name, created_at, last_modified)
			VALUES($1, $2, $3, $4);
			`
		stmt, err := tx.Prepare(sql)
		if err != nil {
			tx.Rollback()
			panic(err)
		}
		defer stmt.Close()

		if _, err := stmt.Exec(m.ID, m.Name, time.Now(), time.Now()); err != nil {
			tx.Rollback() // return an error too, we may want to wrap them
			panic(err)
		}
	}

	for _, f := range m.Foods {
		sql := `
			INSERT INTO food_items(meal_id, food_id, quantity, created_at, last_modified)
			VALUES($1, $2, $3, $4, $5, $6);
		`
		stmt, err := tx.Prepare(sql)
		if err != nil {
			tx.Rollback()
			panic(err)
		}
		defer stmt.Close()

		if _, err := stmt.Exec(f.MealID, f.Food.ID, f.Quantity, time.Now(), time.Now()); err != nil {
			tx.Rollback() // return an error too, we may want to wrap them
			panic(err)
		}
	}

	tx.Commit()
}

func UpdateMeal() {

}

func DeleteMeal() {

}
