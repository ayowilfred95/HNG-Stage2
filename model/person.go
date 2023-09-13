package model


type Person struct {
	ID string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Hobby string `json:"hobby" bson:"hobby"`
}
