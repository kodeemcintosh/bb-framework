package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Users Handlers
func CreateUserHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	fmt.Println(params)
	fmt.Fprintf(w, "Create User!")
}

func UpdateUserInfoHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userID := params["userID"]

	fmt.Println(params)
	fmt.Fprintf(w, "Update user info!")
}

func GetUserInfoHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userID := params["userID"]

	fmt.Println(params)
	fmt.Fprintf(w, "Get user info!")
}

func AddUserSavedFoodHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userID := params["userID"]
	foodID := params["foodID"]

	fmt.Println(params)
	fmt.Fprintf(w, "Add user saved food!")
}

func AddUserSavedMealHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userID := params["userID"]
	mealID := params["mealID"]

	fmt.Println(params)
	fmt.Fprintf(w, "Add user saved meal!")
}

func AddUserConsumedMealHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userID := params["userID"]
	mealID := params["mealID"]

	fmt.Println(params)
	fmt.Fprintf(w, "Add user consumed meal!")
}

// Foods Handlers
func GetFoodsHandler(w http.ResponseWriter, req *http.Request) {
	qry := req.URL.Query()

	name := qry["name"]

	fmt.Fprintf(w, "Get foods!")
}

func GetUserSavedFoodsHandler(w http.ResponseWriter, req *http.Request) {
	qry := req.URL.Query()

	userID := qry["userID"]
	page := qry["page"]
	count := qry["count"]
	sort := qry["sort"]

	fmt.Fprintf(w, "Get user saved foods!")
}

func CreateFoodHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "create foods!")
}

func UpdateFoodHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	foodID := params["foodID"]

	fmt.Fprintf(w, "update foods!")
}

func DeleteFoodHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	foodID := params["foodID"]

	fmt.Fprintf(w, "delete food!")
}

// Meals Handlers
func GetMealsHandler(w http.ResponseWriter, req *http.Request) {
	qry := req.URL.Query()

	userID := qry["userID"]
	page := qry["page"]
	count := qry["count"]
	sort := qry["sort"]

	fmt.Fprintf(w, "Get meals!")
}

func GetUserSavedMealsHandler(w http.ResponseWriter, req *http.Request) {
	qry := req.URL.Query()

	userID := qry["userID"]
	page := qry["page"]
	count := qry["count"]
	sort := qry["sort"]

	fmt.Fprintf(w, "Get user saved meals!")
}

func GetConsumedMealsHandler(w http.ResponseWriter, req *http.Request) {
	qry := req.URL.Query()

	userID := qry["userID"]
	startDate := qry["start"]
	endDate := qry["end"]

	fmt.Fprintf(w, "Get consumed meals!")
}

func CreateMealHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Create meal!")
}

func UpdateMealHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	mealID := params["mealID"]

	fmt.Fprintf(w, "Update meal!")
}

func DeleteMealHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	mealID := params["mealID"]

	fmt.Fprintf(w, "Delete meal!")
}
