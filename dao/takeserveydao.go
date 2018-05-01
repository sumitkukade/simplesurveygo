package dao

import (
	"simplesurveygo/structures"

	"gopkg.in/mgo.v2/bson"
)

func CreateAnswerForUser(ans structures.PostServeyInput, username string) bool {
	session := MgoSession.Clone()
	defer session.Close()
	serveyStruct := structures.PostServey{Username: username, Title: ans.Title, Question: ans.Question}

	sessionClctn := session.DB("simplesurveys").C("userservey")
	sessionClctn.Insert(serveyStruct)
	return true
}

func GetAllUserServeys(username string) []structures.PostServey {
	session := MgoSession.Clone()
	defer session.Close()
	var userserveys []structures.PostServey
	qaClctn := session.DB("simplesurveys").C("userservey")
	qaClctn.Find(bson.M{"username": username}).All(&userserveys)
	return userserveys
}
