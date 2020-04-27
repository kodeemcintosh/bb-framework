package main

import (
	"database/sql"
	"net/http"
	"time"
)

type Environment string

type App struct {
	Env    Environment
	DB     *sql.DB
	Server http.Server
}

const (
	Production Environment = "PROD"
	Staging    Environment = "STAGE"
	Develoment Environment = "DEV"
)

type Food struct {
	ID            string
	Name          string
	Category      string
	Brand         string
	UnitOfMeasure string
	ServingSize   string
	Calories      int
	Fat           int
	Carb          int
	Protein       int
}

type FoodItem struct {
	MealID   string
	Food     Food
	Quantity int
}

type Meal struct {
	ID    string
	Name  string
	Foods []FoodItem
}
type ConsumedMeal struct {
	ID         string
	MealID     string
	UserID     string
	ConsumedAt time.Time
}

type SavedFood struct {
	ID            string
	UserID        string
	Name          string
	Category      string
	Brand         string
	UnitOfMeasure string
	ServingSize   string
	Calories      int
	Fat           int
	Carb          int
	Protein       int
}

type SavedMeal struct {
	MealID string
	UserID string
	Name   string
	Foods  []FoodItem
}

type User struct {
	ID         string
	Name       Name
	Address    Address
	Phone      string
	Email      string
	SavedMeals []Meal
}

type Name struct {
	FirstName string
	LastName  string
	Username  string
}

type Address struct {
	Street1    string
	Street2    string
	City       string
	State      string
	Province   string
	Country    string
	PostalCode string
}

//Route is a structure for the api router call
// type Route struct {
// 	Name        string
// 	Method      string
// 	Pattern     string
// 	HandlerFunc http.HandlerFunc
// }

// //Routes because Route objects have to go somewhere
// type Routes []Route

// type SubRoutes

// type User struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// type Role int
