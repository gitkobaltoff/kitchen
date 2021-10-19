package main

import (
	"strconv"
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
		meal := kitchen.orderList.getMeal(c)
		if meal != nil {
			now := time.Now().Unix()
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
				apparatus , waitApparatus := kitchen.ovens.getApparatusAndWait(now)
				c.timeRequired += waitApparatus
				apparatus.useApparatus(c,meal,now)
			case 2:
				c.apparatusType = 2
				apparatus , waitApparatus := kitchen.stoves.getApparatusAndWait(now)
				c.timeRequired += waitApparatus
				apparatus.useApparatus(c,meal,now)
			}
		}
		delivery := kitchen.orderList.getDelivery()
		if delivery != nil {
			c.statusId = 2
			kitchen.kitchenWeb.deliver(delivery)
		}
		if meal == nil && delivery == nil {
			//Sleep for one second when there is nothing to do
			c.statusId = 0
			time.Sleep(time.Second)
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
		ret += " time left:" +strconv.Itoa(c.timeRequired - int(time.Now().Unix() - c.timeStarted))
	}

	return ret
}
