package bus

import (
	"testing"
)

func TestGetRouteBusInfo(t *testing.T) {
	url := "https://wx.shenghuoquan.cn/WxBusServer/ApiData.do"
	cityName := "厦门市"
	routeName := "641路"
	direction := "2"
	busInfo := GetRouteBusInfo(url, cityName, routeName, direction)

	for i := 0; i < len(busInfo.List); i++ {
		t.Log(busInfo.List[i].Index, busInfo.List[i].StationName, busInfo.List[i].StatusType)
	}
}

func TestGetRouteBusObjByIndex(t *testing.T) {
	url := "https://wx.shenghuoquan.cn/WxBusServer/ApiData.do"
	cityName := "厦门市"
	routeName := "641路"
	direction := "2"
	objIndex := 5
	ObjInfo := GetRouteBusObjByIndex(url, cityName, routeName, direction, objIndex)

	if ObjInfo != nil {
		t.Logf(ObjInfo.MinBusRemark)
	} else {
		t.Log("未发车")
	}
}

func TestGetRouteBusObjByName(t *testing.T) {
	url := "https://wx.shenghuoquan.cn/WxBusServer/ApiData.do"
	cityName := "厦门市"
	routeName := "641路"
	direction := "2"
	objName := "大唐中心站"
	ObjInfo := GetRouteBusObjByName(url, cityName, routeName, direction, objName)

	if ObjInfo != nil {
		t.Logf(ObjInfo.MinBusRemark)
	} else {
		t.Log("未发车")
	}
}

func TestGetBusObj(t *testing.T) {
	url := "https://wx.shenghuoquan.cn/WxBusServer/ApiData.do"
	cityName := "厦门市"
	routeName := "641路"
	direction := "2"
	objIndex := 5
	busInfo := GetRouteBusInfo(url, cityName, routeName, direction)
	// 计算公交距离本索引的详细信息
	ObjInfo := GetBusObj(busInfo, objIndex)
	if ObjInfo != nil {
		t.Logf(ObjInfo.MinBusRemark)
	} else {
		t.Log("未发车")
	}

}
