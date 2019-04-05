package main

import (
	"bus_notify/src/bus"
	"encoding/json"
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Bus Notify!"))
}
func stationInfoHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	cityName := r.Form["cityName"]
	routeName := r.Form["routeName"]
	direction := r.Form["direction"]
	url := "https://wx.shenghuoquan.cn/WxBusServer/ApiData.do"
	fmt.Println(cityName[0], routeName[0], direction[0])
	stationInfo := bus.GetStationInfo(url, cityName[0], routeName[0], direction[0])

	station, _ := json.Marshal(stationInfo)
	w.Header().Set("Content-type", "application/json")
	_, _ = w.Write([]byte(station))

}
func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/station_info", stationInfoHandler)
	fmt.Println("Bus Notify ListenAndServe http://localhost:8888")
	_ = http.ListenAndServe(":8888", nil)
}
