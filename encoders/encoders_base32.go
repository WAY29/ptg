package encoders

import (
	"encoding/base32"
)

type EncoderBase32 struct {
}

func (e *EncoderBase32) Encode(s string) string {
	res := base32.StdEncoding.EncodeToString([]byte(s))
	return res
}

func init() {
	AddEncoder("base32", &EncoderBase32{})
}
