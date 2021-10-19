package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type KitchenHandler struct {
	packetsReceived   int32
	postReceived      int32
	latestOrder       PostOrder
	latestOrderString string
}

func (oh *KitchenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		{
			response := "OK"

			latestOrder := new(PostOrder)
			var requestBody = make([]byte, r.ContentLength)
			r.Body.Read(requestBody)
			json.Unmarshal(requestBody, latestOrder)
			kitchen.orderList.addOrder(parseOrder(latestOrder))

			//Respond with "OK"
			fmt.Fprint(w, response)
		}
	case http.MethodGet:
		{
			fmt.Fprintln(w, "<head><meta http-equiv=\"refresh\" content=\"1\" /></head>")
			fmt.Fprintln(w, makeDiv("Kitchen server is UP on port "+kitchenServerPort))
			if kitchen.connected {
				fmt.Fprintln(w, makeDiv("Kitchen successfully connected to diningHall on address:"+diningHallHost+diningHallPort))
			} else {
				fmt.Fprintln(w, makeDiv("Kitchen did not establish connection to diningHall on address:"+diningHallHost+diningHallPort))
				err := kitchen.kitchenWeb.connectionError
				if err != nil {
					fmt.Fprintln(w, makeDiv("Connection error: "+err.Error()))
				}
			}
			fmt.Fprintln(w, kitchen.getStatus())
		}
	case http.MethodConnect:
		{
			kitchen.connectionSuccessful()
			fmt.Fprint(w, "OK")
		}
	default:
		{
			fmt.Fprintln(w, "UNSUPPORTED METHOD")
		}
	}

}
