package go_gmpay

import "testing"

func TestWithdrawCallback(t *testing.T) {
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

	err := cli.WithdrawCallback(GenWdCallbackRequestDemo(), func(GmPayCallbackReq) error { return nil })
	if err != nil {
		cli.logger.Errorf("Error:%s", err.Error())
		return
	}
}

//	{
//		"status": "completed",
//		"currency": "SGD",
//		"amount": "58.00",
//		"hash": "0e27d90089ee478ade943ba11134b42f46c9f792feb053049f0d247fe83c0921",
//		"ref_no": "202603170543090634",
//		"platform_charge": "1.16",
//		"transaction_reference": "44893BA0F3C3C9BC8F926DD9C9C8F7B6",
//		"updated_at": "2026-03-17 11:30:19",
//		"is_sandbox": 1
//	}
func GenWdCallbackRequestDemo() GmPayCallbackReq {
	return GmPayCallbackReq{
		Status:               "completed",
		Currency:             "SGD",
		Amount:               "58.00",
		Hash:                 "0e27d90089ee478ade943ba11134b42f46c9f792feb053049f0d247fe83c0921",
		RefNo:                "202603170543090634",
		PlatformCharge:       "1.16",
		TransactionReference: "44893BA0F3C3C9BC8F926DD9C9C8F7B6",
		UpdatedAt:            "2026-03-17 11:30:19",
		IsSandBox:            "1",
	}
}
