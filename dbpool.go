package mgoutils

import (
	"fmt"

	"github.com/tukdesk/gopool"
	"gopkg.in/mgo.v2"
)

type MgoPool struct {
	dburl   string
	dbname  string
	session *mgo.Session

	pool *gopool.Pool
}

func NewMgoPool(dburl, dbname string) (*MgoPool, error) {
	if dburl == "" {
		return nil, fmt.Errorf("dburl required")
	}

	if dbname == "" {
		return nil, fmt.Errorf("dbname required")
	}

	s, err := mgo.Dial(dburl)
	if err != nil {
		return nil, err
	}

	mp := &MgoPool{
		dburl:   dburl,
		dbname:  dbname,
		session: s,
	}

	poolCfg := gopool.Config{
		Constructor: mp.newDB,
	}
	pool, err := gopool.NewPool(poolCfg)
	if err != nil {
		return nil, err
	}

	mp.pool = pool
	return mp, nil
}

func (this *MgoPool) newDB() (interface{}, error) {
	return this.session.Copy(), nil
}

func (this *MgoPool) GetSession() *mgo.Session {
	v, _ := this.pool.Get()
	return v.(*mgo.Session)
}

func (this *MgoPool) ReleaseSession(s *mgo.Session) {
	this.pool.Put(s)
	return
}

func (this *MgoPool) GetDB() *mgo.Database {
	return this.GetSession().DB(this.dbname)
}

func (this *MgoPool) ReleaseDB(db *mgo.Database) {
	this.pool.Put(db.Session)
	return
}

func (this *MgoPool) GetCollection(collectionName string) *Collection {
	return &Collection{
		Collection: this.GetDB().C(collectionName),
		p:          this,
	}
}

func (this *MgoPool) ReleaseCollection(c *Collection) {
	this.ReleaseDB(c.Database)
	return
}
