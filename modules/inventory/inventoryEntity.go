package inventory

import "go.mongodb.org/mongo-driver/v2/bson"

type (
	Inventory struct {
		Id       bson.ObjectID `json:"_id" bson:"_id,omitempty"`
		PlayerId string        `json:"player_id" bson:"player_id"`
		ItemId   string        `json:"item_id" bson:"item_id"`
	}
)
