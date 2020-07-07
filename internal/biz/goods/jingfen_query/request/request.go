package request

import "encoding/json"

// UnionOpenGoodsJingfenQuery 京粉精选商品查询接口
type UnionOpenGoodsJingfenQueryRequest struct {
	goodsReq *JFGoodsReq `json:"goods_req"`
}

func NewUnionOpenGoodsJingfenQueryRequest(goodsReq *JFGoodsReq) *UnionOpenGoodsJingfenQueryRequest {
	return &UnionOpenGoodsJingfenQueryRequest{
		goodsReq: goodsReq,
	}
}

func (req *UnionOpenGoodsJingfenQueryRequest) JsonParams() (string, error) {
	goodsReq := map[string]interface{}{
		"goodsReq": &req.goodsReq,
	}
	paramJsonBytes, err := json.Marshal(&goodsReq)
	if err != nil {
		return "", err
	}
	return string(paramJsonBytes), nil
}

func (req *UnionOpenGoodsJingfenQueryRequest) ResponseName() string {
	return "jd_union_open_goods_jingfen_query_response"
}
