package utils

import (
	"unsafe"
)

const (
	uuidStr  = "23456789abcdefghijkmnpqrstuvwxyz" //32
	uuidMask = 1<<5 - 1                           //11111
)

// UUID UUID生成-16bit
// 10-times '-' 6-random string '-' 6-random string
func SUID() string {
	buf := make([]byte, 16)
	idx := 15
	// 15-begin 32*32*32*32*32*32==1073741824
	rand := fastRand() //6
	for ; idx > 9; idx-- {
		buf[idx] = uuidStr[rand&uuidMask]
		rand >>= 5
	}
	now := nowms() //10
	// 50bit
	// 10000-01-01 01:00:00 =>  253402246800000 ms
	// 111001100111011111001110111001111110001010000000 48bit
	for idx = 9; idx >= 0; idx-- {
		buf[idx] = uuidStr[now&uuidMask]
		now >>= 5
	}
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
