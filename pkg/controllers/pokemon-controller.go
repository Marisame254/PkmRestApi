package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Marisame254/PkmRestApi/pkg/models"
	"github.com/Marisame254/PkmRestApi/pkg/utils"
	"github.com/gorilla/mux"
)

var NewPkm models.Pokemon

// Function to Get all pokemons ================================
func GetPkm(w http.ResponseWriter, r *http.Request) {
	newPkms := models.GetAllPkms()
	res, _ := json.Marshal(newPkms)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Function to get a pokemon by Id ============================
func GetPkmById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pkmId := vars["Id"]
	Id, err := strconv.ParseInt(pkmId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	pkmDetails, _ := models.GetPkmById(Id)
	res, _ := json.Marshal(pkmDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Function to Create a pokemon ==============================
func CreatePkm(w http.ResponseWriter, r *http.Request) {
	CreatePkm := &models.Pokemon{}
	utils.ParseBody(r, CreatePkm)
	p := CreatePkm.CreatePkm()
	res, _ := json.Marshal(p)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Function to delete a pokemon =============================
func DeletePkm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pkmId := vars["Id"]
	Id, err := strconv.ParseInt(pkmId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	p := models.DeletePkm(Id)
	res, _ := json.Marshal(p)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Function to update a pokemon =============================
func UpdatePkm(w http.ResponseWriter, r *http.Request) {
	var updatePkm = &models.Pokemon{}
	utils.ParseBody(r, updatePkm)
	vars := mux.Vars(r)
	pkmId := vars["Id"]
	Id, err := strconv.ParseInt(pkmId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	pkmDetails, db := models.GetPkmById(Id)
	if updatePkm.Name != "" {
		pkmDetails.Name = updatePkm.Name
	}

	if updatePkm.Type != "" {
		pkmDetails.Type = updatePkm.Type
	}

	if updatePkm.Hp != 0 {
		pkmDetails.Hp = updatePkm.Hp
	}

	if updatePkm.Attack != 0 {
		pkmDetails.Attack = updatePkm.Attack
	}

	if updatePkm.Defense != 0 {
		pkmDetails.Defense = updatePkm.Defense
	}

	if updatePkm.Sp_attack != 0 {
		pkmDetails.Sp_attack = updatePkm.Sp_attack
	}

	if updatePkm.Sp_defense != 0 {
		pkmDetails.Sp_defense = updatePkm.Sp_defense
	}

	if updatePkm.Speed != 0 {
		pkmDetails.Speed = updatePkm.Speed
	}

	db.Save(&pkmDetails)
	res, _ := json.Marshal(pkmDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
