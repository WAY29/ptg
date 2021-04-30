package encoders

import (
	"crypto/sha1"
	"encoding/hex"
)

type EncoderSha1 struct {
}

func (e *EncoderSha1) Encode(s string) string {
	ee := sha1.New()
	_, err := ee.Write([]byte(s))
	if err == nil {
		return hex.EncodeToString(ee.Sum(nil))
	}
	return s
}

func init() {
	AddEncoder("sha1", &EncoderSha1{})
}
