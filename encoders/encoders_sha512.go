package encoders

import (
	"crypto/sha512"
	"encoding/hex"
)

type EncoderSha512 struct {
}

func (e *EncoderSha512) Encode(s string) string {
	ee := sha512.New()
	_, err := ee.Write([]byte(s))
	if err == nil {
		return hex.EncodeToString(ee.Sum(nil))
	}
	return s
}

func init() {
	AddEncoder("sha512", &EncoderSha512{})
}
