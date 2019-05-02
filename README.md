# 公交服务中台🚌
> 使用Golang语言编写，可以编译不同平台的可执行文件。
> 已经编译了Linux、MacOS和Windows平台的64位可执行文件，下载对应平台可执行程序可直接运行。
Releases：https://github.com/lanshiqin/bus_notify/releases

如果需要请他架构平台的可执行文件，交叉编译请参考文章 https://www.lanshiqin.com/92119e60/
# 快速开始
## 编译
```bash
go build server.go
```
## 运行
```bash
./server
```
运行成功后控制台输出
```bash
Bus Notify Server Start http://localhost:8888
```
打开浏览器访问:
```url
http://localhost:8888
```
## 请求示例
>cityName 城市名称，查询厦门市 cityName=厦门市
>routeName 公交路线名称，查询641 routeName=641路
>direction 方向（正向1 逆向2），查询逆向 direction=2
## 查询指定路线上的所有公交站点信息
http://localhost:8888/station_info?cityName=厦门市&routeName=641路&direction=2

## 查询指定路线当前即将到目标站点索引的所有公交实时信息
>index 站点所在路线上的方向索引值
http://localhost:8888/bus_info?cityName=厦门市&routeName=641路&direction=2&index=5

## 查询指定路线当前即将到目标名称站点的所有公交实时信息
>stationName 站点名称
http://localhost:8888/bus_info?cityName=厦门市&routeName=641路&direction=2&stationName=大唐中心站
```json
{
    "code": 2000,
    "message": "公交信息",
    "data": {
        "MinBusStationName": "大唐中心站",
        "MinBusStationIndex": 4,
        "MinBusStationLat": 24.478626169409846,
        "MinBusStationLng": 118.1900327454922,
        "MinBusLat": 24.47718028558475,
        "MinBusLng": 118.19298748327405,
        "MinBusNumber": "闽DZ6196",
        "MinBusToCurrentDistance": 545.5346399538611,
        "MinBusToCurrentStation": 1,
        "MinBusRemark": "距离1站",
        "WillComeList": [
            {
                "stationName": "大唐中心站",
                "statusType": "2",
                "index": 4,
                "station_lat": 24.478626169409846,
                "station_lng": 118.1900327454922,
                "bus_lat": 24.47718028558475,
                "bus_lng": 118.19298748327405,
                "busNumber": "闽DZ6196",
                "crowdedStatus": "",
                "busToStationNiheDistance": 545.5346399538611,
                "nihePointIndex": 72,
                "angle": 336.81058594658566,
                "routeNumber": "",
                "upperOrDown": "",
                "busIcon": "",
                "_recTime": 1554620145207
            }
        ]
    }
}
```

