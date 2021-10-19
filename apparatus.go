package main

import "sync"

type Apparatus struct {
	busy         int32
	meal         *Meal
	cook         *Cook
	queueWait    int
	prepareMutex sync.Mutex
	valueMutex   sync.Mutex
}

func (a *Apparatus) getTimeLeft(now int64) int {
	a.valueMutex.Lock()
	defer a.valueMutex.Unlock()
	if a.busy == 0 {
		return 0
	}
	return a.meal.getTimeLeft(now) + a.queueWait
}
func (a *Apparatus) setValues(busy int32, cook *Cook, meal *Meal) {
	a.valueMutex.Lock()
	a.busy = busy
	a.meal = meal
	a.cook = cook
	a.valueMutex.Unlock()
}

func (a *Apparatus) addQueueWait(value int) {
	a.valueMutex.Lock()
	a.queueWait += value
	a.valueMutex.Unlock()
}

func (a *Apparatus) getValues() *Apparatus {
	a.valueMutex.Lock()
	defer a.valueMutex.Unlock()
	return a
}

func (a *Apparatus) useApparatus(cook *Cook, meal *Meal, now int64) {
	timeForCurrentMeal := meal.getTimeLeft(now)
	a.addQueueWait(timeForCurrentMeal)
	a.prepareMutex.Lock()
	a.addQueueWait(-timeForCurrentMeal)
	a.setValues(1, cook, meal)

	meal.prepare(cook, now)

	a.setValues(0, nil, nil)
	a.prepareMutex.Unlock()
}
