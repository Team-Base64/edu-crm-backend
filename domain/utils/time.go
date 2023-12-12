package utils

import (
	"log"
	"time"
)

var loc *time.Location

func init() {
	var err error
	loc, err = time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatalln("Load location error: " + err.Error())
	}
}

func TimeToString(t time.Time) string {
	return t.In(loc).Format("15:4 02.01.2006") + " по МСК"
}
