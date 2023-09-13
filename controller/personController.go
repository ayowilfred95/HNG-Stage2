package controller

import (
	"fmt"

	"github.com/ayowilfred95/database"
	"github.com/ayowilfred95/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type person struct {
	Name string `json:"name" bson:"name"`
	Hobby string `json:"hobby" bson:"hobby"`
}

func CreatePerson(c *fiber.Ctx) error {
	// validate the body
	body := new(person)
	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}
	// create the person
	collection := database.GetDBCollection("persons")
	result, err := collection.InsertOne(c.Context(), body)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to create person",
			"message": err.Error(),
		})
	}

	// return the person
	return c.Status(201).JSON(fiber.Map{
		"result": result,
	})
}

func GetPersons(c *fiber.Ctx) error {
	coll := database.GetDBCollection("persons")

	// find all persons
	persons := make([]model.Person, 0)
	cursor, err := coll.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// iterate over the curosr
	for cursor.Next(c.Context()) {
		person := model.Person{}
		err := cursor.Decode(&person)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		persons = append(persons, person)
	}
	return c.Status(200).JSON(fiber.Map{
		"data": persons,
	})
}

func GetPerson(c *fiber.Ctx) error {
	coll := database.GetDBCollection("persons")

	// Get the input (either by ID or name)
	// find the person by name (dynamic input)
	userId := c.Params("userId")
	fmt.Printf("Received userId: %s\n", userId)
	if userId == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Please provide either name or Id",
		})
	}

	person := model.Person{}

	// Attempt to convert the userId to a valid id
	objectID, err := primitive.ObjectIDFromHex(userId)

	var queryFilter bson.M

	if err != nil {
		// If it's not a valid ObjectID, query by "name" instead
		queryFilter = bson.M{
			"$or": []bson.M{
				{"name": userId},
				{"hobby":userId},
			},
		}
	} else {
		// If it's a valid ObjectID, query by "_id"
		queryFilter = bson.M{"_id": objectID}
	}

	err = coll.FindOne(c.Context(), queryFilter).Decode(&person)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).JSON(fiber.Map{
				"error":   "person not found",
				"message": err.Error(),
			})

		}
		fmt.Printf("Error while querying the database: %v\n", err)
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// return the person
	return c.Status(200).JSON(fiber.Map{
		"data": person,
	})
}

// create a struct to hold our data while being updated in the database
type updatePerson struct {
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Hobby string `json:"hobby,omitempty" bson:"hobby,omitempty"`
}

func UpdatePerson(c *fiber.Ctx) error {
	// validate the body
	body := new(updatePerson)
	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid body",
		})
	}

	// get the name
	userId := c.Params("userId")
	fmt.Printf("Received userId: %s\n", userId)
	if userId == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "please provide either name, Id or hobby",
		})

	}
	coll := database.GetDBCollection("persons")
	var result *mongo.UpdateResult
	var err error
	// Attempt to convert the userId to a valid ObjectID
	objectID, err := primitive.ObjectIDFromHex(userId)
	var queryFilter bson.M

	if err != nil {
		// If it's not a valid ObjectID, update by "name" or "hobby" instead
		queryFilter = bson.M{
			"$or": []bson.M{
				{"name": userId},
				{"hobby": userId},
			},
		}
	} else {
		// If it's a valid ObjectID, update by "_id"
		queryFilter = bson.M{"_id": objectID}
	}

	result, err = coll.UpdateOne(c.Context(), queryFilter, bson.M{"$set": body})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to update person",
			"message": err.Error(),
		})
	}

	// return the person
	return c.Status(200).JSON(fiber.Map{
		"result": result,
	})

}

func DeletePerson(c *fiber.Ctx) error {

	// get the name
	userId := c.Params("userId")
	fmt.Printf("Received userId: %s\n", userId)
	if userId == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "please provide either name or id",
		})
	}
	// delete the book
	coll := database.GetDBCollection("persons")
	var result *mongo.DeleteResult
	var err error
	// Attempt to convert the userId to a valid id
	objectID, err := primitive.ObjectIDFromHex(userId)
	var queryFilter bson.M

	if err != nil {
		// If it's not a valid ObjectID, delete by "name" or "hobby" instead
		queryFilter = bson.M{
			"$or": []bson.M{
				{"name": userId},
				{"hobby": userId},
			},
		}
	} else {
		// If it's a valid ObjectID, delete by "_id"
		queryFilter = bson.M{"_id": objectID}
	}

	result, err = coll.DeleteOne(c.Context(), queryFilter)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to delete person",
			"message": err.Error(),
		})
	}

	// return the result
	return c.Status(200).JSON(fiber.Map{
		"result": result,
	})
}
