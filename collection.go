package mgoutils

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Collection struct {
	*mgo.Collection
	p *MgoPool
}

func (this *Collection) FindOne(query, result interface{}) error {
	return this.Collection.Find(query).One(result)
}

func (this *Collection) FindById(id, result interface{}) error {
	return this.Collection.FindId(id).One(result)
}

func (this *Collection) FindAndModify(query, change, result interface{}) error {
	changeObj := mgo.Change{
		Update:    change,
		ReturnNew: true,
	}
	_, err := this.Collection.Find(query).Apply(changeObj, result)
	return err
}

func (this *Collection) FindOrInsert(query, doc, result interface{}) error {
	changeObj := mgo.Change{
		Update:    bson.M{"$setOnInsert": doc},
		Upsert:    true,
		ReturnNew: true,
	}

	_, err := this.Collection.Find(query).Apply(changeObj, result)
	return err
}

func (this *Collection) List(query interface{}, start, limit int, sort []string, result interface{}) error {
	return this.Collection.Find(query).Sort(sort...).Skip(start).Limit(limit).All(result)
}

func (this *Collection) FindAll(query interface{}, sort []string, result interface{}) error {
	return this.Collection.Find(query).Sort(sort...).All(result)
}

func (this *Collection) Release() {
	this.p.ReleaseCollection(this)
}
