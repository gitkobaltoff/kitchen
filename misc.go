package main

import "time"

func makeDiv(str string) string {
	return "<div>" + str + "</div>"
}

func getUnixTimeUnits() int64 {
	return time.Now().UnixNano() / int64(timeUnit)
}
