package ms

type PageRequest struct {
	Pageable
}

type Response[T any] struct {
	Code int64  `json:"code"`
	Msg  string `json:"message"`
	Data T      `json:"data,omitempty"`
}
