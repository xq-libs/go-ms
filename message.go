package ms

import "github.com/xq-libs/go-utils/stringutil"

type Message struct {
	Code         int64
	Key          string
	ParamMap     map[string]string
	DefaultValue string
}

func NewMessage(c int64, k string, dv string) Message {
	return Message{
		Code:         c,
		Key:          k,
		DefaultValue: dv,
	}
}

func (m *Message) AppendParamMap(pm map[string]string) Message {
	return Message{
		Code:         m.Code,
		Key:          m.Key,
		ParamMap:     pm,
		DefaultValue: m.DefaultValue,
	}
}

func (m *Message) AppendParam(k string, v string) Message {
	return Message{
		Code:         m.Code,
		Key:          m.Key,
		ParamMap:     map[string]string{k: v},
		DefaultValue: m.DefaultValue,
	}
}

func (m *Message) AppendParam2(k1 string, v1 string, k2 string, v2 string) Message {
	return Message{
		Code:         m.Code,
		Key:          m.Key,
		ParamMap:     map[string]string{k1: v1, k2: v2},
		DefaultValue: m.DefaultValue,
	}
}

func (m *Message) AppendParam3(k1 string, v1 string, k2 string, v2 string, k3 string, v3 string) Message {
	return Message{
		Code:         m.Code,
		Key:          m.Key,
		ParamMap:     map[string]string{k1: v1, k2: v2, k3: v3},
		DefaultValue: m.DefaultValue,
	}
}

func (m *Message) AppendParam4(k1 string, v1 string, k2 string, v2 string, k3 string, v3 string, k4 string, v4 string) Message {
	return Message{
		Code:         m.Code,
		Key:          m.Key,
		ParamMap:     map[string]string{k1: v1, k2: v2, k3: v3, k4: v4},
		DefaultValue: m.DefaultValue,
	}
}

func (m *Message) AppendParam5(k1 string, v1 string, k2 string, v2 string, k3 string, v3 string, k4 string, v4 string, k5 string, v5 string) Message {
	return Message{
		Code:         m.Code,
		Key:          m.Key,
		ParamMap:     map[string]string{k1: v1, k2: v2, k3: v3, k4: v4, k5: v5},
		DefaultValue: m.DefaultValue,
	}
}

func (m *Message) GetDefaultMessage() string {
	return stringutil.MustParseTemplate(m.DefaultValue, m.ParamMap)
}
