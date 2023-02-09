package ms

import (
	"github.com/samber/lo"
	"github.com/xq-libs/go-utils/types"
)

type User struct {
	ID       string
	TenantId string
	Account  string
	Name     string
}

type Pageable struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

func (p *Pageable) GetOffset() int {
	return (p.Page - 1) * p.Size
}

type Page[T any] struct {
	Pages int64 `json:"pages"`
	Total int64 `json:"total"`
	Items []T   `json:"items,omitempty"`
}

func NewPage[T any](total int64, size int64, items []T) Page[T] {
	return Page[T]{
		Pages: (total + size - 1) / size,
		Total: total,
		Items: items,
	}
}

func PageMapping[T any, R any](p Page[T], mapping types.Function[T, R]) Page[R] {
	return Page[R]{
		Pages: p.Pages,
		Total: p.Total,
		Items: lo.Map(p.Items, func(item T, index int) R {
			return mapping(item)
		}),
	}
}
