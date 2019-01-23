package wxclient

import "testing"

func TestSendRequest(t *testing.T) {
	urlStr := "/sns/jscode2session"
	params := make(map[string]interface{})
	params["appid"] = "wxff6decd518aae0c9"
	params["secret"] = "1941b4816c0c8a6cc5e5e3e3f38fc24e"
	params["code"] = "1941b4816c0c8a6cc5e5e3e3f38f"
	params["grant_type"] = "authorization_code"

	SendRequest(urlStr, params)
}
