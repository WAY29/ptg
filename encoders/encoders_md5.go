package encoders

import (
	"crypto/md5"
	"encoding/hex"
)

type EncoderMd5 struct {
}

func (e *EncoderMd5) Encode(s string) string {
	ee := md5.New()
	_, err := ee.Write([]byte(s))
	if err == nil {
		return hex.EncodeToString(ee.Sum(nil))
	}
	return s
}

func init() {
	AddEncoder("md5", &EncoderMd5{})
}
