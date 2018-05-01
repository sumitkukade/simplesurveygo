package structures

import (
	"gopkg.in/mgo.v2/bson"
)

type QuestionStruct struct {
	Que     string   `bson:"que"`
	Options []string `bson:"options"`
}

type ServeyDetails struct {
	Id          bson.ObjectId    `bson:"_id"`
	Title       string           `bson:"title"`
	Description string           `bson:"description"`
	Question    []QuestionStruct `bson:"question"`
	Status      string           `bson:"status"`
}
type QuestionAnswer struct {
	Que    string `json:"que"`
	Answer string `json:"answer"`
}

type PostServeyInput struct {
	Title    string           `json:"title" bson:"title"`
	Question []QuestionAnswer `json:"question" bson:"question"`
}

type PostServey struct {
	Username string           `bson:"username"`
	Title    string           `bson:"title"`
	Question []QuestionAnswer `bson:"question"`
}
