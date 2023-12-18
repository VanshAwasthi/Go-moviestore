package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/VanshAwasthi/go-movies/pkg/models"
	"github.com/VanshAwasthi/go-movies/pkg/utils"
	"github.com/gorilla/mux"
)

var Newmovies models.Movies

func Getmovies(w http.ResponseWriter, r *http.Request) {
	newmovies := models.GetAllmovies()
	res, _ := json.Marshal(newmovies)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetmoviesById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	moviesId := vars["moviesId"]
	ID, err := strconv.ParseInt(moviesId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	moviesDetails, _ := models.GetmoviesById(ID)
	res, _ := json.Marshal(moviesDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Createmovies(w http.ResponseWriter, r *http.Request) {
	newmovies := &models.Movies{}
	err := utils.ParseBody(r, newmovies)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	createdmovies := newmovies.Createmovies()
	if createdmovies.ID == 0 {
		http.Error(w, "Failed to create movies", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(createdmovies)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func Deletemovies(w http.ResponseWriter, r *http.Request) {
	moviesId := mux.Vars(r)["moviesId"] // Extract moviesId directly from vars
	ID, err := strconv.ParseInt(moviesId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	movies := models.Deletemovies(ID)
	res, _ := json.Marshal(movies)
	w.Header().Set("Content-Type", "application/json") // Corrected content type
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Updatemovies(w http.ResponseWriter, r *http.Request) {
	var updatemovies = &models.Movies{}
	utils.ParseBody(r, updatemovies)
	vars := mux.Vars(r)
	moviesId := vars["moviesId"]
	ID, err := strconv.ParseInt(moviesId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	moviesDetails, db := models.GetmoviesById(ID)
	if updatemovies.Name != "" {
		moviesDetails.Name = updatemovies.Name
	}
	if updatemovies.Author != "" {
		moviesDetails.Author = updatemovies.Author
	}
	if updatemovies.Publication != "" {
		moviesDetails.Publication = updatemovies.Publication
	}
	db.Save(&moviesDetails)
	res, _ := json.Marshal(moviesDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
