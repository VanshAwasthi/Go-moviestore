package models

import (
	"github.com/VanshAwasthi/go-movies/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&movies{})
}

func (b *Movies) Createmovies() *Movies {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllmovies() []Movies {
	var Books []Movies
	db.Find(&Movies)
	return Movies
}

func GetmoviesById(Id int64) (*Movies, *gorm.DB) {
	var getmovies Movies
	db := db.Where("ID = ?", Id).Find(&getmovies)
	return &getmovies, db
}

func Deletemovies(ID int64) Movies {
	var movies Movies
	db.Where("ID = ?", ID).Delete(movies)
	return movies
}
