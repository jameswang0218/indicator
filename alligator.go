package indicator

type Alligator struct {
	lips  *Smma
	teeth *Smma
	jaw   *Smma
}

// NewAlligator 使用 SMMA 初始化鳄鱼线，支持自定义唇线、齿线和颚线的周期。
func NewAlligator() *Alligator {
	return &Alligator{
		lips:  NewSmma(13, 8),
		teeth: NewSmma(8, 5),
		jaw:   NewSmma(5, 3),
	}
}

// Update 更新最新的价格并更新所有三条移动平均线的值。
func (a *Alligator) Update(price float64) (lips, teeth, jaw float64) {
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
func (a *Alligator) GetFutureSegments() (jawSegment, teethSegment, lipsSegment []float64) {
	return a.jaw.GetFutureSegment(), a.teeth.GetFutureSegment(), a.lips.GetFutureSegment()
}

// Clone 创建当前鳄鱼线实例的深拷贝。
func (a *Alligator) Clone() *Alligator {
	return &Alligator{
		lips:  a.lips.Clone(),
		teeth: a.teeth.Clone(),
		jaw:   a.jaw.Clone(),
	}
}
