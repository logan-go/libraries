//对微信API服务进行请求的封装
package wxclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func SendRequest(wxUri string, params map[string]interface{}) {
	urlStr := "https://api.weixin.qq.com" + wxUri

	body := make([]string, 0, len(params))
	for k, v := range params {
		body = append(body, k+"="+fmt.Sprint(v))
	}

	r := strings.NewReader(strings.Join(body, "&"))

	req, err := http.NewRequest("GET", urlStr, r)
	if err != nil {
		fmt.Println("Request New ERROR:", err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request ERROR:", err)
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read Error:", err)
	}

	b := WxCode2Session{}
	json.Unmarshal(resBody, &b)

	fmt.Printf("%+v\n", b)
}
