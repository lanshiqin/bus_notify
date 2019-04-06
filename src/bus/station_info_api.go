package bus

import (
	"bus_notify/src/db"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// 站点数据
type StationData struct {
	StationId    int     `json:"stationId"`    // 站点id
	StationName  string  `json:"stationName"`  // 站点名称
	StationLon   float64 `json:"station_lon"`  // 经度
	StationLat   float64 `json:"station_lat"`  // 纬度
	StationOrder int     `json:"stationOrder"` // 站点序号
	ShowName     string  `json:"showName"`     // 显示名称
	Come         int     `json:"come"`         // 到达状态
	Arrive       int     `json:"arrive"`       // 驾驶状态
}

// 站点信息
type StationInfo struct {
	RouteName   string        `json:"routeName"`   // 公交路线名称
	UpperOrDown string        `json:"upperOrDown"` // 正向1 逆向2
	BeginTime   string        `json:"beginTime"`   // 首班车时间
	EndTime     string        `json:"endTime"`     // 末班车时间
	PlanTime    string        `json:"planTime"`    // 计划发车时间
	Common      string        `json:"commonts"`    // 路线描述
	Data        []StationData `json:"data"`        // 站点数据
}

const StationInfoCmd = "103" // 查询站点信息操作码

// 获取路线上所有站点信息
func GetRouteAllStationInfo(requestURL string, cityName string, routeName string, direction string) *StationInfo {
	response, _ := http.Post(requestURL, "application/x-www-form-urlencoded", strings.NewReader(
		"CMD="+StationInfoCmd+
			"&CITYNAME="+cityName+
			"&LINENAME="+routeName+
			"&DIRECTION="+direction+
			""))
	body, _ := ioutil.ReadAll(response.Body)
	var resp StationInfo
	_ = json.Unmarshal([]byte(string(body)), &resp)
	log.Println("站点信息:" + string(body))
	return &resp
}

// 从redis中获取路线上所有站点信息返回,如果站点信息不存在则请求接口查询并保存
func GetRouteAllStationInfoToSave(requestURL string, cityName string, routeName string, direction string) *StationInfo {
	var stationInfo StationInfo
	// 根据站点名称先从redis中取数据
	value := db.GetString(routeName + "_" + direction)
	// 数据为空则调用接口获取数据
	if value == "" {
		stationInfo = *GetRouteAllStationInfo(requestURL, cityName, routeName, direction)
		obj, _ := json.Marshal(stationInfo)
		// 保存路线上所有站点信息
		_ = db.SetString(routeName+"_"+direction, string(obj))
		for i := 0; i < len(stationInfo.Data); i++ {
			data, _ := json.Marshal(stationInfo.Data[i])
			// 依次保存路线上每个站点的信息，用来后续根据路线和站点名称快速查找站点位置。 key：路线名_站点名_方向
			_ = db.SetString(routeName+"_"+stationInfo.Data[i].StationName+"_"+direction, string(data))
		}

	} else {
		// 将json字符串转换为结构体对象
		_ = json.Unmarshal([]byte(value), &stationInfo)
	}
	return &stationInfo
}

// 根据站点名字获取路线上指定站点的信息
func GetRouteStationData(requestURL string, cityName string, routeName string, direction string, stationName string) *StationData {
	var stationData StationData
	// 先从redis中取信息，如果取不到则调用接口获取
	value := db.GetString(routeName + "_" + stationName + "_" + direction)
	if value == "" {
		// 调用接口获取路线上所有站点信息，遍历站点名称与当前名称进行匹配
		stationInfo := GetRouteAllStationInfoToSave(requestURL, cityName, routeName, direction)
		for i := 0; i < len(stationInfo.Data); i++ {
			if stationInfo.Data[i].StationName == stationName {
				return &stationInfo.Data[i]
			}
		}

	} else {
		_ = json.Unmarshal([]byte(value), &stationData)
	}
	return &stationData
}
