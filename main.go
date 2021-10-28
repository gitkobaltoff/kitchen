package main

import (
	"math/rand"
	"os"
	"time"
)

var diningHallHost = "http://localhost"

const diningHallPort = ":7500"
const kitchenServerPort = ":8000"

const cookN = 3
const ovenN = 3
const stoveN = 2
const orderListMaxSize = 3

const timeUnit = 100 * time.Millisecond

var kitchen Kitchen

func main() {
	rand.Seed(69)
	if args := os.Args; len(args) > 1 {
		//Set the docker internal host
		diningHallHost = args[1]
	}
	kitchen.start()
}
