package encode

import (
	"crypto/md5"
	"encoding/hex"
)

var (
	h = md5.New()
)


func Sum(input []byte) string {
	h.Write(input)
	return hex.EncodeToString(h.Sum(nil))
}