package go_nepay

import (
	"testing"
)

func TestWithdraw(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &NePayInitParams{
		MerchantInfo:      MerchantInfo{MERCHANT_ID, ACCESS_KEY},
		DepositUrl:        DEPOSIT_URL,
		WithdrawUrl:       WITHDRAW_URL,
		NotifyUrl:         NOTIFY_URL,
		WithdrawNotifyUrl: WITHDRAW_NOTIFY_URL,
	})

	//发请求
	resp, err := cli.WithdrawReq(GenWithdrawRequestDemo())
	if err != nil {
		cli.logger.Errorf("err:%s\n", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", resp)
}

func GenWithdrawRequestDemo() NePayWithdrawReq {
	return NePayWithdrawReq{
		Amount:             "1000.00",
		OrderNumber:        "2026020518462935",
		BankCardHolderName: "张三",
		BankCardNumber:     "6217001234567890123",
		BankName:           "中国建设银行",
	}
}
