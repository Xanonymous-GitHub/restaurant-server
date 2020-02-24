package db

import (
	"database/sql"
	"github.com/Xanonymous-GitHub/restaurant-server/secretes"
	"time"
)

type Restaurants struct {
	Name     string
	LastDate time.Time
	Detail   string
}

func (d Restaurants) GetList(token string) {
	//You can't see the sql information.
	//get secrete info.
	sqlInfo, err := secretes.GetSqlInfo(secretes.TeU, token)
	if err != nil {
		panic(err.Error())
	}
	//open the database on server.
	db, err := sql.Open("mysql", sqlInfo["UserName"]+":"+sqlInfo["Password"]+"@localhost/restaurant")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}
