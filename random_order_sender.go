package main


//
////Constants
//const maxTableId = 4
//const maxItemsPerOrder = 10
//const maxPriority = 3
//const maxMaxWait = 200
//const minMaxWait = 40
//const maxCookTime = 200
//const minCookTime = 30
//const maxFoodId = 10
//const maxCookId = 3
//const maxCookingDetails = 4
//
//func getRandomItems(nItems uint32) []uint32 {
//	var ret []uint32
//	for i := uint32(0); i < nItems; i++ {
//		ret = append(ret, uint32(rand.Int31n(maxCookingDetails)))
//	}
//	return ret
//}
//func getRandomCookingDetails(n uint32) []CookDets {
//	var ret []CookDets
//	for i := uint32(0); i < n; i++ {
//		ret = append(ret, CookDets{randU32(maxFoodId), randU32(maxCookId)})
//	}
//	return ret
//}
//
//type CookDets struct {
//	foodId uint32
//	cookId uint32
//}
//type DistributionRequest struct {
//	orderId        uint32
//	tableId        uint32
//	items          []uint32
//	priority       uint32
//	maxWait        uint32
//	pickUpTime     int64
//	cookingTime    uint32
//	cookingDetails []CookDets
//}
//
//func getRandomDistributionRequest(orderId uint32,pickUpTime int64) DistributionRequest {
//	items := getRandomItems(randU32(maxItemsPerOrder))
//	_, maxId := MinMax(items)
//	return DistributionRequest{
//		orderId:        orderId,
//		tableId:        randU32(maxTableId),
//		items:          items,
//		priority:       randU32(maxPriority),
//		maxWait:        randRangeU32(minMaxWait, maxMaxWait),
//		pickUpTime:     pickUpTime,
//		cookingTime:    randRangeU32(minCookTime, maxCookTime),
//		cookingDetails: getRandomCookingDetails(maxId),
//	}
//}
