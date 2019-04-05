package bus

import (
	"testing"
)

func TestGetStation(t *testing.T) {
	url := "https://wx.shenghuoquan.cn/WxBusServer/ApiData.do"
	cityName := "厦门市"
	routeName := "641路"
	direction := "1"
	stationInfo := GetStationInfo(url, cityName, routeName, direction)

	for i := 0; i < len(stationInfo.Data); i++ {
		t.Log(stationInfo.Data[i].StationName)
	}
}
