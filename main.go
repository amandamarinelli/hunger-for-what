package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	Recipe struct {
		Name             string       `json:"name"`
		MinutesToPrepare uint32       `json:"minutes_to_prepare"`
		Ingredients      []Ingredient `json:"ingredients"`
	}

	Ingredient struct {
		Name     string `json:"name"`
		Quantity uint32 `json:"quantity"`
		Unit     string `json:"unit"`
	}
)

var recipes = Recipe{
	Name:             "pasta",
	MinutesToPrepare: 15,
	Ingredients: []Ingredient{
		{
			Name:     "pasta",
			Quantity: 1,
			Unit:     "package",
		},
		{
			Name:     "tomato sauce",
			Quantity: 1,
			Unit:     "package",
		},
	},
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/recipe", getRecipes).Methods(http.MethodGet)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}

func getRecipes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(recipes)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
