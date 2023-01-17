package utils

import (
	"unsafe"
)

const (
	suidStr  = "23456789abcdefghijkmnpqrstuvwxyz" //32
	suidMask = 1<<5 - 1                           //11111
	uuidStr  = "0123456789abcdef"                 //16
	uuidMask = 1<<4 - 1                           //1111
)

// UUID UUID生成-16bit
// 10-times '-' 6-random string '-' 6-random string
func SUID() string {
	buf := make([]byte, 16)
	idx := 15
	// 15-begin 32*32*32*32*32*32==1073741824
	rand := fastRand() //6
	for ; idx > 9; idx-- {
		buf[idx] = suidStr[rand&suidMask]
		rand >>= 5
	}
	now := nowms() //10
	// 50bit
	// 10000-01-01 01:00:00 =>  253402246800000 ms
	// 111001100111011111001110111001111110001010000000 48bit
	for idx = 9; idx >= 0; idx-- {
		buf[idx] = suidStr[now&suidMask]
		now >>= 5
	}
	return *(*string)(unsafe.Pointer(&buf))
}

// UUID
// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx.
func UUID() string {
	buf := make([]byte, 36)
	for idx, cache, remain := 10, fastRand(), 8; idx < 36; {
		if remain == 0 {
			cache, remain = fastRand(), 8
		}
		buf[idx] = uuidStr[cache&uuidMask]
		cache >>= 4
		remain--
		idx++
	}
	// 50bit
	// 10000-01-01 01:00:00 =>  253402246800000 ms
	// 111001100111011111001110111001111110001010000000 48bit
	now := nowms() >> 4 //10
	for idx := 9; idx >= 0; idx-- {
		buf[idx] = uuidStr[now&uuidMask]
		now >>= 4
	}
	buf[8] = '-'
	buf[13] = '-'
	buf[14] = '4'
	buf[18] = '-'
	buf[23] = '-'
	return *(*string)(unsafe.Pointer(&buf))
}

//go:linkname fastRand runtime.fastrand
func fastRand() uint32

//go:noescape
//go:linkname now time.now
func now() (sec int64, nsec int32, mono int64)

// for runtime.walltime
func nowms() int64 {
	sec, nsec, _ := now()
	return sec*1e3 + int64(nsec)/1e6
	// return time.Now().UnixMilli()
}
