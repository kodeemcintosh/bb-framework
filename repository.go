package main

import (
	_ "github.com/lib/pq"

	"time"
)

func InsertUser(u User) error {
	sql := `
		INSERT INTO users(first_name, last_name, user_name, phone, email, street1, street2, city, state, providence, country, postal_code, created_at, last_modified)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14);
	`

	_, err := app.DB.Exec(sql, u.ID, u.Name.FirstName, u.Name.LastName, u.Name.Username, u.Phone, u.Email, u.Address.Street1, u.Address.Street2, u.Address.City, u.Address.State, u.Address.Province, u.Address.Country, u.Address.PostalCode, time.Now, time.Now)
	if err != nil {
		panic(err)
		return err
	}

	return nil
}

func UpdateUserInfo(u User) error {
	sql := `
		UPDATE users SET first_name = $1, last_name = $2, user_name = $3, phone = $4, email = $5, street1 = $6, street2 = $7, city = $8, state = $9, providence = $10, country = $11, postal_code = $12, last_modified = $13)
		WHERE id = $14;
	`

	_, err := app.DB.Exec(sql, u.Name.FirstName, u.Name.LastName, u.Name.Username, u.Phone, u.Email, u.Address.Street1, u.Address.Street2, u.Address.City, u.Address.State, u.Address.Province, u.Address.Country, u.Address.PostalCode, time.Now, u.ID)
	if err != nil {
		panic(err)
		return err
	}

	return nil
}

func SelectUserInfo(uID string) (User, error) {
	var u User
	var a Address
	var n Name

	sql := `
		SELECT first_name, last_name, user_name, phone, email, street1, street2, city, state, province, country, postal_code FROM user
		WHERE id = $1;
	`

	err := app.DB.QueryRow(sql, uID).Scan(&n.FirstName, &n.LastName, &n.Username, &u.Phone, &u.Email, &a.Street1, &a.Street2, &a.City, &a.State, &a.Province, &a.Country, &a.PostalCode)
	if err != nil {
		panic(err)
		return User{}, err
	}

	u.Name = n
	u.Address = a

	return u, nil
}

func InsertCustomFood(uID string, f Food) error {
	sql := `
		INSERT INTO custom_foods(user_id, name, description, brand, category, unit_of_measure, serving_size, calories, fat, carb, protein, created_at, last_modified)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);
	`

	_, err := app.DB.Exec(sql, uID, f.Name, f.Description, f.Brand, f.Category, f.UnitOfMeasure, f.ServingSize, f.Calories, f.Fat, f.Carb, f.Protein, time.Now, time.Now)
	if err != nil {
		panic(err)
		return err
	}

	return nil
}

func UpdateUserFavoriteFoodStatus(fID string, isFavorite bool) error {
	sql := `
		UPDATE custom_foods SET favorite = $1, last_modified = $2
		WHERE id = $3;
	`

	_, err := app.DB.Exec(sql, isFavorite, time.Now, fID)
	if err != nil {
		panic(err)
		return err
	}

	return nil
}

func UpdateFavoriteMealStatus(mID string, isFavorite bool) error {
	sql := `
		UPDATE meals SET favorite = $1, last_modified = $2
		WHERE id = $3;
	`

	_, err := app.DB.Exec(sql, isFavorite, time.Now, mID)
	if err != nil {
		panic(err)
		return err
	}

	return nil
}

func InsertConsumedMeal(cm ConsumedMeal) error {
	sql := `
		INSERT INTO consumed_meals(meal_id, user_id, consumed_at, created_at, last_modified)
		VALUES($1, $2, $3, $4, $5);
	`

	_, err := app.DB.Exec(sql, cm.MealID, cm.UserID, cm.ConsumedAt, time.Now, time.Now)
	if err != nil {
		panic(err)
		return err
	}

	return nil
}

func UpdateConsumedMeal(cm ConsumedMeal) error {
	sql := `
		UPDATE consumed_meals SET meal_id = $1, consumed_at = $2, last_modified = $3
		WHERE id = $4;
	`

	_, err := app.DB.Exec(sql, cm.MealID, cm.ConsumedAt, time.Now, cm.ID)
	if err != nil {
		panic(err)
		return err
	}

	return nil
}

func DeleteConsumedMeal(mID string) error {
	sql := `
		DELETE FROM consumed_meals WHERE id = $1;
	`

	_, err := app.DB.Exec(sql, mID)
	if err != nil {
		panic(err)
		return err
	}

	return nil
}

func SelectFoods(name string, sort string, page int, max int) ([]Food, error) {
	var foods []Food

	// TODO: Add sort, page, and max
	sql := `
		SELECT id, name, description, brand, category, unit_of_measure, serving_size, calories, fat, carb, protein FROM foods
		SORT $1;
	`

	rows, err := app.DB.Query(sql, sort)
	if err != nil {
		panic(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var f Food

		err = rows.Scan(&f.ID, &f.Name, &f.Description, &f.Brand, &f.Category, &f.UnitOfMeasure, &f.ServingSize, &f.Calories, &f.Fat, &f.Carb, &f.Protein)
		if err != nil {
			panic(err)
			return nil, err
		}

		foods = append(foods, f)
	}

	return foods, nil
}

func SelectCustomFoods(uID string, sort string, page int, max int) ([]Food, error) {
	var foods []Food

	sql := `
		SELECT id, name, description, brand, category, unit_of_measure, serving_size, calories, fat, carb, protein FROM foods
		WHERE user_id = $1;
	`

	// TODO: Add sort, page, and max
	rows, err := app.DB.Query(sql, uID)
	if err != nil {
		panic(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var f Food

		err = rows.Scan(&f.ID, &f.Name, &f.Description, &f.Brand, &f.Category, &f.UnitOfMeasure, &f.ServingSize, &f.Calories, &f.Fat, &f.Carb, &f.Protein)
		if err != nil {
			panic(err)
			return nil, err
		}

		foods = append(foods, f)
	}

	return foods, nil
}

func GetUserFavoriteCustomFoods(uID string) ([]Food, error) {
	var foods []Food

	sql := `
		SELECT id, name, description, brand, category, unit_of_measure, serving_size, calories, fat, carb, protein FROM foods
		WHERE user_id = $1 AND favorite = true;
	`

	rows, err := app.DB.Query(sql, uID)
	if err != nil {
		panic(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var f Food

		err = rows.Scan(&f.ID, &f.Name, &f.Description, &f.Brand, &f.Category, &f.UnitOfMeasure, &f.ServingSize, &f.Calories, &f.Fat, &f.Carb, &f.Protein)
		if err != nil {
			panic(err)
			return nil, err
		}

		foods = append(foods, f)
	}

	return foods, nil
}

func InsertFood(f Food) error {
	sql := `
		INSERT INTO foods(name, description, brand, category, unit_of_measure, serving_size, calories, fat, carb, protein, created_at, last_modified)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);
	`

	_, err := app.DB.Exec(sql, f.Name, f.Description, f.Brand, f.Category, f.UnitOfMeasure, f.ServingSize, f.Calories, f.Fat, f.Carb, f.Protein, time.Now, time.Now)
	if err != nil {
		panic(err)
		return err
	}

	return nil
}

func UpdateCustomFood(f Food) error {
	sql := `
		UPDATE custom_foods SET name = $1, description = $2, brand = $3, category = $4, unit_of_measure = $5, serving_size = $6, calories = $7, fat = $8, carb = $9, protein = $10, last_modified = $11
		WHERE id = $12;
	`

	_, err := app.DB.Exec(sql, f.Name, f.Description, f.Brand, f.Category, f.UnitOfMeasure, f.ServingSize, f.Calories, f.Fat, f.Carb, f.Protein, time.Now, f.ID)
	if err != nil {
		panic(err)
		return err
	}

	return nil
}

func DeleteFood(fID string) error {
	sql := `
		DELETE FROM custom_foods WHERE id = $1;
	`

	_, err := app.DB.Exec(sql, fID)
	if err != nil {
		panic(err)
		return err
	}

	return nil
}

func SelectMeals(uID string) ([]Meal, error) {
	var meals []Meal

	// TODO: Add sort, page, and max
	sql := `
		SELECT id, name, description FROM meals
		WHERE user_id = $1;
	`

	rows, err := app.DB.Query(sql, uID)
	if err != nil {
		panic(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var m Meal

		err = rows.Scan(&m.ID, &m.Name, &m.Description)
		if err != nil {
			panic(err)
			return nil, err
		}

		meals = append(meals, m)
	}

	for _, m := range meals {
		sql := `
			SELECT f.id, f.name, f.description, f.brand, f.category, f.unit_of_measure, f.serving_size, f.calories, f.fat, f.carb, f.protein, fi.id, fi.quantity FROM foods f
				LEFT JOIN food_items fi ON f.id = fi.food_id
				LEFT JOIN meals m ON fi.meal_id = m.id
			WHERE m.id = $1;
		`
		rows, err = app.DB.Query(sql, m.ID)
		if err != nil {
			panic(err)
			return nil, err
		}

		for rows.Next() {
			var f Food
			var fi FoodItem

			err = rows.Scan(&f.ID, &f.Name, &f.Description, &f.Brand, &f.Category, &f.UnitOfMeasure, &f.ServingSize, &f.Calories, &f.Fat, &f.Carb, &f.Protein, &fi.ID, &fi.Quantity)
			if err != nil {
				panic(err)
				return nil, err
			}

			fi.MealID = m.ID
			fi.Food = f

			m.Foods = append(m.Foods, fi)
		}

		defer rows.Close()
	}

	return meals, nil
}

func SelectFavoriteMeals(uID string) ([]Meal, error) {
	var meals []Meal

	// TODO: Add sort, page, and max
	sql := `
		SELECT id, name, description FROM meals
		WHERE user_id = $1 AND favorite = true;
	`

	rows, err := app.DB.Query(sql, uID)
	if err != nil {
		panic(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var m Meal

		err = rows.Scan(&m.ID, &m.Name, &m.Description)
		if err != nil {
			panic(err)
			return nil, err
		}

		meals = append(meals, m)
	}

	for _, m := range meals {
		sql := `
			SELECT f.id, f.name, f.description, f.brand, f.category, f.unit_of_measure, f.serving_size, f.calories, f.fat, f.carb, f.protein, fi.id, fi.quantity FROM foods f
				LEFT JOIN food_items fi ON f.id = fi.food_id
				LEFT JOIN meals m ON fi.meal_id = m.id
			WHERE m.id = $1;
		`
		rows, err = app.DB.Query(sql, m.ID)
		if err != nil {
			panic(err)
			return nil, err
		}

		for rows.Next() {
			var f Food
			var fi FoodItem

			err = rows.Scan(&f.ID, &f.Name, &f.Description, &f.Brand, &f.Category, &f.UnitOfMeasure, &f.ServingSize, &f.Calories, &f.Fat, &f.Carb, &f.Protein, &fi.ID, &fi.Quantity)
			if err != nil {
				panic(err)
				return nil, err
			}

			fi.MealID = m.ID
			fi.Food = f

			m.Foods = append(m.Foods, fi)
		}

		defer rows.Close()
	}

	return meals, nil
}

func SelectConsumedMeals(uID string, start time.Time, end time.Time) ([]Meal, error) {
	var meals []Meal

	// TODO: Add sort, page, and max
	sql := `
		SELECT m.id, m.name, m.description FROM consumed_meals cm
			LEFT JOIN meals m ON cm.meal_id = m.id
		WHERE user_id = $1 AND m.consumed_at BETWEEN $2 AND $3;
	`

	rows, err := app.DB.Query(sql, uID, start, end)
	if err != nil {
		panic(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var m Meal

		err = rows.Scan(&m.ID, &m.Name, &m.Description)
		if err != nil {
			panic(err)
			return nil, err
		}

		meals = append(meals, m)
	}

	for _, m := range meals {
		sql := `
			SELECT f.id, f.name, f.description, f.brand, f.category, f.unit_of_measure, f.serving_size, f.calories, f.fat, f.carb, f.protein, fi.id, fi.quantity FROM foods f
				LEFT JOIN food_items fi ON f.id = fi.food_id
				LEFT JOIN meals m ON fi.meal_id = m.id
			WHERE m.id = $1;
		`
		rows, err = app.DB.Query(sql, m.ID)
		if err != nil {
			panic(err)
			return nil, err
		}

		for rows.Next() {
			var f Food
			var fi FoodItem

			err = rows.Scan(&f.ID, &f.Name, &f.Description, &f.Brand, &f.Category, &f.UnitOfMeasure, &f.ServingSize, &f.Calories, &f.Fat, &f.Carb, &f.Protein, &fi.ID, &fi.Quantity)
			if err != nil {
				panic(err)
				return nil, err
			}

			fi.MealID = m.ID
			fi.Food = f

			m.Foods = append(m.Foods, fi)
		}

		defer rows.Close()
	}

	return meals, nil
}

func InsertMeal(m Meal) error {
	tx, err := app.DB.Begin()
	if err != nil {
		panic(err)
		return err
	}

	{
		sql := `
			INSERT INTO meals(user_id, name, description, created_at, last_modified)
			VALUES($1, $2, $3, $4, $5);
			`
		stmt, err := tx.Prepare(sql)
		if err != nil {
			tx.Rollback()
			panic(err)
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(m.ID, m.UserID, m.Name, m.Description, time.Now(), time.Now()); err != nil {
			tx.Rollback() // return an error too, we may want to wrap them
			panic(err)
			return err
		}
	}

	for _, f := range m.Foods {
		sql := `
			INSERT INTO food_items(meal_id, food_id, quantity, created_at, last_modified)
			VALUES($1, $2, $3, $4, $5);
		`
		stmt, err := tx.Prepare(sql)
		if err != nil {
			tx.Rollback()
			panic(err)
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(f.MealID, f.Food.ID, f.Quantity, time.Now(), time.Now()); err != nil {
			tx.Rollback() // return an error too, we may want to wrap them
			panic(err)
			return err
		}
	}

	tx.Commit()

	return nil
}

func UpdateFoodItem(fi FoodItem) error {
	sql := `
		UPDATE food_items SET quantity = $1
		WHERE id = $2;
	`

	_, err := app.DB.Exec(sql, fi.Quantity, fi.ID)
	if err != nil {
		panic(err)
		return err
	}

	return nil
}

func DeleteFoodItem(fiID string) error {
	sql := `
		DELETE FROM food_items
		WHERE id = $1;
	`

	_, err := app.DB.Exec(sql, fiID)
	if err != nil {
		panic(err)
		return err
	}

	return nil
}

func DeleteMeal(mID string) error {
	sql := `
		DELETE FROM meals
		WHERE id = $1;
	`

	_, err := app.DB.Exec(sql, mID)
	if err != nil {
		panic(err)
		return err
	}

	return nil
}
