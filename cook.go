package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var cookStatus = [...]string{" Waiting.", " Preparing", " Delivering"}

type Cook struct {
	id            int
	rank          int
	proficiency   int
	name          string
	catchPhrase   string
	atWork        int32
	statusId      int
	orderId       int
	foodId        int
	apparatusType int
	timeStarted   int64
	timeRequired  int
}

func NewCook(cook *Cook) *Cook {
	ret := new(Cook)

	ret.id = cook.id
	ret.rank = cook.rank
	ret.proficiency = cook.proficiency
	ret.name = cook.name
	ret.catchPhrase = cook.catchPhrase
	ret.atWork = 0
	ret.statusId = 0
	ret.orderId = 0
	ret.foodId = 0
	ret.apparatusType = 0
	ret.timeStarted = 0
	ret.timeRequired = 0

	return ret
}

func (c *Cook) startWorking() {
	c.atWork = 1
	for c.atWork == 1 {
		didATask := false
		var wg sync.WaitGroup
		stuffToDo := rand.Intn(c.proficiency) + 1
		for i := 0; i < stuffToDo; i++ {
			wg.Add(1)
			go func() {
				meal := kitchen.orderList.getMeal(c)
				if meal != nil {
					didATask = true
					now := getUnixTimeUnits()
					c.statusId = 1
					c.orderId = meal.parent.id
					c.foodId = meal.foodId
					c.timeStarted = now
					c.timeRequired = meal.timeRequired
					switch meal.apparatus {
					case 0:
						c.apparatusType = 0
						meal.prepare(c, now)
					case 1:
						c.apparatusType = 1
						apparatus, waitApparatus := kitchen.ovens.getApparatusAndWait(now)
						c.timeRequired += waitApparatus
						apparatus.useApparatus(c, meal, now)
					case 2:
						c.apparatusType = 2
						apparatus, waitApparatus := kitchen.stoves.getApparatusAndWait(now)
						c.timeRequired += waitApparatus
						apparatus.useApparatus(c, meal, now)
					}
				}
				wg.Done()
			}()
		}
		wg.Wait()

		delivery := kitchen.orderList.getDelivery()
		if delivery != nil {
			success := false
			for success == false {
				didATask = true
				c.statusId = 2
				success = kitchen.kitchenWeb.deliver(delivery)
				if success == false {
					fmt.Println("OH NO")
				}
			}
		}
		if !didATask {
			//Sleep for one second when there is nothing to do
			c.statusId = 0
			time.Sleep(timeUnit)
		}
	}
}

func (c *Cook) stopWorking() {
	atomic.StoreInt32(&c.atWork, 0)
}

func (c *Cook) getStatus() string {
	ret := "Cook " + c.name + " id:" + strconv.Itoa(c.id) + cookStatus[c.statusId] + " "
	if c.statusId != 0 {
		ret += menu[c.foodId].name + " for order id:" + strconv.Itoa(c.orderId)
		if c.apparatusType != 0 {
			ret += " using " + idToApparatus[c.apparatusType]
		}
		ret += " time left:" + strconv.Itoa(c.timeRequired-int(getUnixTimeUnits()-c.timeStarted))
	}

	return ret
}
