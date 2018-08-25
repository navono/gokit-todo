package db

import (
	mgo "gopkg.in/mgo.v2"
)

var mgoSession *mgo.Session
var mongoConnStr = "mongodb://localhost:27017"

// GetMongoSession Creates a new session if mgoSession is nil i.e there is no active mongo session.
//If there is an active mongo session it will return a Clone
func GetMongoSession() (*mgo.Session, error) {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(mongoConnStr)
		if err != nil {
			return nil, err
		}
	}
	return mgoSession.Clone(), nil
}

// package db

// import (
// 	"context"

// 	mongo "github.com/mongodb/mongo-go-driver/mongo"
// )

// var mongoSession *mongo.Session
// var mongoClient *mongo.Client
// var mongoConnStr = "mongodb://localhost:27017"

// // GetMongoDB Creates a new MongoDB
// func GetMongoDB() (*mongo.Session, error) {
// 	if mongoSession == nil {
// 		client, err := mongo.Connect(context.Background(), mongoConnStr, nil)

// 		mongoSession, err = client.StartSession()

// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	return mongoSession, nil

// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// session, err := client.StartSession();
// 	// client.

// 	// db := client.Database("todo_app")
// 	// return db, nil
// }

// // GetMongoDBClient create a client
// func GetMongoDBClient() (*mongo.Client, error) {
// 	if mongoClient == nil {
// 		var err error
// 		mongoClient, err = mongo.Connect(context.Background(), mongoConnStr, nil)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	return mongoClient, nil
// }
