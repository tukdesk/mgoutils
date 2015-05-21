package mgoutils

import (
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func TestUtils(t *testing.T) {
	id := NewId()

	_, ok := IdFromString(id.Hex())
	if !ok {
		t.Error("exected to be true, got false")
	}

	_, ok2 := IdFromString("abcdefg")
	if ok2 {
		t.Error("expected to be false, got true")
	}

	emptyId := bson.ObjectId("")
	isEmpty := IsEmptyObjectId(emptyId)
	if !isEmpty {
		t.Error("expected to be true, got false")
	}
}
