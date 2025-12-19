package models

import "fmt"

// 矩形
type Rectangle struct {
	// 长
	Length float64
	// 宽
	Width float64
}

func (r Rectangle) Area() {
	fmt.Printf("矩形面积: %.2f\n", r.Length*r.Width)

}

func (r Rectangle) Perimeter() {
	fmt.Printf("矩形周长: %.2f\n", 2*(r.Length+r.Width))
}
