package request

type EliteId int32

const (
	GoodCoupon           EliteId = 1   // 1-好券商品
	SuperHypermarket             = 2   // 2-超级大卖场
	NineDivision                 = 10  // 10-9.9专区
	HotSell                      = 22  // 22-热销爆品
	Commend                      = 23  // 23-为你推荐
	DigitalHomeAppliance         = 24  // 24-数码家电
	SuperMarket                  = 25  // 25-超市
	MotherAndBabyToys            = 26  // 26-母婴玩具
	FurnitureDaily               = 27  // 27-家具日用
	BeautyMakeup                 = 28  // 28-美妆穿搭,
	HealthCare                   = 29  // 29-医药保健
	BooksStationary              = 30  // 30-图书文具
	TodayRecommend               = 31  // 31-今日必推
	BrandHQGoods                 = 32  // 32-品牌好货
	SeckillGoods                 = 33  // 33-秒杀商品
	PinGouGoods                  = 34  // 34-拼购商品
	HighIncome                   = 40  // 40-高收益
	SelfSupportHotSell           = 41  // 41-自营热卖榜
	NewArrival                   = 109 // 109-新品首发
	SelfSupport                  = 110 // 110-自营
	FirstPurchase                = 125 // 125-首购商品
	HighCommission               = 129 // 129-高佣榜单
	VideoGoods                   = 130 // 130-视频商品
)

type JFGoodsReq struct {
	EliteId   EliteId `json:"eliteId,omitempty"`   // 频道ID
	PageIndex *int32  `json:"pageIndex,omitempty"` // 页码 默认1
	PageSize  *int32  `json:"PageSize,omitempty"`  // 每页数量，默认20，上限50
	SortName  *string `json:"sortName,omitempty"`  // 排序字段
	Sort      *string `json:"sort,omitempty"`      // asc,desc升降序,默认降序
	Pid       *string `json:"pid,omitempty"`       // 联盟id_应用id_推广位id，三段式
	Fields    *string `json:"fields,omitempty"`    // 支持出参数据筛选，逗号','分隔，目前可用：videoInfo,documentInfo
}

func NewJFGoodsReq(eliteId EliteId, pageIndex *int32, pageSize *int32, sortName *string, sort *string, pid *string, fields *string) *JFGoodsReq {
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
