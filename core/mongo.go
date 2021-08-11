package core

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"
)

func init() {
	bson.SetJSONTagFallback(true)
}

type AuthDao struct {
	DataBase   string
	Collection string
	Session    *mgo.Session
}

func NewAuthDao(conf *AuthConfig) (*AuthDao, error) {
	session, err := mgo.Dial(conf.Mongo.Url)
	if err != nil {
		log.Errorf("mgo dial err %+v", err)
		return nil, err
	}
	session.SetPoolLimit(conf.Mongo.PoolSize)

	err = session.DB(conf.Mongo.DataBase).C(conf.Mongo.Collection).EnsureIndex(
		mgo.Index{
			Key:        []string{"push_agent", "bundle_id"},
			Unique:     true,
			Background: false,
			Sparse:     true,
		})
	if err != nil {
		log.Errorf("mgo ensure index err %+v", err)
		return nil, err
	}

	return &AuthDao{
		DataBase:   conf.Mongo.DataBase,
		Collection: conf.Mongo.Collection,
		Session:    session,
	}, nil
}

func (d *AuthDao) GetAuth(id UniqueId) (*AuthToken, error) {
	s := d.Session.Clone()
	defer s.Close()

	auth := &AuthToken{}
	err := s.DB(d.DataBase).C(d.Collection).Find(
		bson.M{"push_agent": id.PushAgent, "bundle_id": id.BundleId}).One(auth)
	if err == mgo.ErrNotFound {
		log.Infof("mgo not found push agent %s bundle id %s", id.PushAgent, id.BundleId)
		return nil, nil
	}
	if err != nil {
		log.Errorf("mgo find err %+v", err)
		return nil, err
	}

	return auth, nil
}

func (d *AuthDao) SetAuth(id UniqueId, auth *AuthToken) error {
	s := d.Session.Clone()
	defer s.Close()

	_, err := s.DB(d.DataBase).C(d.Collection).Upsert(
		bson.M{"push_agent": id.PushAgent, "bundle_id": id.BundleId},
		bson.M{"$set": bson.M{"token": auth.Token, "expire_at": auth.ExpireAt}},
	)
	if err != nil {
		log.Errorf("mgo upsert err %+v", err)
		return err
	}

	return nil
}

func (d *AuthDao) Close() {
	d.Session.Close()
}

type StatsDao struct {
	DataBase   string
	Collection string
	Session    *mgo.Session
}

func NewStatsDao(conf *StatsConfig) (*StatsDao, error) {
	session, err := mgo.Dial(conf.Mongo.Url)
	if err != nil {
		log.Errorf("mgo dial err %+v", err)
		return nil, err
	}
	session.SetPoolLimit(conf.Mongo.PoolSize)

	err = session.DB(conf.Mongo.DataBase).C(conf.Mongo.Collection).EnsureIndex(
		mgo.Index{
			Key:        []string{"push_agent", "bundle_id"},
			Unique:     true,
			Background: false,
			Sparse:     true,
		})
	if err != nil {
		log.Errorf("mgo ensure index err %+v", err)
		return nil, err
	}

	return &StatsDao{
		DataBase:   conf.Mongo.DataBase,
		Collection: conf.Mongo.Collection,
		Session:    session,
	}, nil
}

func (d *StatsDao) Close() {
	d.Session.Close()
}
