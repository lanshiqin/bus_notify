package info

import (
	. "bus_notify/src/bus"
	"bus_notify/src/db"
	"strconv"
	"testing"
)

func TestGetRoute(t *testing.T) {

	dbUrl := "root:BvcxzgfD1.@/bus_notify?charset=utf8&parseTime=True&loc=Local"
	mysql := db.MySQL{Url: dbUrl}
	mysql.InitMySQL()

	if !mysql.DB.HasTable(&StationInfo{}) {
		mysql.DB.CreateTable(&StationInfo{})
	}
	if !mysql.DB.HasTable(&StationData{}) {
		mysql.DB.CreateTable(&StationData{})
	}

	for i := 0; i < 1000; i++ {
		stationInfo := GetRoute(strconv.Itoa(i))
		if stationInfo.Data != nil {
			// 保存路线数据
			mysql.DB.Create(&stationInfo)
			for j := 0; j < len(stationInfo.Data); j++ {
				// 保存站点数据
				stationInfo.Data[j].StationId = i
				mysql.DB.Create(&stationInfo.Data[j])
			}
		}
	}

}
