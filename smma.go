package indicator

type Smma struct {
	period int
	prices []float64
	smma   float64
	smmas  []float64
	offset int
	age    uint64
}

type Price struct {
	High  float64
	Low   float64
	Close float64
}

func NewSmma(period int, offset int) *Smma {
	if period <= 0 || offset < 0 {
		panic("Invalid period or offset")
	}
	return &Smma{
		period: period,
		offset: offset,
	}
}

func (s *Smma) Update(price Price) {
	mid := (price.High + price.Low) / 2
	s.prices = append(s.prices, mid)
	s.age++

	if s.age < uint64(s.period) {
		return
	} else if s.age == uint64(s.period) {
		sum1 := 0.0
		for _, p := range s.prices {
			sum1 += p
		}
		s.smma = sum1 / float64(s.period)
	} else if s.age > uint64(s.period) {
		s.smma = (s.smma*(float64(s.period)-1) + mid) / float64(s.period)
	}

	s.smmas = append(s.smmas, s.smma)

	if len(s.prices) > s.period*4 {
		s.prices = s.prices[len(s.prices)-s.period*4:]
	}
	if len(s.smmas) > s.period*4 {
		s.smmas = s.smmas[len(s.smmas)-s.period*4:]
	}
}

func (s *Smma) GetPrice() float64 {
	if len(s.smmas) == 0 || s.offset <= 0 {
		return 0
	}

	index := len(s.smmas) - (s.offset + 1)
	if index < 0 {
		index = 0
	}

	return s.smmas[index]
}

func (s *Smma) GetPreviousPrice() float64 {
	if len(s.smmas) == 0 {
		return 0
	}
	return s.smmas[len(s.smmas)-1]
}

func (s *Smma) GetFutureSegment() []float64 {
	if len(s.smmas) == 0 || s.offset <= 0 {
		return nil
	}

	startIndex := len(s.smmas) - s.offset
	if startIndex < 0 {
		startIndex = 0
	}

	return s.smmas[startIndex:]
}

func (s *Smma) Clone() *Smma {
	clone := &Smma{
		period: s.period,
		prices: append([]float64{}, s.prices...),
		smma:   s.smma,
		smmas:  append([]float64{}, s.smmas...),
		offset: s.offset,
		age:    s.age,
	}
	return clone
}
