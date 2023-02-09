package ms

type Message struct {
	Code         int64
	Key          string
	Args         []string
	DefaultValue string
}

func NewMessage(c int64, k string, dv string) Message {
	return Message{
		Code:         c,
		Key:          k,
		DefaultValue: dv,
	}
}

func (m *Message) AddArgs(args ...string) {
	m.Args = args
}
