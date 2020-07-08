package union_jd_sdk

import (
	"encoding/json"
	"github.com/fujianbang/union-jd-sdk/internal"
	"github.com/fujianbang/union-jd-sdk/internal/biz"
	"github.com/fujianbang/union-jd-sdk/internal/biz/goods/jingfen_query/response"
	"github.com/fujianbang/union-jd-sdk/internal/toolkit"
	"go.uber.org/zap"
	"time"
)

const ServerUrl = "https://router.jd.com/api"

type JdClient struct {
	accessToken string
	appKey      string
	appSecret   string
}

func NewJdClient(accessToken, appKey, appSecret string) *JdClient {
	return &JdClient{accessToken: accessToken, appKey: appKey, appSecret: appSecret}
}

func (c *JdClient) Execute(req biz.Request) (interface{}, error) {
	// get business params
	jsonParams, err := req.JsonParams()
	if err != nil {
		zap.L().Error("获取JsonParams失败", zap.Error(err))
		return nil, err
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
	respBytes, err := toolkit.HttpGet(ServerUrl, params)
	if err != nil {
		zap.L().Error("http请求失败", zap.Error(err))
		return nil, err
	}

	var respObj map[string]interface{}
	if err := json.Unmarshal(respBytes, &respObj); err != nil {
		zap.L().Error("JSON反序列化失败", zap.Error(err))
		return nil, err
	}

	responseMessage := respObj[req.ResponseName()].(map[string]interface{})

	respCode := responseMessage["code"].(string)
	respResult := responseMessage["result"].(string)

	zap.L().Debug("响应结果", zap.String("code", respCode))

	var result response.UnionOpenGoodsJingfenQueryResponse
	if err := json.Unmarshal([]byte(respResult), &result); err != nil {
		zap.L().Error("消息反序列化失败", zap.Error(err))
		return nil, nil
	}

	zap.L().Debug("响应解析结果", zap.Int32("code", result.Code), zap.String("message", result.Message),
		zap.String("responseId", result.RequestId), zap.Any("data", result.Data))

	return result, nil
}
