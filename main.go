package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/terminate25/filmserver/config"

	. "github.com/terminate25/filmserver/api"
	. "github.com/terminate25/filmserver/dao"
)


func init() {
	var config = Config{}
	var dao = FilmsDAO{}

	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()

	InitDAOInstance(dao)

}


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/films", AllFilmsEndPoint).Methods("GET")
	r.HandleFunc("/films", CreateFilmEndPoint).Methods("POST")
	r.HandleFunc("/films", UpdateFilmEndPoint).Methods("PUT")
	r.HandleFunc("/films", DeleteFilmEndPoint).Methods("DELETE")
	r.HandleFunc("/films/{id}", FindFilmEndpoint).Methods("GET")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}
