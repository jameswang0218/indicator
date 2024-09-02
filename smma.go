package indicator

// Smma 结构体，包含平滑类型和历史位数
type Smma struct {
	period    int32
	prices    []float64
	smma      float64
	isInitial bool
	offset    int32 // 用于指定要获取的历史位数
}

// NewSmma 创建一个新的 SMMA 实例，支持历史位数
func NewSmma(period int32, offset int32) *Smma {
	return &Smma{
		period:    period,
		isInitial: true,
		offset:    offset,
	}
}

// Update 更新当前价格并计算 SMMA
func (s *Smma) Update(price float64) {
	// 后续周期，计算 SMMA(i)
	s.smma = (s.smma*float64(s.period-1) + price) / float64(s.period)
	// 将计算出来的 SMMA 值加入价格队列
	s.prices = append(s.prices, s.smma)
	if len(s.prices) > 100 {
		s.prices = s.prices[len(s.prices)-50:]
	}
}

// GetPrice 返回指定历史位数的价格值
func (s *Smma) GetPrice() float64 {
	if len(s.prices) == 0 || s.offset <= 0 {
		return 0
	}

	index := len(s.prices) - int(s.offset)
	if index < 0 {
		index = 0
	}

	return s.prices[index]
}

// GetPreviousPrice 返回未平滑的价格
func (s *Smma) GetPreviousPrice() float64 {
	return s.prices[len(s.prices)-1]
}

// GetFutureSegment 返回未来数据段的价格值
func (s *Smma) GetFutureSegment() []float64 {
	if len(s.prices) == 0 || s.offset < 0 {
		return nil
	}

	startIndex := len(s.prices) - int(s.offset)
	if startIndex < 0 {
		startIndex = 0
	}

	return s.prices[startIndex:]
}

// Clone 创建当前 SMMA 实例的深拷贝
func (s *Smma) Clone() *Smma {
	clone := &Smma{
		period:    s.period,
		prices:    append([]float64{}, s.prices...), // 深拷贝价格数据
		smma:      s.smma,
		isInitial: s.isInitial,
		offset:    s.offset,
	}
	return clone
}
