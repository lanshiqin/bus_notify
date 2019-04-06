package bus

import (
	"testing"
)

func TestGetRouteAllStationInfo(t *testing.T) {
	url := "https://wx.shenghuoquan.cn/WxBusServer/ApiData.do"
	cityName := "厦门市"
	routeName := "641路"
	direction := "2"
	stationInfo := GetRouteAllStationInfo(url, cityName, routeName, direction)

	for i := 0; i < len(stationInfo.Data); i++ {
		// 打印站点索引位置和站点名称
		t.Log(stationInfo.Data[i].StationOrder, stationInfo.Data[i].StationName)
	}
}

func TestGetRouteAllStationInfoToSave(t *testing.T) {
	url := "https://wx.shenghuoquan.cn/WxBusServer/ApiData.do"
	cityName := "厦门市"
	routeName := "641路"
	direction := "2"
	stationInfo := GetRouteAllStationInfoToSave(url, cityName, routeName, direction)

	for i := 0; i < len(stationInfo.Data); i++ {
		// 打印站点索引位置和站点名称
		t.Log(stationInfo.Data[i].StationOrder, stationInfo.Data[i].StationName)
	}
}

func TestGetRouteStationData(t *testing.T) {
	url := "https://wx.shenghuoquan.cn/WxBusServer/ApiData.do"
	cityName := "厦门市"
	routeName := "641路"
	direction := "2"
	stationName := "大唐中心站"
	stationInfo := GetRouteStationData(url, cityName, routeName, direction, stationName)
	// 打印站点索引位置和站点名称
	t.Log(stationInfo.StationOrder, stationInfo.StationName)

}
