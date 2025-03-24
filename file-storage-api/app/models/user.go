package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User model for MongoDB
type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Username     string             `bson:"username"`
	Password     string             `bson:"password"`
	StorageQuota int64              `bson:"storage_quota"` // in bytes
	UsedStorage  int64              `bson:"used_storage"`  // in bytes
}
