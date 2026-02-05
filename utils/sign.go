package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/samber/lo"
	"github.com/spf13/cast"
)

func Sign(params map[string]string, key string) (string, error) {
	// 1. 依照 ASCII 顺序由小到大做排序
	//  key1=value1&key2=value2...的方式组出字串，最后再加上&secret_key={密钥}
	keys := lo.Keys(params)
	sort.Strings(keys)

	var sb strings.Builder
	for _, k := range keys {
		value := cast.ToString(params[k])
		if k != "sign" && value != "" { // && value != ""
			//只有非空才可以参与签名
			sb.WriteString(fmt.Sprintf("%s=%s&", k, url.QueryEscape(value)))
		}
	}
	signStr := sb.String()
	signStr += fmt.Sprintf("secret_key=%s", key)
	signStr, err := url.QueryUnescape(signStr)
	if err != nil {
		fmt.Println("QueryUnescape error:", err)
		return "", err
	}

	fmt.Printf("[rawString]%s\n", signStr)

	// 第2步骤产生签名字串做 md5 加签得到sign
	hash := md5.Sum([]byte(signStr))
	signResult := hex.EncodeToString(hash[:])

	fmt.Printf("MD5签名str: %s\n\n", signResult)
	return signResult, nil
}

func Verify(params map[string]string, signKey string) (bool, error) {
	// Check if signature exists in params
	signature, exists := params["sign"]
	if !exists {
		return false, nil
	}

	// Remove signature from params for verification
	delete(params, "sign")

	// Generate current signature
	currentSignature, err := Sign(params, signKey)
	if err != nil {
		return false, fmt.Errorf("signature generation failed: %w", err)
	}

	// Compare signatures
	return signature == currentSignature, nil
}

// 入金&出金回调-成功-验签
func VerifyCallback(params map[string]interface{}, signKey string) bool {
	// 1. 依照 ASCII 顺序由小到大做排序
	//  key1=value1&key2=value2...的方式组出字串，最后再加上&secret_key={密钥}
	keys := lo.Keys(params)
	sort.Strings(keys)

	var sb strings.Builder
	for _, k := range keys {
		value := cast.ToString(params[k])
		if k != "sign" {
			//只有非空才可以参与签名
			sb.WriteString(fmt.Sprintf("%s=%s&", k, url.QueryEscape(value)))
		}
	}
	signStr := sb.String()
	signStr += fmt.Sprintf("secret_key=%s", signKey)

	fmt.Printf("[rawString]%s\n", signStr)

	// 第2步骤产生签名字串做 md5 加签得到sign
	hash := md5.Sum([]byte(signStr))
	signResult := hex.EncodeToString(hash[:])

	fmt.Printf("MD5签名: %s\n\n", signResult)
	fmt.Printf("回调sign值: %s\n\n", params["sign"])
	return signResult == params["sign"]
}
