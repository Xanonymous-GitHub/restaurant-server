package db

import (
	"database/sql"
	"github.com/Xanonymous-GitHub/restaurant-server/secretes"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Restaurants struct {
	Name     string
	LastDate string
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
	db, err := sql.Open("mysql", sqlInfo["UserName"]+":"+sqlInfo["Password"]+"@tcp(localhost)/restaurant")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// Execute the query
	results, err := db.Query("SELECT * FROM restaurant")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var restaurant Restaurants
		err = results.Scan(&restaurant.Name, &restaurant.LastDate, &restaurant.Detail)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(restaurant.Name)
		log.Printf(restaurant.LastDate)
		log.Printf(restaurant.Detail)
	}
}
