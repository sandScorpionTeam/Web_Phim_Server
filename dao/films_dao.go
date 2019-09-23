package dao

import (
	"log"

	. "github.com/terminate25/filmserver/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// FilmsDAO instance
type FilmsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

// COLLECTION films collection name
const (
	COLLECTION = "films"
)

//Connect Establish a connection to database
func (m *FilmsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {

		log.Fatal(err)
	}

	db = session.DB(m.Database)
}

// FindAll Find list films
func (m *FilmsDAO) FindAll() ([]Film, error) {
	var films []Film
	err := db.C(COLLECTION).Find(bson.M{}).All(&films)
	return films, err
}

// FindByID find a film by its id
func (m *FilmsDAO) FindByID(id string) (Film, error) {
	var films Film
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&films)
	return films, err
}

// Insert insert film into database
func (m *FilmsDAO) Insert(film Film) error {
	err := db.C(COLLECTION).Insert(&film)
	return err
}

// Delete delete an existing film
func (m *FilmsDAO) Delete(film Film) error {
	err := db.C(COLLECTION).Remove(&film)
	return err
}

// Update an existing film
func (m *FilmsDAO) Update(film Film) error {
	err := db.C(COLLECTION).UpdateId(film.ID, &film)
	return err
}
