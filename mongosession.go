package mongosession

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/globalsign/mgo"
)

//MgoSession and session
type MgoSession struct {
	Session *mgo.Session
}

func newMgoSession(s *mgo.Session) *MgoSession {
	return &MgoSession{s}
}

//GetMongoSession start a new mongo session
func GetMongoSession(msg string, databaseURI string) (*MgoSession, error) {
	mongoSession, err := mgo.Dial(databaseURI) // dial uri can be any mongodatabase
	defer mongoSession.Close()
	if err != nil {
		return nil, errors.New(fmt.Sprintf(`[MongoDB] %s`, err))
	}

	mongoSession.SetMode(mgo.Monotonic, true)
	mongoSession.SetSocketTimeout(1 * time.Hour) //extend for long result

	log.Printf("[MongoDB] session started! %s", msg)
	return newMgoSession(mongoSession), nil
}
