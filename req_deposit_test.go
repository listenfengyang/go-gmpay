package go_nepay

import (
	"testing"
)

func TestDeposit(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &NePayInitParams{
		MerchantInfo:      MerchantInfo{MERCHANT_ID, ACCESS_KEY},
		DepositUrl:        DEPOSIT_URL,
		WithdrawUrl:       WITHDRAW_URL,
		NotifyUrl:         NOTIFY_URL,
		ReturnUrl:         RETURN_URL,
		WithdrawNotifyUrl: WITHDRAW_NOTIFY_URL,
	})
	//发请求
	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		cli.logger.Errorf("err:%s\n", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", resp)
}

func GenDepositRequestDemo() NePayDepositReq {
	return NePayDepositReq{
		Amount:      "5000.00",
		OrderNumber: "202602058352366721",
		RealName:    "张三",
		ClientIp:    "127.0.0.1",
	}
}
