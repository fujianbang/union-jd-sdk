package response

type Wrapper struct {
	UnionOpenGoodsJingfenQueryResponse UnionOpenGoodsJingfenQueryResponse `json:"jd_union_open_goods_jingfen_query_response"`
}

type UnionOpenGoodsJingfenQueryResponse struct {
	QueryResult JingfenQueryResult `json:"result"`
}

type JingfenQueryResult struct {
	Code       int32  `json:"code"`
	Message    string `json:"message"`
	TotalCount int64  `json:"total_count"`
}
