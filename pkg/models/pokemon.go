package models

import (
	"github.com/Marisame254/PkmRestApi/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Pokemon struct {
	gorm.Model
	Name       string `json:"name"`
	Type       string `json:"type"`
	Hp         int    `json:"hp"`
	Attack     int    `json:"attack"`
	Defence    int    `json:"defence"`
	Sp_attack  int    `json:"sp_attack"`
	Sp_defence int    `json:"sp_defence"`
	Speed      int    `json:"speed"`
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
	var Pokemons []Pokemon
	db.Find(&Pokemons)
	return Pokemons
}

func GetPkmById(Id int64) (*Pokemon, *gorm.DB) {
	var getPkm Pokemon
	db := db.Where("Id = ?", Id).Find(&getPkm)
	return &getPkm, db
}

func DeletePkm(Id int64) Pokemon {
	var pokemon Pokemon
	db.Where("Id = ?", Id).Delete(&pokemon)
	return pokemon
}
