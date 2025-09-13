package utils

import (
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func ConvertToObject(id string) bson.ObjectID {
	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Error: Convert string to object id failed: %s", err.Error())
	}
	return objectId
}
