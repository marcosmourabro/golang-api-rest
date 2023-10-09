package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marcosmourabro/go-rest-api/database"
	"github.com/marcosmourabro/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalitie
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func GetPersonalities(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var newPersonalitie models.Personalitie
	database.DB.First(&newPersonalitie, id)

	json.NewEncoder(w).Encode(newPersonalitie)
}

func CreatePersonalities(w http.ResponseWriter, r *http.Request) {
	var newPersonalities models.Personalitie
	json.NewDecoder(r.Body).Decode(&newPersonalities)
	database.DB.Create(&newPersonalities)
	json.NewEncoder(w).Encode(newPersonalities)
}

func DeletePersonalities(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalities models.Personalitie
	database.DB.Delete(&personalities, id)
	json.NewEncoder(w).Encode(personalities)
}

func UpdatePersonalities(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalities models.Personalitie
	database.DB.First(&personalities, id)
	json.NewDecoder(r.Body).Decode(&personalities)
	database.DB.Save(&personalities)
	json.NewEncoder(w).Encode(personalities)
}
