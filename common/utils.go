package common

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ZhangYet/ein"
)

// 公用功能
// 读取数据的请求
// TODO 进一步封装: 包括重试、错误处理
func GetLastQuotes() ([]*ein.LastQuote, error) {
	ret := []*ein.LastQuote{}
	resp, err := http.Get("https://api.iextrading.com/1.0/tops/last")
	if err != nil {
		return ret, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(body, &ret)
	return ret, err
}
