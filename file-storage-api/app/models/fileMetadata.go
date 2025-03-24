package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// FileMetadata stores details about uploaded files
type FileMetadata struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Username   string             `bson:"username"`
	Filename   string             `bson:"filename"`
	Size       int64              `bson:"size"`
	UploadTime time.Time          `bson:"upload_time"`
}
