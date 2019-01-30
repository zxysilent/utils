package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	FormatDateTime = "2006-01-02 15:04:05"

	// FmtyyyyMMdd 年月日
	FmtyyyyMMdd = "20060102"
	chars       = "0123456789_abcdefghijkl-mnopqrstuvwxyz" //ABCDEFGHIJKLMNOPQRSTUVWXYZ
	charsLen    = len(chars)
	mask        = 1<<6 - 1
)

var rng = rand.NewSource(time.Now().UnixNano())

// RandStr 返回指定长度的随机字符串
func RandStr(ln int) string {
	/* chars 38个字符
	 * rng.Int63() 每次产出64bit的随机数,每次我们使用6bit(2^6=64) 可以使用10次
	 */
	buf := strings.Builder{}
	for idx, cache, remain := ln, rng.Int63(), 10; idx > 0; {
		if remain == 0 {
			cache, remain = rng.Int63(), 10
		}
		buf.WriteByte(chars[int(cache&mask)%charsLen])
		cache >>= 6
		remain--
		idx--
	}
	return buf.String()
}

// Atoi 字符串转数字 def 默认值
func Atoi(s string, def ...int) (int, error) {
	rtn, err := strconv.Atoi(s)
	if err != nil && len(def) > 0 {
		return def[0], err
	}
	return rtn, nil
}

// InOf 目标是否在数组中
func InOf(in int, arr []int) bool {
	for idx := range arr {
		if in == arr[idx] {
			return true
		}
	}
	return false
}
