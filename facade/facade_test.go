package facade

import (
	"testing"
	"time"
)

func Test_Facade(t *testing.T) {
	userID := 1
	giftID := 2

	// 预先存入1000积分
	err := MockPointsService.SaveUserPoints(userID, 1000)
	if err != nil {
		t.Error(err)
		return
	}

	// 预先存入1个库存
	err = MockInventoryService.SaveStock(giftID, 1)
	if err != nil {
		t.Error(err)
		return
	}

	request := &GiftExchangeRequest{
		ID: 1,
		UserID: userID,
		GiftID: giftID,
		CreateTime: time.Now().Unix(),
	}

	err, orderNo := MockGiftExchangeService.Exchange(request)
	if err != nil {
		t.Log(err)
	}

	t.Logf("shipping order no = %v\n", orderNo)
}
