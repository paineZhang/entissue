package util

import "math/rand"

const randomBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-=!@#$^&*()_+[]{};':\",./<>?|\\"

func NewRandomString(l int) string {
	if l <= 0 {
		panic("长度无效")
	}

	var ret = make([]byte, l)

	for index := 0; index < l; index++ {
		ret[index] = randomBytes[rand.Intn(len(randomBytes))]
	}

	return string(ret)
}
