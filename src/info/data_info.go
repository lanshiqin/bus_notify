package info

import (
	. "bus_notify/src/bus"
	"log"
)

func GetRoute(routerName string) *StationInfo {

	url := "https://wx.shenghuoquan.cn/WxBusServer/ApiData.do"
	cityName := "厦门市"
	routeName := routerName + "路"
	direction := "2"
	stationInfo := GetRouteAllStationInfo(url, cityName, routeName, direction)

	for i := 0; i < len(stationInfo.Data); i++ {
		// 打印站点索引位置和站点名称
		log.Println(stationInfo.Data[i].StationOrder, stationInfo.Data[i].StationName)
	}

	return stationInfo
}
