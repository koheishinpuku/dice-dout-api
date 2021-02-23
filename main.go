package main

import (
	"github.com/koheishinpuku/dice-dout-api/db"
	"github.com/koheishinpuku/dice-dout-api/routers"
)

func main() {
	db.Init()

	routers.Init()
}
