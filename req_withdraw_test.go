package go_gmpay

import (
	"testing"
)

func TestWithdraw(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &GmPayInitParams{
		MerchantInfo: MerchantInfo{
			ApiKey:    API_KEY,
			SecretKey: SECRET_KEY,
		},
		DepositUrl:        DEPOSIT_URL,
		WithdrawUrl:       WITHDRAW_URL,
		WithdrawNotifyUrl: WITHDRAW_NOTIFY_URL,
		ReturnUrl:         RETURN_URL,
	})

	//发请求
	resp, err := cli.WithdrawReq(GenWithdrawRequestDemo())
	if err != nil {
		cli.logger.Errorf("err:%s\n", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", resp)
}

func GenWithdrawRequestDemo() GmPayWithdrawReq {
	return GmPayWithdrawReq{
		ApiKey:         API_KEY,
		RefNo:          "2023082415202",
		Amount:         "50",
		Currency:       "SGD",
		PlayerId:       "123456",
		BankName:       "BOC",
		BankholderName: "lis",
		BankAccount:    "3523",
	}
}
