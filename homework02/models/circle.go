package models

import (
	"fmt"
	"math"
)

// 圆
type Circle struct {
	// 半径
	Radius float64
}

func (c Circle) Area() {
	fmt.Printf("圆面积: %.2f\n", math.Pi*c.Radius*c.Radius)
}

// Perimeter 计算圆形周长
func (c Circle) Perimeter() {
	fmt.Printf("圆周长: %.2f\n", 2*math.Pi*c.Radius)
}
