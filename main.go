package main

import (
	"fmt"
	"src/util"
)

func main() {

	//http.Get(url)
	//var params map[string]string
	params := make(map[string]string)
	params["redirect"] = ""
	params["unionId"] = "10001"
	params["openId"] = ""
	params["source"] = "app"

	tdb := util.New("5e84462851af1", "dee11a8e7c9a63eef48a6bf9179fa937")
	url := tdb.GetUrl(params, "/v1/user/login/auto-login")
	fmt.Println(url)
	params1 := make(map[string]string)
	params1["unionId"] = "10001"
	getUrl := tdb.Get(params1, "/sdk/api/query-user-points-amount")
	fmt.Println(getUrl)
}
