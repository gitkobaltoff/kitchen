package main

import "time"

type Kitchen struct {
	kitchenWeb KitchenWeb
	orderList  *OrderList
	ovens      *ApparatusList
	stoves     *ApparatusList
	cookList   *CookList
	connected  bool
}

func (k *Kitchen) start() {
	k.cookList = NewCookList()
	k.orderList = NewOrderList()
	k.ovens = newApparatus(ovenN)
	k.stoves = newApparatus(stoveN)

	go k.tryConnectDiningHall()
	k.kitchenWeb.start()
}

func (k *Kitchen) tryConnectDiningHall() {
	k.connected = false
	for k.connected {
		if k.kitchenWeb.establishConnection() {
			k.connectionSuccessful()
			break
		} else {
			time.Sleep(time.Second)
		}
	}
}

func (k *Kitchen) deliver(delivery *Delivery) {
	k.kitchenWeb.deliver(delivery)
}

func (k *Kitchen) connectionSuccessful() {
	if k.connected {
		return
	}
	k.connected = true
	k.cookList.start()
}

func (k *Kitchen) getStatus() string {
	ret := "Cooks:"
	for _, cook := range k.cookList.cookList {
		ret += makeDiv(cook.getStatus())
	}
	ret += "Ovens:"
	ret += k.ovens.getStatus()
	ret += "Stoves:"
	ret += k.stoves.getStatus()
	ret += "OrderList:"
	ret += k.orderList.getStatus()

	return ret
}
