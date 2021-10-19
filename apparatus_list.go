package main

import (
	"math"
	"strconv"
	"sync"
)

type ApparatusList struct {
	numOfApparatus int
	list           []*Apparatus
	listMutex      sync.Mutex
}

func (al *ApparatusList) getTimeLeft(now int64) int {
	minWait := math.MaxInt32
	for i, _ := range al.list {
		timeLeft := al.list[i].getTimeLeft(now)
		if timeLeft == 0 {
			return 0
		}
		if minWait > timeLeft {
			minWait = timeLeft
		}
	}
	return minWait
}

func newApparatus(numOfApparatus int) *ApparatusList {
	ret := new(ApparatusList)
	ret.numOfApparatus = numOfApparatus
	for i := 0; i < numOfApparatus; i++ {
		ret.list = append(ret.list, new(Apparatus))
	}
	return ret
}

func (al *ApparatusList) getApparatusAndWait(now int64) (*Apparatus, int) {


	al.listMutex.Lock()
	appa := al.list[0].getValues()
	minWait := math.MaxInt32

	//Get the first oven to finish
	for _, loopAppa := range al.list {
		timeLeft := loopAppa.getTimeLeft(now)
		if timeLeft == 0 {
			minWait = 0
			appa = loopAppa.getValues()
			break
		}
		if minWait > timeLeft {
			minWait = timeLeft
			appa = loopAppa.getValues()
		}
	}
	al.listMutex.Unlock()

	return appa, minWait
}

func (al *ApparatusList) getStatus() string {
	ret := ""
	for i, apparatus := range al.list {
		identification := "Id:" + strconv.Itoa(i)
		if apparatus.busy == 1 {
			identification += " Used by cook id:"
			if apparatus.cook != nil {
				identification += strconv.Itoa(apparatus.cook.id)
			}
		} else {
			identification += " Free"
		}
		ret += makeDiv(identification)
	}
	return ret
}
