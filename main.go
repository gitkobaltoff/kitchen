package main

import (
	"fmt"
	"log"
	"net/http"
)

type KitchenHandler struct {
}

func (KitchenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received something")
	if r.Method == http.MethodPost {
		//TODO make buffer static
		var buffer = make([]byte, r.ContentLength)
		r.Body.Read(buffer)

		//parseDiningHallRequest(buffer)
		fmt.Fprintln(w, "Kitchen http request post method detected.")
		fmt.Fprintln(w, "Kitchen request detected.\nPost Method Body:\n"+string(buffer))
	} else {
		fmt.Fprintln(w, "Kitchen server is UP on port "+kitchenServerPort)
	}
}

//func parseDiningHallRequest(buffer []byte) map[string]string {
//	decoder := json.Decoder{}
//	result := decoder.Decode(string(buffer))
//
//	//TODO return a order type
//	return result
//}

const diningHallPort = ":7500"
const kitchenServerPort = ":8000"

func main() {
	var kitchenServer http.Server
	kitchenServer.Addr = kitchenServerPort
	kitchenServer.Handler = KitchenHandler{}

	fmt.Println("Kitchen is listening and serving on port:"+kitchenServerPort)
	if err := kitchenServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
