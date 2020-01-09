package dbmodel

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	authCollection = "mqtt_user"
	aclCollection  = "mqtt_acl"
)

type AuthDetail struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	IsAdmin  bool               `bson:"is_admin"`
	Salt     string             `bson:"salt"`
}

type AclDetail struct {
	ClientId    string   `bson:"clientid"`
	Username    string   `bson:"username"`
	PubRestrict []string `bson:"publish"`
	SubRestrict []string `bson:"subscribe"`
	PubSub      []string `bson:"pubsub"`
}
