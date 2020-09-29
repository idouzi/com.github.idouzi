package util

import (
	"crypto/md5"
	"encoding/hex"
	"sort"
)

//签名算法
func MD5Sign(params map[string]string, appSecret string) string {

	//检查参数
	_, ok1 := params["timestamp"]
	_, ok2 := params["nonceStr"]
	if !ok1 || !ok2 {
		return "缺少必备参数!"
	}
	var keys []string

	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var signStr string
	for _, k := range keys {
		signStr += k + "=" + params[k] + "&"
	}
	signStr += "appSecret=" + appSecret
	hash := md5.New()
	hash.Write([]byte(signStr))
	return hex.EncodeToString(hash.Sum(nil))
}
