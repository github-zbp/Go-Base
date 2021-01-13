package method

import (
	"image/color"
	"math"
)

type Point struct {X, Y float64}	// 定义一个Point类(表示一个点)
type ColorPoint struct {	// 定义一个有颜色的点类
	Point	// 嵌入Point这个类型，此时会隐式的为ColorPoint生成一个名字也叫Point的成员
	Color color.RGBA	// 颜色成员，类型为color包的RGBA类型（底层其实使一个struct）
}
type Points []*Point

// 向量相加
func (point *Point) Add(anotherPoint *Point) *Point {
	var res *Point
	res = &Point{X:point.X+anotherPoint.X, Y:point.Y+anotherPoint.Y}
	return res
}

// 向量相加
func (point *Point) Sub(anotherPoint *Point) *Point {
	var res *Point
	res = &Point{X:point.X-anotherPoint.X, Y:point.Y-anotherPoint.Y}
	return res
}

func (point *Point) Distance(anotherPoint *Point) float64 {
	return math.Hypot(point.X-anotherPoint.X, point.Y-anotherPoint.Y)
}

// 平移points中的多个点，平移的偏移量是offset，会影响原points得到平移后的点
// add为true表示做向量相加的平移，为false是向量相减的平移
func (points Points) TranslateBy(offset *Point, add bool) {
	var op func(point *Point, anotherPoint *Point) *Point	// 定义了一个没有函数体的闭包，所以这个op是一个零值函数，即nil

	// 为op赋值为方法表达式
	if add {
		op = (*Point).Add
	}else{
		op = (*Point).Sub
	}

	for i := range points{
		points[i] = op(points[i], offset)
	}
}