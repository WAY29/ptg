package encoders

import (
	"encoding/base64"
)

type EncoderBase64 struct {
}

func (e *EncoderBase64) Encode(s string) string {
	res := base64.StdEncoding.EncodeToString([]byte(s))
	return res
}

func init() {
	AddEncoder("base64", &EncoderBase64{})
}
