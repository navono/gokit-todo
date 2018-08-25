package io

import (
	"encoding/json"

	objectid "github.com/mongodb/mongo-go-driver/bson/objectid"
)

// Todo for service
type Todo struct {
	ID       objectid.ObjectID `json:"id" bson:"_id"`
	Title    string            `json:"title"  bson:"title"`
	Complete bool              `json:"complete" bson:"complete"`
}

func (t Todo) String() string {
	b, err := json.Marshal(t)
	if err != nil {
		return "unsupported value type"
	}
	return string(b)
}
