package indicator

// Smma 结构体，包含平滑类型和历史位数
type Smma struct {
	period    int32
	prices    []float64
	smma      float64
	smmas     []float64
	offset    int32 // 用于指定要获取的历史位数
	isInitial bool
}

// NewSmma 创建一个新的 SMMA 实例，支持周期和历史位数
func NewSmma(period int32, offset int32) *Smma {
	return &Smma{
		period:    period,
		offset:    offset,
		isInitial: true,
	}
}

// Update 更新当前价格并计算 SMMA
func (s *Smma) Update(price float64) {
	// 将新价格加入价格队列
	s.prices = append(s.prices, price)

	if int32(len(s.prices)) < s.period {
		// 还没有足够的数据来计算 SMMA
		return
	}

	if s.isInitial {
		// 第一个周期，计算 SMA 作为初始 SMMA
		sum := 0.0
		// 使用周期长度内的数据计算 SMA
		startIndex := len(s.prices) - int(s.period)
		for _, p := range s.prices[startIndex:] {
			sum += p
		}
		s.smma = sum / float64(s.period)
		s.isInitial = false
	} else {
		// 后续周期，计算 SMMA
		s.smma = (s.smma*float64(s.period-1) + price) / float64(s.period)
	}

	s.smmas = append(s.smmas, s.smma)

	// 长度处理
	if int32(len(s.prices)) > s.period*4 {
		s.prices = s.prices[len(s.prices)-int(s.period*4):] // 保持价格队列长度为周期长度的4倍
	}
	if int32(len(s.smmas)) > s.period*4 {
		s.smmas = s.smmas[len(s.smmas)-int(s.period*4):] // 保持 smmas 列表长度为周期长度的4倍
	}
}

// GetPrice 返回指定历史位数的价格值
func (s *Smma) GetPrice() float64 {
	if len(s.smmas) == 0 || s.offset <= 0 {
		return 0
	}

	index := len(s.smmas) - int(s.offset)
	if index < 0 {
		index = 0
	}

	return s.smmas[index]
}

// GetPreviousPrice 返回未平滑的价格
func (s *Smma) GetPreviousPrice() float64 {
	if len(s.smmas) == 0 {
		return 0
	}
	return s.smmas[len(s.smmas)-1]
}

// GetFutureSegment 返回未来数据段的价格值
func (s *Smma) GetFutureSegment() []float64 {
	if len(s.smmas) == 0 || s.offset < 0 {
		return nil
	}

	startIndex := len(s.smmas) - int(s.offset)
	if startIndex < 0 {
		startIndex = 0
	}

	return s.smmas[startIndex:]
}

// Clone 创建当前 SMMA 实例的深拷贝
func (s *Smma) Clone() *Smma {
	clone := &Smma{
		period:    s.period,
		prices:    append([]float64{}, s.prices...), // 深拷贝价格数据
		smma:      s.smma,
		smmas:     append([]float64{}, s.smmas...), // 深拷贝 smma 数据
		offset:    s.offset,
		isInitial: s.isInitial,
	}
	return clone
}
