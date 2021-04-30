package encoders

import (
	"crypto/sha256"
	"encoding/hex"
)

type EncoderSha256 struct {
}

func (e *EncoderSha256) Encode(s string) string {
	ee := sha256.New()
	_, err := ee.Write([]byte(s))
	if err == nil {
		return hex.EncodeToString(ee.Sum(nil))
	}
	return s
}

func init() {
	AddEncoder("sha256", &EncoderSha256{})
}
