package indicator

import (
	"math"
	"sync"
)

type Boll struct {
	n      int
	k      float64
	prices []float64
	smaSum float64
	mid    float64
	up     float64
	low    float64
	m      sync.RWMutex
}

// NewBoll 初始化 Boll 结构体
func NewBoll(n int, k float64) *Boll {
	return &Boll{
		n:      n,
		k:      k,
		prices: make([]float64, 0, n),
	}
}

// AddPrice 添加一个新的价格数据点，并更新布林带
func (b *Boll) AddPrice(price float64) {
	defer b.m.Unlock()
	b.m.Lock()
	if len(b.prices) >= b.n {
		// 移除最早的价格并更新 smaSum
		b.smaSum -= b.prices[0]
		b.prices = b.prices[1:]
	}
	// 添加新的价格并更新 smaSum
	b.prices = append(b.prices, price)
	b.smaSum += price

	// 如果达到n个数据点，更新布林带
	if len(b.prices) == b.n {
		b.calculate()
	}
}

// calcSMA 计算简单移动平均值
func (b *Boll) calcSMA() float64 {
	return b.smaSum / float64(b.n)
}

// calcSTD 计算标准差
func (b *Boll) calcSTD(ma float64) float64 {
	var sum float64
	for _, price := range b.prices {
		sum += (price - ma) * (price - ma)
	}
	return math.Sqrt(sum / float64(b.n))
}

// calculate 计算布林带
func (b *Boll) calculate() {
	b.mid = b.calcSMA()
	std := b.calcSTD(b.mid)
	b.up = b.mid + b.k*std
	b.low = b.mid - b.k*std
}

// GetBoll 返回当前布林带的值
func (b *Boll) GetBoll() (mid, up, low float64) {
	// 如果数据不足，返回0
	if len(b.prices) < b.n {
		return 0, 0, 0
	}
	return b.mid, b.up, b.low
}

// Clone 创建当前 Boll 对象的深拷贝
func (b *Boll) Clone() *Boll {
	// 创建一个新的 Boll 对象
	clone := &Boll{
		n:      b.n,
		k:      b.k,
		prices: make([]float64, len(b.prices)),
		smaSum: b.smaSum,
		mid:    b.mid,
		up:     b.up,
		low:    b.low,
	}
	// 复制 prices 切片的数据
	copy(clone.prices, b.prices)
	return clone
}
