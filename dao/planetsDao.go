package dao

import (
	"log"

	"github.com/wallacebenevides/star-wars-api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Database

type PlanetsDAO struct {
	Hosts    string
	Database string
}

const (
	COLLECTION = "planets"
)

func (m *PlanetsDAO) Connect() {

	session, err := mgo.Dial(m.Hosts)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func GetAllPlanets() ([]models.Planet, error) {
	var planets []models.Planet
	err := db.C(COLLECTION).Find(bson.M{}).All(&planets)
	return planets, err
}

func CreatePlanet(planet models.Planet) error {
	err := db.C(COLLECTION).Insert(&planet)
	return err
}

func GetPlanetByID(id string) (models.Planet, error) {
	var planet models.Planet
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&planet)
	return planet, err
}

func FindPlanetByName(name string) ([]models.Planet, error) {
	var planets []models.Planet
	filter := bson.D{{
		"name",
		bson.RegEx{Pattern: name},
	}}
	err := db.C(COLLECTION).Find(filter).All(&planets)
	return planets, err
}
