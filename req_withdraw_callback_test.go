package go_gmpay

import (
	"testing"
)

func TestWithdrawCallback(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &NePayInitParams{
		MerchantInfo:      MerchantInfo{MERCHANT_ID, ACCESS_KEY},
		DepositUrl:        DEPOSIT_URL,
		WithdrawUrl:       WITHDRAW_URL,
		NotifyUrl:         NOTIFY_URL,
		WithdrawNotifyUrl: WITHDRAW_NOTIFY_URL,
	})

	err := cli.WithdrawCallback(GenWdRequestDemo(), func(NePayCallbackReq) error { return nil })
	if err != nil {
		cli.logger.Errorf("Error:%s", err.Error())
		return
	}
}

func GenWdRequestDemo() NePayCallbackReq {
	return NePayCallbackReq{
		Data: CallbackData{
			OrderNumber:       "2026020518462935",
			SystemOrderNumber: "GX20260205111915634496",
			UserName:          "CPT02",
			Amount:            "1000.00",
			Status:            5,
			Sign:              "62e05643cddd2fa2cd12fd46ba761073",
		},
		HttpStatusCode: 200,
		ErrorCode:      0,
		Message:        "u5f02u6b65u56deu8c03",
	}
}
