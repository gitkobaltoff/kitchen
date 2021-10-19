package main

import (
	"sync"
	"time"
)

type Meal struct {
	prepared      int32
	busy          int32
	timeRequired  int
	complexity    int
	apparatus     int //0 nil, 1 oven, 2 stove
	preparingTime int64
	foodId        int
	cookId        int
	parent        *Order
	valueMutex    sync.Mutex
}

func (m *Meal) getTimeLeft(now int64) int {
	//now := time.Now().Unix()
	if m.busy == 1 {
		elapsed := int(now - m.preparingTime)
		return m.timeRequired - elapsed
	}
	elapsed := int(now - m.parent.pickUpTime)
	limit := m.parent.maxWait
	priority := m.parent.priority
	return limit - elapsed - m.timeRequired - priority
}

func (m *Meal) get() *Meal {
	m.valueMutex.Lock()
	defer m.valueMutex.Unlock()
	return m
}

func (m *Meal) set(meal *Meal){
	m.valueMutex.Lock()
	defer m.valueMutex.Unlock()
	m.parent = meal.parent
	m.busy = meal.busy
	m.prepared = meal.prepared
}

func (m *Meal) getBusyMeal() *Meal {
	m.busy = 1
	return m
}

func (m *Meal) prepare(cook *Cook, now int64) {
	writeMeal := m.get()
	if writeMeal.prepared == 1 {
		return
	}

	writeMeal.busy = 1
	writeMeal.preparingTime = now
	writeMeal.cookId = cook.id
	m.set(writeMeal)
	time.Sleep(time.Duration(m.timeRequired) * time.Second)
	writeMeal.busy = 1
	writeMeal.prepared = 1
	writeMeal.parent.mealCounter -= 1
	m.set(writeMeal)
}
func newMeal(parent *Order, id int) *Meal {
	food := menu[id]
	return &Meal{
		prepared:      0,
		busy:          0,
		timeRequired:  food.preparationTime,
		complexity:    food.complexity,
		apparatus:     apparatusToId[food.cookingApparatus],
		preparingTime: 0,
		foodId:        id,
		cookId:        -1,
		parent:        parent,
		valueMutex:    sync.Mutex{},
	}
}
