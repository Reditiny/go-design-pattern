package main

import "design_pattern/2_3_proxy/static"

func main() {
	// 通过火车站代理来买票
	stationProxy := new(static.ThirdParty)
	stationProxy.Sell()
}
