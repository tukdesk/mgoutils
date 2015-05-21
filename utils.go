package mgoutils

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	ErrNotFound   = mgo.ErrNotFound
	IsDup         = mgo.IsDup
	EmptyObjectId = bson.ObjectId("")
)

func NewId() bson.ObjectId {
	return bson.NewObjectId()
}

func IdFromString(s string) (bson.ObjectId, bool) {
	if bson.IsObjectIdHex(s) {
		return bson.ObjectIdHex(s), true
	}
	return EmptyObjectId, false
}

func IsNotFound(err error) bool {
	return err == mgo.ErrNotFound
}

func IsEmptyObjectId(id bson.ObjectId) bool {
	return id == EmptyObjectId
}
