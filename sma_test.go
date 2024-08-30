package indicator

import (
	"fmt"
	"testing"
)

func TestSma(t *testing.T) {
	// 测试数据
	prices := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedAverages := []float64{
		1.0, // (1)
		1.5, // (1+2)/2
		2.0, // (1+2+3)/3
		2.5, // (1+2+3+4)/4
		3.0, // (1+2+3+4+5)/5
		4.0, // (2+3+4+5+6)/5
		5.0, // (3+4+5+6+7)/5
		6.0, // (4+5+6+7+8)/5
		7.0, // (5+6+7+8+9)/5
		8.0, // (6+7+8+9+10)/5
	}

	// 创建 Sma 对象，period 为 5
	sma := NewSma(5)

	for i, price := range prices {
		avg := sma.Update(price)
		fmt.Println(avg)
		if avg != expectedAverages[i] {
			t.Errorf("TestSma failed at index %d: expected %v, got %v", i, expectedAverages[i], avg)
		}
	}
}
