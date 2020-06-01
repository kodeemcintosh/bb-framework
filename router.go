package main

import (
	// "net/http"

	"github.com/gorilla/mux"
)

//Router is the big, bad router that gets called in the main function to do the heavy lifting
func Router() *mux.Router {
	// create new router
	r := mux.NewRouter().StrictSlash(true)
	// 	Host("{subdomain:[a-z]+}.merchforce.io")
	// PathPrefix("/api/v1")

	// apply auth middleware
	// r.Use(middleware.OktaAuth)

	// USERS
	r.
		HandleFunc("/users", CreateUserHandler).
		Methods("POST").
		Name("CreateUser")
	r.
		HandleFunc("/users/{userID}", UpdateUserInfoHandler).
		Methods("POST").
		Name("UpdateUserInfo")
	r.
		HandleFunc("/users/{userID}", GetUserInfoHandler).
		Methods("GET").
		Name("GetUserInfo")
	r.
		HandleFunc("/foods/custom", CreateCustomFoodHandler).
		Methods("POST").
		Name("CreateCustomFood")

		// FOODS
	r.
		HandleFunc("/foods", GetFoodsHandler).
		Methods("GET").
		Name("GetFoods")
	r.
		HandleFunc("/foods/custom", GetCustomFoodsHandler).
		Methods("GET").
		Name("GetCustomFoods")
	r.
		HandleFunc("/foods", CreateFoodHandler).
		Methods("POST").
		Name("CreateFood")
	r.
		HandleFunc("/foods/{foodID}", UpdateCustomFoodHandler).
		Methods("POST").
		Name("UpdateCustomFood")
	r.
		HandleFunc("/foods/{foodID}", DeleteFoodHandler).
		Methods("DELETE").
		Name("DeleteFood")

		// MEALS
	r.
		HandleFunc("/meals", GetMealsHandler).
		Methods("GET").
		Name("GetMeals")
	r.
		HandleFunc("/meals/favorite", GetFavoriteMealsHandler).
		Methods("GET").
		Name("GetFavoriteMeals")
	r.
		HandleFunc("/meals/favorite", UpdateFavoriteMealStatusHandler).
		Methods("POST").
		Name("UpdateFavoriteMealStatus")
	r.
		HandleFunc("/meals/consumed", CreateConsumedMealHandler).
		Methods("POST").
		Name("CreateConsumedMeal")
	r.
		HandleFunc("/meals/consumed", GetConsumedMealsHandler).
		Methods("GET").
		Name("GetConsumedMeals")
	r.
		HandleFunc("/meals/consumed/{mealID}, UpdateConsumedMealsHandler).
		Methods("POST").
		Name("UpdateConsumedMeals")
	r.
		HandleFunc("/meals/consumed/{mealID}", DeleteConsumedMealsHandler).
		Methods("DELETE").
		Name("DeleteConsumedMeals")
	r.
		HandleFunc("/meals", CreateMealHandler).
		Methods("POST").
		Name("CreateMeal")
	r.
		HandleFunc("/meals/{mealID}", DeleteMealHandler).
		Methods("DELETE").
		Name("DeleteMeal")
	r.
		HandleFunc("/meals/food-item", UpdateFoodItemHandler).
		Methods("POST").
		Name("UpdateFoodItem")
	r.
		HandleFunc("/meals/food-item/{foodItemID}", DeleteFoodItemHandler).
		Methods("POST").
		Name("DeleteFoodItem")

	return r
}
