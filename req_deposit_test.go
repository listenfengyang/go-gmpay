package go_gmpay

import (
	"testing"
)

func TestDeposit(t *testing.T) {
	vLog := VLog{}
	// 构造client
	cli := NewClient(vLog, &GmPayInitParams{
		MerchantInfo: MerchantInfo{
			ApiKey:    API_KEY,
			SecretKey: SECRET_KEY,
		},
		DepositUrl:       DEPOSIT_URL,
		WithdrawUrl:      WITHDRAW_URL,
		DepositNotifyUrl: DEPOSIT_CALLBACK_URL,
		ReturnUrl:        RETURN_URL,
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
		RefNo:           "20230824152007",
		Amount:          "100",
		PaymentMethodId: "2",
		Currency:        "SGD",
		PlayerId:        "123456",
	}
}
