package go_gmpay

import (
	"testing"
)

func TestDeposit(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &GmPayInitParams{
		MerchantInfo: MerchantInfo{
			ApiKey:    API_KEY,
			SecretKey: SECRET_KEY,
		},
		DepositUrl:  DEPOSIT_URL,
		WithdrawUrl: WITHDRAW_URL,
	})
	//发请求
	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		cli.logger.Errorf("err:%s\n", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", resp)
}

func GenDepositRequestDemo() GmPayDepositReq {
	return GmPayDepositReq{
		ApiKey:          API_KEY,
		RefNo:           "20230824152000000001",
		Amount:          "100",
		PaymentMethodId: "1",
		Currency:        "SGD",
		ReturnUrl:       "https://www.baidu.com",
		CallbackUrl:     "https://www.baidu.com",
		PlayerId:        "123456",
	}
}
