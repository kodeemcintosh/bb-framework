package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Users Handlers
func CreateUserHandler(w http.ResponseWriter, req *http.Request) {
	var u User

	err := json.NewDecoder(req.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = InsertUser(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Create User!")
}

func UpdateUserInfoHandler(w http.ResponseWriter, req *http.Request) {
	var u User

	err := json.NewDecoder(req.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = UpdateUserInfo(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Update user info!")
}

func GetUserInfoHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	uID := params["userID"]

	u, err := SelectUserInfo(uID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Get user info!")
}

func CreateCustomFoodHandler(w http.ResponseWriter, req *http.Request) {
	var f Food

	err := json.NewDecoder(req.Body).Decode(&f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	params := mux.Vars(req)
	userID := params["userID"]

	err = InsertCustomFood(userID, f)

	fmt.Fprintf(w, "Add user saved food!")
}

func CreateConsumedMealHandler(w http.ResponseWriter, req *http.Request) {
	var cm ConsumedMeal

	err := json.NewDecoder(req.Body).Decode(&cm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = InsertConsumedMeal(cm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Add user consumed meal!")
}

// Foods Handlers
func GetFoodsHandler(w http.ResponseWriter, req *http.Request) {
	qry := req.URL.Query()

	name := qry["name"]
	sort := qry["sort"]
	page := qry["page"]
	max := qry["max"]

	foods, err := SelectFoods(name, sort, page, max)

	fmt.Fprintf(w, "Get foods!")
}

func GetCustomFoodsHandler(w http.ResponseWriter, req *http.Request) {
	qry := req.URL.Query()

	userID := qry["userID"]
	page := qry["page"]
	count := qry["count"]
	sort := qry["sort"]

	foods, err := SelectCustomFoods(userID, sort, page, max)

	fmt.Fprintf(w, "Get user saved foods!")
}

func CreateFoodHandler(w http.ResponseWriter, req *http.Request) {
	var f Food

	err := json.NewDecoder(req.Body).Decode(&f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = InsertFood(f)

	fmt.Fprintf(w, "create foods!")
}

func UpdateCustomFoodHandler(w http.ResponseWriter, req *http.Request) {
	var f Food

	err := json.NewDecoder(req.Body).Decode(&f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	params := mux.Vars(req)
	foodID := params["foodID"]

	err = UpdateCustomFood(f)
	fmt.Fprintf(w, "update foods!")
}

func DeleteFoodHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	foodID := params["foodID"]

	err := DeleteFood(foodID)

	fmt.Fprintf(w, "delete food!")
}

// Meals Handlers
func GetMealsHandler(w http.ResponseWriter, req *http.Request) {
	qry := req.URL.Query()

	userID := qry["userID"]
	page := qry["page"]
	count := qry["count"]
	sort := qry["sort"]

	meals, err := SelectMeals(userID)

	fmt.Fprintf(w, "Get meals!")
}

func GetFavoriteMealsHandler(w http.ResponseWriter, req *http.Request) {
	qry := req.URL.Query()

	userID := qry["userID"]
	page := qry["page"]
	count := qry["count"]
	sort := qry["sort"]

	meals, err := SelectFavoriteMeals(uID)

	fmt.Fprintf(w, "Get user saved meals!")
}

func UpdateFavoriteMealStatusHandler(w http.ResponseWriter, req *http.Request) {
	var mID string
	var isFavorite bool

	err := json.NewDecoder(req.Body).Decode(&mID, &isFavorite)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := UpdateFavoriteMealStatus(mID, isFavorite)

	fmt.Fprintf(w, "Update Favorite meals status meals!")
}

func GetConsumedMealsHandler(w http.ResponseWriter, req *http.Request) {
	qry := req.URL.Query()

	userID := qry["userID"]
	start := qry["start"]
	end := qry["end"]

	meals, err := SelectConsumedMeals(userID, start, end)

	fmt.Fprintf(w, "Get consumed meals!")
}

func UpdateConsumedMealsHandler(w http.ResponseWriter, req *http.Request) {
	var cm ConsumedMeal

	err := json.NewDecoder(req.Body).Decode(&cm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = UpdateConsumedMeals(cm)

	fmt.Fprintf(w, "Get consumed meals!")
}

func DeleteConsumedMealsHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	mID := params["mealID"]

	err = DeleteConsumedMeals(mID)

	fmt.Fprintf(w, "Get consumed meals!")
}

func CreateMealHandler(w http.ResponseWriter, req *http.Request) {
	var m Meal

	err := json.NewDecoder(req.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = InsertMeal(m)
	fmt.Fprintf(w, "Add user saved meal!")
}

func UpdateFoodItemHandler(w http.ResponseWriter, req *http.Request) {
	var fi FoodItem

	err := json.NewDecoder(req.Body).Decode(&fi)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = UpdateFoodItem(fi)

	fmt.Fprintf(w, "Update meal!")
}

func DeleteFoodItemHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	fiID := params["foodItemID"]

	err := DeleteFoodItem(fiID)

	fmt.Fprintf(w, "Delete meal!")
}

func DeleteMealHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	mID := params["mealID"]

	err := DeleteMeal(mID)

	fmt.Fprintf(w, "Delete meal!")
}
