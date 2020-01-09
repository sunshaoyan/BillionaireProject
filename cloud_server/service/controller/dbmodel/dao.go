package dbmodel

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hackathon/base/db"
)

func InsertRecord(record *AuthDetail) (string, error) {
	db := db.MongoDB()
	collection := db.Collection(authCollection)
	res, err := collection.InsertOne(context.Background(), record)
	return res.InsertedID.(primitive.ObjectID).Hex(), err
}

func InsertRestriction(restriction *AclDetail) error {
	db := db.MongoDB()
	collection := db.Collection(aclCollection)
	_, err := collection.InsertOne(context.Background(), restriction)
	return err
}

func UpdatePassword(objectId primitive.ObjectID, password string) error {
	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"password": password}}
	_, err := db.MongoDB().Collection(authCollection).UpdateOne(context.Background(), filter, update)
	return err
}

func QueryRecord(thingId string) (*AuthDetail, error) {
	objectThingId, err := primitive.ObjectIDFromHex(thingId)
	if err != nil {
		return nil, err
	}
	result := &AuthDetail{}
	err = db.MongoDB().Collection(authCollection).FindOne(context.Background(), bson.M{"_id": objectThingId}).Decode(result)
	return result, err
}
