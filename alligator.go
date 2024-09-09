package indicator

import (
	"math"
	"sync"
)

type Alligator struct {
	lips  *Smma
	teeth *Smma
	jaw   *Smma
	m     sync.RWMutex
}

// NewAlligator 使用 SMMA 初始化鳄鱼线，支持自定义唇线、齿线和颚线的周期。
func NewAlligator() *Alligator {
	return &Alligator{
		lips:  NewSmma(5, 3),
		teeth: NewSmma(8, 5),
		jaw:   NewSmma(13, 8),
	}
}

// Update 更新最新的价格并更新所有三条移动平均线的值。
func (a *Alligator) Update(price Price) (lips, teeth, jaw float64) {
	defer a.m.Unlock()
	a.m.Lock()
	a.lips.Update(price)
	a.teeth.Update(price)
	a.jaw.Update(price)
	return a.lips.GetPrice(), a.teeth.GetPrice(), a.jaw.GetPrice()
}

// GetValues 返回当前唇线、齿线和颚线的值。
func (a *Alligator) GetValues() (lips, teeth, jaw float64) {
	return a.lips.GetPrice(), a.teeth.GetPrice(), a.jaw.GetPrice()
}

// GetPreviousValues 返回没有平滑过的，唇线、齿线和颚线的值。
func (a *Alligator) GetPreviousValues() (lips, teeth, jaw float64) {
	return a.lips.GetPreviousPrice(), a.teeth.GetPreviousPrice(), a.jaw.GetPreviousPrice()
}

// GetFutureSegments 返回所有 SMMA 实例的未来数据段
func (a *Alligator) GetFutureSegments() (lipsSegment, teethSegment, jawSegment []float64) {
	return a.lips.GetFutureSegment(), a.teeth.GetFutureSegment(), a.jaw.GetFutureSegment()
}

// Clone 创建当前鳄鱼线实例的深拷贝。
func (a *Alligator) Clone() *Alligator {
	return &Alligator{
		lips:  a.lips.Clone(),
		teeth: a.teeth.Clone(),
		jaw:   a.jaw.Clone(),
	}
}

func TruncateWithMath(num float64, precision int) float64 {
	factor := math.Pow10(precision)
	return math.Round(num*factor) / factor
}
