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

func GetPkm(w http.ResponseWriter, r *http.Request) {
	newPkms := models.GetAllPkms()
	res, _ := json.Marshal(newPkms)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPkmById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pkmId := vars["pkmId"]
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

func CreatePkm(w http.ResponseWriter, r *http.Request) {
	CreatePkm := &models.Pokemon{}
	utils.ParseBody(r, CreatePkm)
	p := CreatePkm.CreatePkm()
	res, _ := json.Marshal(p)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeletePkm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pkmId := vars["pkmId"]
	Id, err := strconv.ParseInt(pkmId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	pokemon := models.DeletePkm(Id)
	res, _ := json.Marshal(pokemon)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdatePkm(w http.ResponseWriter, r *http.Request) {
	var updatePkm = &models.Pokemon{}
	utils.ParseBody(r, updatePkm)
	vars := mux.Vars(r)
	pkmId := vars["pkmId"]
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

	if updatePkm.Defence != 0 {
		pkmDetails.Defence = updatePkm.Defence
	}

	if updatePkm.Sp_attack != 0 {
		pkmDetails.Sp_attack = updatePkm.Sp_attack
	}

	if updatePkm.Sp_defence != 0 {
		pkmDetails.Sp_defence = updatePkm.Sp_defence
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
