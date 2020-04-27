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
		HandleFunc("/users/{userID}/profile", UpdateUserInfoHandler).
		Methods("POST").
		Name("UpdateUserInfo")
	r.
		HandleFunc("/users/{userID}/profile", GetUserInfoHandler).
		Methods("GET").
		Name("GetUserInfo")
	r.
		HandleFunc("/users/{userID}/saved/foods/{foodID}", AddUserSavedFoodHandler).
		Methods("POST").
		Name("AddUserSavedFood")
	r.
		HandleFunc("/users/{userID}/saved/meals/{mealID}", AddUserSavedMealHandler).
		Methods("POST").
		Name("AddUserSavedMeal")
	r.
		HandleFunc("/users/{userID}/consumed/{mealID}", AddUserConsumedMealHandler).
		Methods("POST").
		Name("AddUserConsumedMeal")

		// FOODS
	r.
		HandleFunc("/foods", GetFoodsHandler).
		Methods("GET").
		Name("GetFoods")
	r.
		HandleFunc("/foods/saved", GetUserSavedFoodsHandler).
		Methods("GET").
		Name("GetUserSavedFoods")
	r.
		HandleFunc("/foods", CreateFoodHandler).
		Methods("POST").
		Name("CreateFood")
	r.
		HandleFunc("/foods/{foodID}", UpdateFoodHandler).
		Methods("POST").
		Name("UpdateFood")
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
		HandleFunc("/meals/saved", GetUserSavedMealsHandler).
		Methods("GET").
		Name("GetUserSavedMeals")
	r.
		HandleFunc("/meals/consumed", GetConsumedMealsHandler).
		Methods("GET").
		Name("GetConsumedMeals")
	r.
		HandleFunc("/meals", CreateMealHandler).
		Methods("POST").
		Name("CreateMeal")
	r.
		HandleFunc("/meals/{mealID}", UpdateMealHandler).
		Methods("POST").
		Name("UpdateMeal")
	r.
		HandleFunc("/meals/{mealID}", DeleteMealHandler).
		Methods("DELETE").
		Name("DeleteMeal")

	return r
}
