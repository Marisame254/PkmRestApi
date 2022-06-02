package models

import (
	"github.com/Marisame254/PkmRestApi/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Pokemon struct {
	//gorm.Model
	Pokemonid  int64  `json:"pokemonid"`
	Name       string `json:"name"`
	Type1id    string `json:"type1id"`
	Hp         int64  `json:"hp"`
	Attack     int64  `json:"attack"`
	Defense    int64  `json:"defense"`
	Sp_attack  int64  `json:"sp_attack"`
	Sp_defense int64  `json:"sp_defense"`
	Speed      int64  `json:"speed"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Pokemon{})
}

func (p *Pokemon) CreatePkm() *Pokemon {
	db.NewRecord(p)
	db.Create(&p)
	return p
}

func GetAllPkms() []Pokemon {
	var Pkms []Pokemon
	db.Find(&Pkms)
	return Pkms
}

func GetPkmById(Id int64) (*Pokemon, *gorm.DB) {
	var getPkm Pokemon
	db := db.Where("pokemonid = ?", Id).Find(&getPkm)
	return &getPkm, db
}

func DeletePkm(Id int64) Pokemon {
	var pokemon Pokemon
	db.Where("pokemonid = ?", Id).Delete(&pokemon)
	return pokemon
}
