package main

import "github.com/Xanonymous-GitHub/restaurant-server/db"

func run() {
	var tmp db.Restaurants
	tmp.GetList("0202494c:0063")
}

func main() {
	run()
}
