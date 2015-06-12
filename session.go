package mgoutils

import (
	"io"

	"gopkg.in/mgo.v2"
)

var _ io.Closer = &Session{}

type Session struct {
	*mgo.Session
}

func (this *Session) DB(dbname string) *Database {
	return &Database{
		session:  this,
		Database: this.Session.DB(dbname),
	}
}

func (this *Session) Close() error {
	this.Session.Close()
	return nil
}
