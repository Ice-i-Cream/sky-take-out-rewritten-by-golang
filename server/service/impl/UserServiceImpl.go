package impl

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
	"net/url"
	"sky-take-out/common/constant"
	"sky-take-out/pojo/dto"
	"sky-take-out/pojo/entity"
	"sky-take-out/resources/commonParams"
	"sky-take-out/resources/mapperParams"
	"time"
)

type UserServiceImpl struct {
}

func (u *UserServiceImpl) WxLogin(dto dto.UserLoginDTO) (user entity.User, err error) {
	openid, _ := getOpenid(dto.Code)
	if openid == "" {
		return user, fmt.Errorf(constant.LOGIN_FAILED)
	}
	user, err = mapperParams.UserMapper.GetByOpenid(openid)
	if err != nil {
		user = entity.User{
			OpenID:     openid,
			CreateTime: time.Now(),
		}
		commonParams.Tx, err = commonParams.Db.Begin()
		err = mapperParams.UserMapper.Insert(user)
		err = commonParams.Tx.Commit()
	}
	return user, err
}

// 模拟 HttpClientUtil.doGet 方法
func doGet(urlStr string, params map[string]string) (string, error) {
	// 构建 URL 参数
	values := url.Values{}
	for key, value := range params {
		values.Add(key, value)
	}
	fullURL := fmt.Sprintf("%s?%s", urlStr, values.Encode())

	// 发送 GET 请求
	resp, err := http.Get(fullURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// 获取微信用户的 openid
func getOpenid(code string) (string, error) {
	// 准备请求参数
	params := map[string]string{
		"appid":      commonParams.WechatAppid,
		"secret":     commonParams.WechatSecret,
		"js_code":    code,
		"grant_type": "authorization_code",
	}

	// 发送 GET 请求
	jsonStr, err := doGet(commonParams.WX_LOGIN, params)
	if err != nil {
		return "", err
	}

	// 解析 JSON 响应
	var jsonObject map[string]interface{}
	err = jsoniter.Unmarshal([]byte(jsonStr), &jsonObject)
	if err != nil {
		return "", err
	}

	// 获取 openid
	openid, ok := jsonObject["openid"].(string)
	if !ok {
		return "", fmt.Errorf("未找到 openid 字段或类型不匹配")
	}

	return openid, nil
}
