package indicator

import "sync"

type ERsi struct {
	period      int32
	gainsEma    *Ema
	lossesEma   *Ema
	currentERsi float64
	prevPrice   float64
	firstPrice  bool
	m           sync.RWMutex
}

func NewERsi(weight int32) *ERsi {
	return &ERsi{
		period:     weight,
		gainsEma:   NewEma(weight),
		lossesEma:  NewEma(weight),
		firstPrice: true,
	}
}

func (this *ERsi) Update(price float64) float64 {
	defer this.m.Unlock()
	this.m.Lock()
	if this.firstPrice {
		this.prevPrice = price
		this.firstPrice = false
		return 0 // 初始状态没有RSI值
	}

	change := price - this.prevPrice
	if change > 0 {
		this.gainsEma.Update(change)
		this.lossesEma.Update(0)
	} else {
		this.gainsEma.Update(0)
		this.lossesEma.Update(-change)
	}

	if this.gainsEma.age >= uint32(this.period) {
		avgGain := this.gainsEma.GetPrice()
		avgLoss := this.lossesEma.GetPrice()
		rs := avgGain / avgLoss
		this.currentERsi = 100 - (100 / (1 + rs))
	}

	this.prevPrice = price
	return this.currentERsi
}

func (this *ERsi) GetERsi() float64 {
	return this.currentERsi
}

func (this *ERsi) Clone() *ERsi {
	return &ERsi{
		period:      this.period,
		gainsEma:    this.gainsEma.Clone(),
		lossesEma:   this.lossesEma.Clone(),
		currentERsi: this.currentERsi,
		prevPrice:   this.prevPrice,
		firstPrice:  this.firstPrice,
	}
}
