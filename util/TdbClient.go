package util

import (
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"time"
)

//域名常量
const domain = "https://points-mall.henshihui.com"

type initTdb struct {
	appKey    string
	appSecret string
}

//初始化
func New(appKey string, appSecret string) *initTdb {
	return &initTdb{
		appKey:    appKey,
		appSecret: appSecret,
	}
}

func (tdb *initTdb) Get(params map[string]string, method string) string {
	url := tdb.GetUrl(params, method)
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

//获取接口
func (tdb *initTdb) GetUrl(params map[string]string, method string) string {

	var url = domain + method + "?"
	publicParams := tdb.getPublicParams()
	allMap := tdb.putAllMap(params, publicParams)
	md5Sign := MD5Sign(allMap, tdb.appSecret)
	allMap["sign"] = md5Sign

	var keys []string

	for k := range allMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		url += k + "=" + allMap[k] + "&"
	}
	s := url[:len(url)-1]
	return s
}

//合并map
func (tdb *initTdb) putAllMap(params1 map[string]string, params2 map[string]string) map[string]string {
	n := make(map[string]string)
	for k1, v1 := range params1 {
		for k2, v2 := range params2 {
			if k1 == k2 {
				n[k1] = v2
			} else {
				if _, ok := n[k1]; !ok {
					n[k1] = v1
				}
				if _, ok := n[k2]; !ok {
					n[k2] = v2
				}
			}
		}
	}
	return n
}

//获取公共参数
func (tdb *initTdb) getPublicParams() map[string]string {
	params := make(map[string]string)
	rmndr := RandStringBytesRmndr(25)
	timeUnix := time.Now().Unix()
	params["timestamp"] = strconv.FormatInt(timeUnix, 10)
	params["nonceStr"] = rmndr
	params["appKey"] = tdb.appKey
	params["s_ver"] = "1"
	return params
}
