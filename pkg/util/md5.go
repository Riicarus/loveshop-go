package util

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(str string, salt string, interation int) string {
	b := []byte(str)
	s := []byte(salt)
	h := md5.New()
	h.Write(s)
	h.Write(b)

	var res []byte
	res = h.Sum(nil)

	for i := 0; i < interation; i++ {
		h.Reset()
		h.Write(res)
		res = h.Sum(nil)
	}
	
	return hex.EncodeToString(res)
}