package bus

import (
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

func GetStationInfo(requestURL string, cityName string, routeName string, direction string) *StationInfo {
	response, _ := http.Post(requestURL, "application/x-www-form-urlencoded", strings.NewReader(
		"CMD="+StationInfoCmd+
			"&CITYNAME="+cityName+
			"&LINENAME="+routeName+
			"&DIRECTION="+direction+
			""))
	body, _ := ioutil.ReadAll(response.Body)
	var resp StationInfo
	_ = json.Unmarshal([]byte(string(body)), &resp)
	log.Println(string(body))
	return &resp
}
