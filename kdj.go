package indicator

import (
	"container/list"
)

type Kdj struct {
	n1     int
	n2     int
	n3     int
	dequeH *list.List
	dequeL *list.List
	k      float64
	d      float64
}

func NewKdj(n1 int, n2 int, n3 int) *Kdj {
	return &Kdj{
		n1:     n1,
		n2:     n2,
		n3:     n3,
		dequeH: list.New(),
		dequeL: list.New(),
		k:      50,
		d:      50,
	}
}

func (k *Kdj) Update(kline Kline) (float64, float64, float64) {
	k.dequeH.PushBack(kline.High)
	if k.dequeH.Len() > k.n1 {
		k.dequeH.Remove(k.dequeH.Front())
	}

	k.dequeL.PushBack(kline.Low)
	if k.dequeL.Len() > k.n1 {
		k.dequeL.Remove(k.dequeL.Front())
	}

	if k.dequeH.Len() < k.n1 || k.dequeL.Len() < k.n1 {
		return 0, 0, 0 // not enough data
	}

	h := k.maxHigh()
	l := k.minLow()
	rsv := 100 * (kline.Close - l) / (h - l)
	if rsv != rsv { // check for NaN
		rsv = 0
	}

	k.k = (rsv + 2*k.k) / float64(k.n2)
	k.d = (k.k + 2*k.d) / float64(k.n3)
	j := 3*k.k - 2*k.d

	return k.k, k.d, j
}

func (k *Kdj) maxHigh() float64 {
	max_ := k.dequeH.Front().Value.(float64)
	for e := k.dequeH.Front(); e != nil; e = e.Next() {
		if e.Value.(float64) > max_ {
			max_ = e.Value.(float64)
		}
	}
	return max_
}

func (k *Kdj) minLow() float64 {
	min_ := k.dequeL.Front().Value.(float64)
	for e := k.dequeL.Front(); e != nil; e = e.Next() {
		if e.Value.(float64) < min_ {
			min_ = e.Value.(float64)
		}
	}
	return min_
}

// Get returns the current K, D, and J values
func (k *Kdj) Get() (float64, float64, float64) {
	j := 3*k.k - 2*k.d
	return k.k, k.d, j
}

// Clone creates and returns a copy of the current KDJ object
func (k *Kdj) Clone() *Kdj {
	clone := &Kdj{
		n1:     k.n1,
		n2:     k.n2,
		n3:     k.n3,
		dequeH: list.New(),
		dequeL: list.New(),
		k:      k.k,
		d:      k.d,
	}
	for e := k.dequeH.Front(); e != nil; e = e.Next() {
		clone.dequeH.PushBack(e.Value.(float64))
	}
	for e := k.dequeL.Front(); e != nil; e = e.Next() {
		clone.dequeL.PushBack(e.Value.(float64))
	}
	return clone
}
