package ms

type Message struct {
	Code         int64
	Key          string
	ArgMap       map[string]string
	DefaultValue string
}

func NewMessage(c int64, k string, dv string) Message {
	return Message{
		Code:         c,
		Key:          k,
		DefaultValue: dv,
	}
}

func (m *Message) AddArg(k string, v string) {
	m.ArgMap[k] = v
}
