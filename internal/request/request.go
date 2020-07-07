package request

import (
	"encoding/json"
)

type Request interface {
	JsonParams() (string, error)
}

// UnionOpenGoodsJingfenQuery 京粉精选商品查询接口
type UnionOpenGoodsJingfenQueryRequest struct {
	goodsReq *JFGoodsReq
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

type JFGoodsReq struct {
	EliteId   int32   `json:"eliteId,omitempty"`   // 频道ID
	PageIndex *int32  `json:"pageIndex,omitempty"` // 页码 默认1
	PageSize  *int32  `json:"PageSize,omitempty"`  // 每页数量，默认20，上限50
	SortName  *string `json:"sortName,omitempty"`  // 排序字段
	Sort      *string `json:"sort,omitempty"`      // asc,desc升降序,默认降序
	Pid       *string `json:"pid,omitempty"`       // 联盟id_应用id_推广位id，三段式
	Fields    *string `json:"fields,omitempty"`    // 支持出参数据筛选，逗号','分隔，目前可用：videoInfo,documentInfo
}

func NewJFGoodsReq(eliteId int32, pageIndex *int32, pageSize *int32, sortName *string, sort *string, pid *string, fields *string) *JFGoodsReq {
	return &JFGoodsReq{
		EliteId:   eliteId,
		PageIndex: pageIndex,
		PageSize:  pageSize,
		SortName:  sortName,
		Sort:      sort,
		Pid:       pid,
		Fields:    fields,
	}
}
