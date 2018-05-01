package dao

import (
	"simplesurveygo/structures"

	"gopkg.in/mgo.v2/bson"
)

func GetServeysByStatus(status string) []structures.ServeyDetails {
	session := MgoSession.Clone()
	defer session.Close()
	var servey []structures.ServeyDetails
	serveyClctn := session.DB("simplesurveys").C("servey")
	query := serveyClctn.Find(bson.M{"status": status})
	err := query.All(&servey)
	if err != nil {
		return []structures.ServeyDetails{}
	}
	return servey
}
func GetServeysByTitle(title string) structures.ServeyDetails {
	session := MgoSession.Clone()
	defer session.Close()
	var servey structures.ServeyDetails
	serveyClctn := session.DB("simplesurveys").C("servey")
	query := serveyClctn.Find(bson.M{"title": title})
	err := query.One(&servey)
	if err != nil {
		return structures.ServeyDetails{}
	}
	return servey
}
