package wxpay

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

func VerifySign(needVerifyM map[string]interface{}, sign string) bool {
	signCalc := CalcSign(needVerifyM, "API_KEY")

	//	fmt.Printf("计算出来的sign: %v\n", signCalc)
	//	fmt.Printf("微信异步通知sign: %v\n", sign)
	if sign == signCalc {
		//fmt.Println("签名校验通过!")
		return true
	}

	//fmt.Println("签名校验失败!")
	return false
}

func CalcSign(mReq map[string]interface{}, key string) (sign string) {
	//fmt.Println("微信支付签名计算, API KEY:", key)
	//STEP 1, 对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range mReq {
		sorted_keys = append(sorted_keys, k)
	}

	sort.Strings(sorted_keys)

	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sorted_keys {
		//fmt.Printf("k=%v, v=%v\n", k, mReq[k])
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	//STEP3, 在键值对的最后加上key=API_KEY
	if key != "" {
		signStrings = signStrings + "key=" + key
	}

	//STEP4, 进行MD5签名并且将所有字符转为大写.
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	return upperSign
}
