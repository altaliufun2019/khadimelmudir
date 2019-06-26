package Database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mudiralmaham/utils/Logger"
)

var(
	Client *mongo.Client
	err    error
)


func DatabaseInit() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	Client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		Logger.ErrorLogger.Println("Error in connecting to mongodb:", err)
		return
	}
	Logger.GeneralLogger.Println("connected to Database")

	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		Logger.ErrorLogger.Println("Error in pinging to mongodb:", err)
		return
	}
	Logger.GeneralLogger.Println("connection to Database checked successfully")
}


func DisconnectDB() {
	err := Client.Disconnect(context.TODO())

	if err != nil {
		Logger.ErrorLogger.Println("Error in disconnecting to Database:", err)
	} else {
		Logger.GeneralLogger.Println("disconnected from Database")
	}
}


