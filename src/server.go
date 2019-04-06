package main

import (
	"bus_notify/src/bus"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	resp, _ := json.Marshal(&ApiResponse{Code: 2000, Message: "Bus Notify", Data: "copyright 蓝士钦"})
	_, _ = w.Write([]byte(resp))
}

// 检查入参是否为空
func CheckParameterIsNil(parameter ...[]string) bool {
	for i := 0; i < len(parameter); i++ {
		if parameter[i] == nil {
			return true
		}
	}
	return false
}

type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 查询站点数据处理
func stationInfoHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	cityName := r.Form["cityName"]
	routeName := r.Form["routeName"]
	direction := r.Form["direction"]
	if CheckParameterIsNil(cityName, routeName, direction) {
		resp, _ := json.Marshal(&ApiResponse{Code: 5000, Message: "请输入合法参数"})
		_, _ = w.Write([]byte(resp))
		return
	}
	url := "https://wx.shenghuoquan.cn/WxBusServer/ApiData.do"
	fmt.Println("入参: ", cityName[0], routeName[0], direction[0])
	stationInfo := bus.GetRouteAllStationInfoToSave(url, cityName[0], routeName[0], direction[0])

	station, _ := json.Marshal(&ApiResponse{Code: 2000, Message: "公交站点", Data: stationInfo})

	_, _ = w.Write([]byte(station))

}

// 查询公交信息处理
func busInfoHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	cityName := r.Form["cityName"]
	routeName := r.Form["routeName"]
	stationName := r.Form["stationName"]
	direction := r.Form["direction"]
	if CheckParameterIsNil(cityName, routeName, direction) {
		resp, _ := json.Marshal(&ApiResponse{Code: 5000, Message: "请输入合法参数"})
		_, _ = w.Write([]byte(resp))
		return
	}

	url := "https://wx.shenghuoquan.cn/WxBusServer/ApiData.do"
	fmt.Println("入参: ", cityName[0], routeName[0], direction[0])

	var objInfo *bus.ObjInfo
	if stationName != nil {
		objInfo = bus.GetRouteBusObjByName(url, cityName[0], routeName[0], direction[0], stationName[0])
	} else {
		index := r.Form["index"]
		if CheckParameterIsNil(index) {
			resp, _ := json.Marshal(&ApiResponse{Code: 5000, Message: "请输入合法参数"})
			_, _ = w.Write([]byte(resp))
			return
		}
		objIndex, _ := strconv.Atoi(index[0])
		objInfo = bus.GetRouteBusObjByIndex(url, cityName[0], routeName[0], direction[0], objIndex)
	}

	station, _ := json.Marshal(&ApiResponse{Code: 2000, Message: "公交信息", Data: objInfo})
	_, _ = w.Write([]byte(station))

}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/station_info", stationInfoHandler)
	http.HandleFunc("/bus_info", busInfoHandler)
	fmt.Println("Bus Notify ListenAndServe http://localhost:8888")
	_ = http.ListenAndServe(":8888", nil)
}
