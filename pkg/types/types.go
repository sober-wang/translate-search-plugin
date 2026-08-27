package types

import "encoding/json"

type ReqBody struct {
	Ver   string `json:"ver"`
	MID   string `json:"mID"`
	Count string `json:"cont"`
}

type RespBody struct {
	Ver  string `json:"ver"`
	MID  string `json:"mID"`
	Cont []Cont `json:"cont"`
}
type Cont struct {
	Group string `json:"group"`
	Items []Item `json:"items"`
}
type Item struct {
	Item string `json:"item"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Type string `json:"type"`
}

// BuildResp 构建返回体
func (rpb *RespBody) BuildResp(res string) string {
	// 这里你应该实现具体的搜索逻辑，并返回符合 dde-grand-search 格式的结果
	// 结果格式可以参考其源码，通常包含 "name", "icon", "path", "type" 等字段
	items := []Item{
		{
			Item: "Translate",
			Name: res,
			Icon: "UosAiAssistant",
			Type: "Translate/result",
		},
	}
	cont := Cont{
		Group: "Translate",
		Items: items,
	}

	rpb.Cont = []Cont{cont}
	rtres, _ := json.Marshal(rpb)
	return string(rtres)
}

// BuildEmptyResp 不合法的请求构建空的返回体
func (rpb *RespBody) BuildEmptyResp() string {
	rpb.Cont = []Cont{}
	rtres, _ := json.Marshal(rpb)
	return string(rtres)
}
