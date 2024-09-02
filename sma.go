package indicator

import (
	"container/list"
)

// Sma 结构体定义
type Sma struct {
	period int
	queue  *list.List
	sum    float64
}

// NewSma 创建一个新的 Sma 对象
func NewSma(period int) *Sma {
	return &Sma{
		period: period,
		queue:  list.New(),
	}
}

// Update 更新 Sma 值
func (s *Sma) Update(price float64) float64 {
	if s.queue.Len() >= s.period {
		s.sum -= s.queue.Remove(s.queue.Front()).(float64)
	}
	s.queue.PushBack(price)
	s.sum += price
	return s.sum / float64(s.queue.Len())
}

// Clone 创建并返回当前 Sma 对象的克隆副本
func (s *Sma) Clone() *Sma {
	clone := &Sma{
		period: s.period,
		queue:  list.New(),
		sum:    s.sum,
	}
	for e := s.queue.Front(); e != nil; e = e.Next() {
		clone.queue.PushBack(e.Value)
	}
	return clone
}
