package facade

import "errors"

// 礼品信息实体, Points 就是积分的意思
type GiftInfo struct {
	ID     int
	Name   string
	Points int
}

func NewGiftInfo(id int, name string, points int) *GiftInfo {
	return &GiftInfo{
		ID:     id,
		Name:   name,
		Points: points,
	}
}

// 积分兑换礼品请求
type GiftExchangeRequest struct {
	ID         int
	UserID     int
	GiftID     int
	CreateTime int64
}

// 礼品兑换服务
type IGiftExchangeService interface {
	// 兑换礼品，并返回物流单号
	Exchange(request *GiftExchangeRequest) (error, string)
}

// 模拟用户积分管理服务的接口: 用户积分服务
type IPointsService interface {
	GetUserPoints(uid int) (error, int)
	SaveUserPoints(uid int, points int) error
}

// 模拟库存管理服务的接口: stock 就是库存的意思
type IInventoryService interface {
	GetGift(goodsID int) *GiftInfo
	GetStock(goodsID int) (error, int)
	SaveStock(goodsID int, num int) error
}

// 模拟物流下单服务的接口: 物流下单服务
type IShippingService interface {
	CreateShippingOrder(uid int, goodsID int) (error, string)
}

type xMockGiftExchangeService struct {
}

func newMockGiftExchangeService() IGiftExchangeService {
	return &xMockGiftExchangeService{}
}

var MockGiftExchangeService = newMockGiftExchangeService()

// 模拟环境下未考虑事务提交和回滚
func (m *xMockGiftExchangeService) Exchange(request *GiftExchangeRequest) (error, string) {
	gift := MockInventoryService.GetGift(request.GiftID)
	if gift == nil {
		return errors.New("gift not found"), ""
	}

	err, points := MockPointsService.GetUserPoints(request.UserID)
	if err != nil {
		return err, ""
	}
	if points < gift.Points {
		return errors.New("insufficient user points"), ""
	}

	err, stock := MockInventoryService.GetStock(gift.ID)
	if err != nil {
		return err, ""
	}
	if stock <= 0 {
		return errors.New("insufficient gift stock"), ""
	}

	err = MockInventoryService.SaveStock(gift.ID, stock-1)
	if err != nil {
		return err, ""
	}
	err = MockPointsService.SaveUserPoints(request.UserID, points - gift.Points)
	if err != nil {
		return err, ""
	}

	err, orderNo := MockShippingService.CreateShippingOrder(request.UserID, gift.ID)
	if err != nil {
		return err, ""
	}
	return nil, orderNo
}


// 模拟实现用户积分管理服务
var MockPointsService = newMockPointsService()

type xMockPointsService struct {
	userPoints map[int]int
}

func newMockPointsService() IPointsService {
	return &xMockPointsService{
		userPoints: make(map[int]int, 16),
	}
}

func (m *xMockPointsService) GetUserPoints(uid int) (error, int) {
	n, ok := m.userPoints[uid]
	if ok {
		return nil, n
	} else {
		return errors.New("user not found"), 0
	}
}

func (m *xMockPointsService) SaveUserPoints(uid int, points int) error {
	m.userPoints[uid] = points
	return nil
}

// 模拟实现库存管理服务
var MockInventoryService = newMockInventoryService()

type xMockInventoryService struct {
	goodsStock map[int]int
}

func newMockInventoryService() IInventoryService {
	return &xMockInventoryService{
		goodsStock: make(map[int]int, 16),
	}
}

func (m *xMockInventoryService) GetGift(id int) *GiftInfo {
	return NewGiftInfo(id, "mock gift", 100)
}

func (m *xMockInventoryService) GetStock(goodsID int) (error, int) {
	n, ok := m.goodsStock[goodsID]
	if ok {
		return nil, n
	} else {
		return nil, 0
	}
}

func (m *xMockInventoryService) SaveStock(goodsID int, num int) error {
	m.goodsStock[goodsID] = num
	return nil
}

// 模拟实现物流下单服务
var MockShippingService = newMockShippingService()

type xMockShippingService struct {
}

func newMockShippingService() IShippingService {
	return &xMockShippingService{}
}

func (m *xMockShippingService) CreateShippingOrder(uid int, goodsId int) (error, string) {
	return nil, "shipping-order-666"
}
