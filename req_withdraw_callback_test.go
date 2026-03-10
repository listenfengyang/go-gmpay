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

// {\"status\":\"completed\",\"currency\":\"SGD\",\"amount\":\"50.00\",\"hash\":\"5eb21b6da9ba06c26145c768923d611badb7ad8af26eb29cdd9aa12d30973b14\",\"ref_no\":\"2023082415202\",\"platform_charge\":\"1.00\",\"transaction_reference\":\"AB7DF280DF50B8FB4F6262AE6ECCC605\",\"updated_at\":\"2026-03-10 17:32:32\",\"is_sandbox\":1}
func GenWdCallbackRequestDemo() GmPayCallbackReq {
	return GmPayCallbackReq{
		Status:               "completed",
		Currency:             "SGD",
		Amount:               "50.00",
		Hash:                 "5eb21b6da9ba06c26145c768923d611badb7ad8af26eb29cdd9aa12d30973b14",
		RefNo:                "2023082415202",
		PlatformCharge:       "1.00",
		TransactionReference: "AB7DF280DF50B8FB4F6262AE6ECCC605",
		UpdatedAt:            "2026-03-10 17:32:32",
		IsSandBox:            "1",
	}
}
