package mgoutils

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Collection struct {
	session *Session
	*mgo.Collection
	p *MgoPool
}

func (this *Collection) FindOne(query, projection map[string]interface{}, result interface{}) error {
	return this.Collection.Find(query).Select(projection).One(result)
}

func (this *Collection) FindById(id interface{}, projection map[string]interface{}, result interface{}) error {
	return this.Collection.FindId(id).Select(projection).One(result)
}

func (this *Collection) FindAndModify(query interface{}, projection map[string]interface{}, change, result interface{}) error {
	changeObj := mgo.Change{
		Update:    change,
		ReturnNew: true,
	}

	_, err := this.Collection.Find(query).Select(projection).Apply(changeObj, result)
	return err
}

func (this *Collection) FindOrInsert(query map[string]interface{}, doc, result interface{}) (bool, error) {
	changeObj := mgo.Change{
		Update:    bson.M{"$setOnInsert": doc},
		Upsert:    true,
		ReturnNew: true,
	}

	res, err := this.Collection.Find(query).Apply(changeObj, result)
	if err != nil {
		return false, err
	}

	return res.UpsertedId != nil, nil
}

func (this *Collection) List(query interface{}, projection map[string]interface{}, start, limit int, sort []string, result interface{}) error {
	return this.Collection.Find(query).Select(projection).Sort(sort...).Skip(start).Limit(limit).All(result)
}

func (this *Collection) FindAll(query interface{}, projection map[string]interface{}, sort []string, result interface{}) error {
	return this.Collection.Find(query).Select(projection).Sort(sort...).All(result)
}

func (this *Collection) Count(query interface{}) (int, error) {
	return this.Collection.Find(query).Count()
}

func (this *Collection) Release() {
	this.p.ReleaseCollection(this)
}
