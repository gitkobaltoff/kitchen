package main

import "math/rand"

type CookList struct {
	cookList      []*Cook
	cookIdCounter int
}

func NewCookList() *CookList {
	ret := new(CookList)
	ret.cookIdCounter = 0
	for i := 0; i < cookN; i++ {
		randomCook := cookPersonas[rand.Intn(len(cookPersonas))]
		randomCook.id = ret.cookIdCounter
		ret.cookIdCounter++
		if i == 0 {
			randomCook.rank = 3
		}
		ret.cookList = append(ret.cookList, NewCook(&randomCook))
	}
	return ret
}


func (cl CookList) start() {
	for _, cook := range cl.cookList {
		go cook.startWorking()
	}
}

var cookPersonas = []Cook{{
	rank:        1,
	proficiency: 1,
	name:        "Jimmy Cook",
	catchPhrase: "YES!",
}, {
	rank:        2,
	proficiency: 2,
	name:        "Andy",
	catchPhrase: "Why am i here?",
}, {
	rank:        1,
	proficiency: 3,
	name:        "Karen",
	catchPhrase: "Abolish the patriarchy",
}, {
	rank:        3,
	proficiency: 2,
	name:        "Vanessa",
	catchPhrase: "The cake is a lie",
}, {
	rank:        3,
	proficiency: 3,
	name:        "Gordon Ramsay",
	catchPhrase: "WHERE IS THE LAMB SAUCE?",
}}
