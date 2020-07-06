package union_jd_sdk

import (
	"encoding/json"
	"go.uber.org/zap"
	"time"
	"union-jd-sdk/config"
	"union-jd-sdk/request"
	"union-jd-sdk/toolkit"
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

func (c *JdClient) PostQueryJingfenGoods() error {
	// 请求数据
	req := &request.UnionOpenGoodsJingfenQueryRequest{
		EliteId: "1",
		//PageIndex: "",
		//PageSize:  "",
		//SortName:  "",
		//Sort:      "",
		//Pid:       "",
		//Fields:    "",
	}
	bizData, err := json.Marshal(&req)
	if err != nil {
		zap.L().Error("JSON序列化失败", zap.Error(err))
		return err
	}
	zap.L().Debug("请求数据", zap.String("data", string(bizData)))

	// 封装参数
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	goodsReq := map[string]string{
		"goodsReq": string(bizData),
	}
	paramJsonBytes, err := json.Marshal(&goodsReq)
	if err != nil {
		return err
	}
	// 签名
	signParams := map[string]string{
		"app_key":     c.appKey,
		"format":      "json",
		"method":      config.MethodQueryJingfenGoods,
		"param_json":  string(paramJsonBytes),
		"sign_method": "md5",
		"timestamp":   timestamp,
		"v":           "1.0",
	}
	signValue := toolkit.Sign(signParams, c.appSecret)
	params := request.Config{
		Version:     "1.0",
		Method:      config.MethodQueryJingfenGoods,
		AccessToken: c.accessToken,
		AppKey:      c.appKey,
		SignMethod:  "md5",
		Format:      "json",
		Timestamp:   timestamp,
		Sign:        signValue,
		ParamJson:   string(paramJsonBytes),
	}
	zap.L().Debug("请求参数", zap.Any("data", params))

	resp, err := toolkit.HttpGet(SERVER_URL, params)
	if err != nil {
		zap.L().Error("http请求失败", zap.Error(err))
		return err
	}

	zap.L().Debug("响应结果", zap.ByteString("data", resp))
	return nil
}
