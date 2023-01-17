package utils

import (
	"math/rand"
	"time"
	"unsafe"
)

const (
	digitStr  = "0123456789"
	randStr   = "23456789abcdefghijkmnpqrstuvwxyz" //32
	randMask  = 1<<5 - 1                           //11111
	digitLen  = len(digitStr)
	digitMask = 1<<4 - 1
)

var rng = rand.NewSource(time.Now().UnixNano())

// RandStr 返回指定长度的随机字符串
// 包含数字 小写字母
func RandStr(ln int) string {
	/* chars 36个字符
	 * rng.Int63() 每次产出64bit的随机数,每次我们使用6bit(2^6=64) 可以使用10次
	 */
	buf := make([]byte, ln)
	for idx, cache, remain := 0, rng.Int63(), 10; idx < ln; {
		if remain == 0 {
			cache, remain = rng.Int63(), 10
		}
		buf[idx] = randStr[cache&randMask]
		cache >>= 6
		remain--
		idx++
	}
	return *(*string)(unsafe.Pointer(&buf))
}

// RandDigitStr 返回指定长度的随机字符串
// 只包含数字
func RandDigitStr(ln int) string {
	/* digits 10个字符
	 * rng.Int63() 每次产出64bit的随机数,每次我们使用4bit(2^4=16) 可以使用16次
	 */
	buf := make([]byte, ln)
	for idx, cache, remain := 0, rng.Int63(), 16; idx < ln; {
		if remain == 0 {
			cache, remain = rng.Int63(), 16
		}
		buf[idx] = digitStr[int(cache&digitMask)%digitLen]
		cache >>= 4
		remain--
		idx++
	}
	return *(*string)(unsafe.Pointer(&buf))
}
