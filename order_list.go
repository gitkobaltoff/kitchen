package main

import (
	"math"
	"strconv"
	"sync"
)

func removeFromArr(arr *[]*Meal, ptr *Meal) {
	index := -1
	for i, meal := range *arr {
		if meal == ptr {
			index = i
			break
		}
	}
	if index != -1 {
		*arr = append((*arr)[:index], (*arr)[index+1:]...)
	}
}

type OrderList struct {
	deliveryMutex sync.Mutex
	mealMutex     sync.Mutex
	ovenList      []*Meal
	stoveList     []*Meal
	nilList       []*Meal
	orderArr      []*Order
	maxLen        int
}

func getPriority(meal *Meal, timeLeft int) float64 {
	maxWait := meal.parent.maxWait
	orderPriority := meal.parent.priority
	numMeals := int(meal.parent.mealCounter)
	return math.Tanh(float64(numMeals+meal.complexity+orderPriority)/3)*2 - math.Tanh(1-float64(timeLeft)/float64(maxWait))
}

func NewOrderList() *OrderList {
	ret := new(OrderList)
	ret.deliveryMutex = sync.Mutex{}
	ret.mealMutex = sync.Mutex{}
	ret.ovenList = []*Meal{}
	ret.stoveList = []*Meal{}
	ret.nilList = []*Meal{}
	ret.orderArr = []*Order{}
	ret.maxLen = orderListMaxSize
	return ret
}

func (orderList *OrderList) addOrder(order *Order) bool {
	orderList.deliveryMutex.Lock()
	defer orderList.deliveryMutex.Unlock()
	if len(orderList.orderArr) >= orderList.maxLen{
		return false
	}
	orderList.orderArr = append(orderList.orderArr, order)
	for _, meal := range order.mealList {
		apparatusId := meal.apparatus
		switch apparatusId {
		case 0:
			orderList.nilList = append(orderList.nilList, meal)
		case 1:
			orderList.ovenList = append(orderList.ovenList, meal)
		case 2:
			orderList.stoveList = append(orderList.stoveList, meal)
		}
	}
	return true
}

func (orderList *OrderList) getDelivery() *Delivery {
	//Prevent threads from getting the same delivery
	orderList.deliveryMutex.Lock()
	defer orderList.deliveryMutex.Unlock()

	for i, order := range orderList.orderArr {
		if order.isReady() {
			for _, meal := range order.mealList {
				apparatusId := meal.apparatus
				switch apparatusId {
				case 0:
					removeFromArr(&orderList.nilList, meal)
				case 1:
					removeFromArr(&orderList.ovenList, meal)
				case 2:
					removeFromArr(&orderList.stoveList, meal)
				}
			}
			orderList.orderArr = append(orderList.orderArr[:i], orderList.orderArr[i+1:]...)
			return newDelivery(order)
		}
	}
	return nil
}

func (orderList *OrderList) getMeal(cook *Cook) *Meal {
	//Prevent threads from taking the same meal
	orderList.mealMutex.Lock()
	defer orderList.mealMutex.Unlock()

	now := getUnixTimeUnits()
	var priority float64 = 0
	var ret *Meal
	ovenTimeLeft := kitchen.ovens.getTimeLeft(now)
	for _, meal := range orderList.ovenList {
		readMeal := meal.get()
		if readMeal.prepared == 0 && readMeal.busy == 0 && readMeal.complexity <= cook.rank {
			localPriority := getPriority(readMeal, readMeal.getTimeLeft(now)+ovenTimeLeft)
			if priority < localPriority {
				priority = localPriority
				ret = readMeal
			}
		}
	}
	stoveTimeLeft := kitchen.stoves.getTimeLeft(now)
	for _, meal := range orderList.stoveList {
		readMeal := meal.get()
		if readMeal.prepared == 0 && readMeal.busy == 0 && readMeal.complexity <= cook.rank {
			localPriority := getPriority(readMeal, readMeal.getTimeLeft(now)+stoveTimeLeft)
			if priority < localPriority {
				priority = localPriority
				ret = readMeal
			}
		}
	}
	for _, meal := range orderList.nilList {
		readMeal := meal.get()
		if readMeal.prepared == 0 && readMeal.busy == 0 && readMeal.complexity <= cook.rank {
			localPriority := getPriority(readMeal, readMeal.getTimeLeft(now))
			if priority < localPriority {
				priority = localPriority
				ret = readMeal
			}
		}
	}

	if ret != nil {
		return ret.get()
	}

	return ret
}

func (orderList *OrderList) getStatus() string {
	var ret string

	now := getUnixTimeUnits()
	for _, order := range orderList.orderArr {

		ret += makeDiv("Order id:" + strconv.Itoa(order.id) + " Meals to prepare:" + strconv.Itoa(int(order.mealCounter)) + "/" + strconv.Itoa(len(order.items)) +
			" Time passed:" + strconv.Itoa(int(now-order.pickUpTime)) + " Max wait:" + strconv.Itoa(order.maxWait))
	}
	return ret
}
