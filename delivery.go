package main

import "time"

type Delivery struct {
	OrderId        int            `json:"order_id"`
	TableId        int            `json:"table_id"`
	Items          []int          `json:"items"`
	Priority       int            `json:"priority"`
	MaxWait        int            `json:"max_wait"`
	PickUpTime     int64          `json:"pick_up_time"`
	CookingTime    int            `json:"cooking_time"`
	CookingDetails []MealDelivery `json:"cooking_details"`
}

func newDelivery(order *Order) *Delivery {
	ret := new(Delivery)
	ret.OrderId = order.id
	ret.TableId = order.tableId
	ret.Items = order.items
	ret.Priority = order.priority
	ret.MaxWait = order.maxWait
	ret.PickUpTime = order.pickUpTime
	ret.CookingTime = int(time.Now().Unix() - order.pickUpTime)
	var cookingDetails []MealDelivery
	for _, meal := range order.mealList {
		cookingDetails = append(cookingDetails, MealDelivery{meal.foodId, meal.cookId})
	}
	ret.CookingDetails = cookingDetails
	return ret
}
