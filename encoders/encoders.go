package encoders

var Encoders map[string]Encoder

type Encoder interface {
	Encode(s string) string
}

func init() {
	Encoders = make(map[string]Encoder)
}

type EncoderBase struct {
}

func (e *EncoderBase) Encode(s string) string {
	return s
}

func AddEncoder(name string, e Encoder) {
	Encoders[name] = e
}

func GetEncoder(name string) Encoder {
	e, ok := Encoders[name]
	if !ok {
		return &EncoderBase{}
	}
	return e
}
