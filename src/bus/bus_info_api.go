package bus

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// 主要数据
type InfoData struct {
	StationName string `json:"stationName"` // 站点名称
	Index       int    `json:"index"`       // 站点在路线上的索引位置 从0开始
	Arrive      int    `json:"arrive"`      // 驾驶状态 0行驶中 1停止行驶
	Come        int    `json:"come"`        // 到站状态 0已到站 1即将到站
	ShowType    int    `json:"showType"`    // 显示类型
}

// 详细数据
type InfoList struct {
	StationName              string  `json:"stationName"`              // 站点名称
	StatusType               string  `json:"statusType"`               // 状态类别 0已到站 2即将到站
	Index                    int     `json:"index"`                    // 站点在路线上的索引位置 从0开始
	StationLat               float64 `json:"station_lat"`              // 经度
	StationLng               float64 `json:"station_lng"`              // 纬度
	BusLat                   float64 `json:"bus_lat"`                  // 公交经度
	BusLng                   float64 `json:"bus_lng"`                  // 公交纬度
	BusNumber                string  `json:"busNumber"`                // 公交车牌号码
	CrowdedStatus            string  `json:"crowdedStatus"`            // 拥挤状态
	BusToStationNiheDistance float64 `json:"busToStationNiheDistance"` // 公交到站距离
	NihePointIndex           int     `json:"nihePointIndex"`           // 附近点指数
	Angle                    float64 `json:"angle"`                    // 角度
	RouteNumber              string  `json:"routeNumber"`              // 路由号
	UpperOrDown              string  `json:"upperOrDown"`              // 正向或逆向
	BusIcon                  string  `json:"busIcon"`                  // 公交图标
	RecTime                  int     `json:"_recTime"`                 // 获取时间
}

// 公交信息
type Info struct {
	Data []InfoData `json:"data"`
	List []InfoList `json:"list"`
}

// 指定对象信息
type ObjInfo struct {
	MinBusStationName       string     // 最小距离的公交信息站点名称
	MinBusStationIndex      int        // 最小距离的公交信息站点在路线上的索引位置
	MinBusStationLat        float64    // 最小距离的公交站点经度
	MinBusStationLng        float64    // 最小距离的公交站点纬度
	MinBusLat               float64    // 最小距离的公交经度
	MinBusLng               float64    // 最小距离的公交纬度
	MinBusNumber            string     // 最小距离的公交公交车牌号码
	MinBusToCurrentDistance float64    // 最小距离的公交到站距离米
	MinBusToCurrentStation  int        // 最小距离的公交距离当前站的站点数
	MinBusRemark            string     // 最小距离的公交信息备注
	WillComeList            []InfoList // 即将到站的所有公交详细数据
}

const busInfoCmd = "104"          // 查询站点信息操作码
const IntMax = int(^uint(0) >> 1) // 无符号整型最大值

// 获取指定路线的实时公交信息
func GetRouteBusInfo(requestURL string, cityName string, routeName string, direction string) *Info {
	response, _ := http.Post(requestURL, "application/x-www-form-urlencoded", strings.NewReader(
		"CMD="+busInfoCmd+
			"&CITYNAME="+cityName+
			"&LINENAME="+routeName+
			"&DIRECTION="+direction+
			""))
	body, _ := ioutil.ReadAll(response.Body)
	var resp Info
	_ = json.Unmarshal([]byte(string(body)), &resp)
	log.Println("实时公交信息：" + string(body))
	return &resp
}

// 根据指定路线索引位置，获取即将开往本索引的公交信息
func GetRouteBusObjByIndex(requestURL string, cityName string, routeName string, direction string, objIndex int) *ObjInfo {
	// 请求接口获取路线上的公交信息
	stationInfo := GetRouteBusInfo(requestURL, cityName, routeName, direction)
	// 根据路线上的公交信息和指定索引位置，计算公交距离本索引的详细信息
	return GetBusObj(stationInfo, objIndex)

}

// 根据指定路线站点名字，获取即将开往本站点的公交信息
func GetRouteBusObjByName(requestURL string, cityName string, routeName string, direction string, objName string) *ObjInfo {
	// 请求接口获取路线上的公交信息
	stationInfo := GetRouteBusInfo(requestURL, cityName, routeName, direction)
	// 根据站点名称获取站点信息
	stationData := GetRouteStationData(requestURL, cityName, routeName, direction, objName)
	// 根据路线上的公交信息和指定索引位置，计算公交距离本索引的详细信息
	return GetBusObj(stationInfo, stationData.StationOrder)

}

// 根据路线上的公交信息和指定索引位置，计算公交距离本索引的详细信息
func GetBusObj(info *Info, stationIndex int) *ObjInfo {
	var objInfo ObjInfo
	var willComeList []InfoList
	indexTmp := 0
	objInfo.MinBusToCurrentStation = IntMax
	// 遍历所有公交信息节点，计算距离最短的公交信息
	for i := 0; i < len(info.List); i++ {
		// 从当前站点大于公交索引的位置开始
		if stationIndex >= info.List[i].Index {
			willComeList = append(willComeList, info.List[i])
			// 计算出当前站点与目标站点的索引数值
			indexTmp = stationIndex - info.List[i].Index
			// 距离目标索引站点的数值小于或等于当前站点数值
			if indexTmp <= objInfo.MinBusToCurrentStation {
				// 交换两个变量
				objInfo.MinBusToCurrentStation, indexTmp = indexTmp, objInfo.MinBusToCurrentStation
				// 公交数据
				objInfo.MinBusStationName = info.List[i].StationName
				objInfo.MinBusStationIndex = info.List[i].Index
				objInfo.MinBusStationLat = info.List[i].StationLat
				objInfo.MinBusStationLng = info.List[i].StationLng
				objInfo.MinBusLat = info.List[i].BusLat
				objInfo.MinBusLng = info.List[i].BusLng
				objInfo.MinBusNumber = info.List[i].BusNumber
				objInfo.MinBusToCurrentDistance = info.List[i].BusToStationNiheDistance
			}
			objInfo.WillComeList = willComeList
		}
	}
	if objInfo.MinBusToCurrentStation == 0 {
		objInfo.MinBusRemark = "已经到站"
	} else {
		objInfo.MinBusRemark = "距离" + strconv.Itoa(objInfo.MinBusToCurrentStation) + "站"
	}
	if len(willComeList) == 0 {
		objInfo.MinBusRemark = "暂无数据"
	}
	log.Println(objInfo.MinBusRemark)
	return &objInfo
}
