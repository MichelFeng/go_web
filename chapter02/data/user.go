package data

import "time"

type User struct {
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

func (u User) CreateSession() Session {

	return Session{}
}

func (s Session) Check() (bool, error) {
	return false, nil
}
