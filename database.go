package mgoutils

import (
	"gopkg.in/mgo.v2"
)

type Database struct {
	session *Session
	*mgo.Database
}

func (this *Database) C(name string) *Collection {
	return &Collection{
		session:    this.session,
		Collection: this.Database.C(name),
	}
}
