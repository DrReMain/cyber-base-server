package pagi

import "github.com/DrReMain/cyber-base-server/biz/hertz_gen/common/pagination"

type Pagi struct {
	total     int64
	more      bool
	num, size int
}

func NewPagi(total int64, more bool, num, size int) *Pagi {
	return &Pagi{total: total, more: more, num: num, size: size}
}

func (p *Pagi) Trans() *pagination.P {
	return &pagination.P{
		Total:    p.total,
		More:     p.more,
		PageNum:  int32(p.num),
		PageSize: int32(p.size),
	}
}
