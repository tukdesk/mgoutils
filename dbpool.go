package mgoutils

import (
	"fmt"
	"io"
	"time"

	"github.com/facebookgo/rpool"
	"gopkg.in/mgo.v2"
)

type MgoPool struct {
	dburl   string
	dbname  string
	session *mgo.Session

	pool *rpool.Pool
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

	pool := &rpool.Pool{
		New:           mp.newDB,
		Max:           20,
		MinIdle:       3,
		IdleTimeout:   30 * time.Minute,
		ClosePoolSize: 3,
	}

	mp.pool = pool
	return mp, nil
}

func (this *MgoPool) newDB() (io.Closer, error) {
	session := &Session{Session: this.session.Copy()}
	return session, nil
}

func (this *MgoPool) GetSession() *Session {
	v, _ := this.pool.Acquire()
	return v.(*Session)
}

func (this *MgoPool) ReleaseSession(s *Session) {
	this.pool.Release(s)
	return
}

func (this *MgoPool) GetDB() *Database {
	return this.GetSession().DB(this.dbname)
}

func (this *MgoPool) ReleaseDB(db *Database) {
	this.ReleaseSession(db.session)
	return
}

func (this *MgoPool) GetCollection(collectionName string) *Collection {
	return this.GetDB().C(collectionName)
}

func (this *MgoPool) ReleaseCollection(c *Collection) {
	this.ReleaseSession(c.session)
	return
}
