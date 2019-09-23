package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"

	. "github.com/terminate25/filmserver/models"

	. "github.com/terminate25/filmserver/dao"
)

var dao = FilmsDAO{}

// AllFilmsEndPoint Get list of films
func AllFilmsEndPoint(w http.ResponseWriter, r *http.Request) {
	films, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, films)
}

// FindFilmEndpoint Get film by its ID
func FindFilmEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	film, err := dao.FindByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid film id")
		return
	}
	respondWithJSON(w, http.StatusOK, film)
}

// CreateFilmEndPoint POST new film
func CreateFilmEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var film Film
	if err := json.NewDecoder(r.Body).Decode(&film); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	film.ID = bson.NewObjectId()
	if err := dao.Insert(film); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, film)
}

//UpdateFilmEndPoint PUT update an existing film
func UpdateFilmEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var film Film
	if err := json.NewDecoder(r.Body).Decode(&film); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(film); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, film)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

// DeleteFilmEndPoint DELETE an existing film
func DeleteFilmEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var film Film
	if err := json.NewDecoder(r.Body).Decode(&film); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(film); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"error": "success"})
}

// respondWithJSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}


func InitDAOInstance(filmsDAO FilmsDAO) {
	dao = filmsDAO
}
