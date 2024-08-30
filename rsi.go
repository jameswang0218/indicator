package indicator

import (
	"sync"
)

// Rsi 表示相对强弱指数 (RSI)
type Rsi struct {
	period      int
	m           sync.RWMutex
	currentRsi  float64
	prices      []float64 // 存储价格数据
	rsis        []float64 // 存储rsi价格数据
	avgGain     float64   // 平均增益
	avgLoss     float64   // 平均损失
	updateCount int       // 更新计数
}

// NewRsi 创建一个新的 Rsi 实例
func NewRsi(period int) *Rsi {
	return &Rsi{
		period: period,
	}
}

// Update 更新 RSI 并返回当前值
func (r *Rsi) Update(price float64) float64 {
	defer r.m.Unlock()
	r.m.Lock()

	// 添加价格到 prices 切片
	r.prices = append(r.prices, price)

	// 如果 prices 的长度大于两倍 period，则删除最旧的价格
	if len(r.prices) > 2*r.period {
		r.prices = r.prices[1:]
	}
	if len(r.rsis) > 3*r.period {
		r.rsis = r.rsis[1:]
	}

	if len(r.prices) < 2 {
		return 0 // 如果只有一个价格数据，无法计算 RSI，返回默认值
	}

	change := price - r.prices[len(r.prices)-2]

	var gain, loss float64
	if change > 0 {
		gain = change
		loss = 0
	} else {
		gain = 0
		loss = -change
	}

	// 更新平均增益和平均损失
	r.updateCount++
	if r.updateCount <= r.period {
		r.avgGain = (r.avgGain*float64(r.updateCount-1) + gain) / float64(r.updateCount)
		r.avgLoss = (r.avgLoss*float64(r.updateCount-1) + loss) / float64(r.updateCount)
	} else {
		r.avgGain = (r.avgGain*float64(r.period-1) + gain) / float64(r.period)
		r.avgLoss = (r.avgLoss*float64(r.period-1) + loss) / float64(r.period)
	}

	if r.avgLoss == 0 {
		r.currentRsi = 100
	} else {
		rs := r.avgGain / r.avgLoss
		r.currentRsi = 100 - (100 / (1 + rs))
	}
	r.rsis = append(r.rsis, r.currentRsi)
	return r.currentRsi
}

// GetCurrentRsi 返回当前的 RSI 值
func (r *Rsi) GetCurrentRsi() float64 {
	return r.currentRsi
}

// Clone 返回当前 Rsi 的克隆实例
func (r *Rsi) Clone() *Rsi {
	return &Rsi{
		period:      r.period,
		currentRsi:  r.currentRsi,
		prices:      append([]float64{}, r.prices...), // 复制切片内容，避免共享底层数组
		avgGain:     r.avgGain,
		avgLoss:     r.avgLoss,
		updateCount: r.updateCount,
	}
}

// GetRsiForIndex 获取指定索引位置的 RSI 值
func (r *Rsi) GetRsiForIndex(index int) float64 {
	if index < 0 || index >= len(r.rsis) {
		return 0.0 // 或者返回一个合适的默认值
	}
	return r.rsis[index]
}

// inRange 检查当前 RSI 是否在指定范围内
func inRange(previousRsi float64, lower, upper float64) bool {
	return previousRsi >= float64(lower) && previousRsi <= float64(upper)
}
