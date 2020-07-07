package union_jd_sdk

import (
	"go.uber.org/zap"
	"time"
	"union-jd-sdk/internal"
	"union-jd-sdk/internal/request"
	"union-jd-sdk/internal/toolkit"
)

const SERVER_URL = "https://router.jd.com/api"

type JdClient struct {
	accessToken string
	appKey      string
	appSecret   string
}

func NewJdClient(accessToken, appKey, appSecret string) *JdClient {
	return &JdClient{accessToken: accessToken, appKey: appKey, appSecret: appSecret}
}

func (c *JdClient) Execute(req request.Request) error {
	// get business params
	jsonParams, err := req.JsonParams()
	if err != nil {
		zap.L().Error("获取JsonParams失败", zap.Error(err))
		return err
	}
	zap.L().Debug("业务参数", zap.String("JsonParams", jsonParams))

	// sign
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	signParams := map[string]string{
		"app_key":     c.appKey,
		"format":      "json",
		"method":      MethodQueryJingfenGoods,
		"param_json":  jsonParams,
		"sign_method": "md5",
		"timestamp":   timestamp,
		"v":           "1.0",
	}
	signValue := toolkit.Sign(signParams, c.appSecret)
	params := internal.Config{
		Version:     "1.0",
		Method:      MethodQueryJingfenGoods,
		AccessToken: c.accessToken,
		AppKey:      c.appKey,
		SignMethod:  "md5",
		Format:      "json",
		Timestamp:   timestamp,
		Sign:        signValue,
		ParamJson:   jsonParams,
	}
	zap.L().Debug("请求参数", zap.Any("data", params))

	// 请求JD Api服务器
	resp, err := toolkit.HttpGet(SERVER_URL, params)
	if err != nil {
		zap.L().Error("http请求失败", zap.Error(err))
		return err
	}

	zap.L().Debug("响应结果", zap.ByteString("data", resp))
	return nil
}
