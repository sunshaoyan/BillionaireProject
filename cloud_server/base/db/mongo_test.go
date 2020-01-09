package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hackathon/conf"
	"testing"
)

func init() {
	conf.Configure.MongodbURL = `mongodb://xdms_admin:xdmspasswd@10.10.100.21:27017/xdms?replicaSet=ib-dev`
	conf.Configure.MongodbMode = 2
	MongoConnect()
}

func Test_Mongo(t *testing.T) {
	res, err := MongoDB().Collection("test").InsertOne(context.Background(), bson.M{"hello": "world"})
	if err != nil {
		t.Errorf("insert error:%s", err.Error())
	}
	id := res.InsertedID
	t.Logf("InsertedID %s", id)
	result := struct {
		Key1 int
		Key2 int
	}{}
	//filter := bson.D{{"hello", "ddd"}}
	filter := bson.D{
		{"key1", bson.D{{"$in", bson.A{11}}}},
	}
	op := options.FindOne().SetProjection(bson.D{{"key1", 1}})
	err = MongoDB().Collection("test").FindOne(context.Background(), filter, op).Decode(&result)
	if err != nil {
		t.Errorf("find error:%s", err.Error())
	}
	t.Logf("find result [%v]", result)

}
