package wxclient

//通用返回
type WxResponse struct {
	ErrorCode int    `json:"errorcode"`
	ErrMsg    string `json:"errmsg"`
}

//code2session接口返回
type WxCode2Session struct {
	WxResponse
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
}
