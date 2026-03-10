package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// SHA256 transaction_reference + amount + currency
func Sign(params map[string]string, secretKey string) (string, error) {
	if ref, ok := params["transaction_reference"]; ok {
		params["ref_no"] = ref
	} else if ref, ok := params["ref_no"]; ok {
		params["ref_no"] = ref
	} else {
		return "", fmt.Errorf("ref_no or transaction_reference is required")
	}
	if amount, ok := params["amount"]; ok {
		params["amount"] = amount
	} else {
		return "", fmt.Errorf("amount is required")
	}
	if currency, ok := params["currency"]; ok {
		params["currency"] = currency
	} else {
		return "", fmt.Errorf("currency is required")
	}

	signStr := fmt.Sprintf("%s%s%s", params["ref_no"], params["amount"], params["currency"])
	fmt.Printf("SHA256签名before: %s\n\n", signStr)

	// 计算 HMAC-SHA256
	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(signStr))
	sum := mac.Sum(nil)

	// 输出十六进制字符串
	callbackHash := hex.EncodeToString(sum)
	fmt.Printf("SHA256签名after: %s\n\n", callbackHash)
	return callbackHash, nil
}

func Verify(params map[string]string, secretKey string) (bool, error) {
	// Check if signature exists in params
	signature, exists := params["hash"]
	if !exists {
		return false, nil
	}

	// Remove signature from params for verification
	delete(params, "hash")

	// Generate current signature
	currentSignature, err := Sign(params, secretKey)
	if err != nil {
		return false, fmt.Errorf("signature generation failed: %w", err)
	}

	// Compare signatures
	return signature == currentSignature, nil
}
