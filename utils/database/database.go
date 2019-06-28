package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mudiralmaham/utils/Logger"
)

var (
	databaseName = "makhzan"
	Client       *mongo.Client
	DB           *mongo.Database
	err          error
)

func DatabaseInit() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	Client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		Logger.ErrorLogger.Println("Error in connecting to mongodb:", err)
		return
	}
	Logger.GeneralLogger.Println("connected to DB")

	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		Logger.ErrorLogger.Println("Error in pinging to mongodb:", err)
		return
	}
	Logger.GeneralLogger.Println("connection to DB checked successfully")

	DB = Client.Database(databaseName)
}

func DisconnectDB() {
	err := Client.Disconnect(context.TODO())

	if err != nil {
		Logger.ErrorLogger.Println("Error in disconnecting to DB:", err)
	} else {
		Logger.GeneralLogger.Println("disconnected from DB")
	}
}
