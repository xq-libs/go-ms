package ms

type Response[T any] struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data,omitempty"`
}

func NewResponse[T any](data T, err error) Response[T] {
	if err != nil {
		return Response[T]{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return Response[T]{
		Code: 0,
		Msg:  "Ok",
		Data: data,
	}
}
