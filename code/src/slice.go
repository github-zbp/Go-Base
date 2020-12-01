package main

import "fmt"

func main() {
	// example1
	//arr := [10] int {1,2,3,4,5,6,7,8,9,10}	// 创建一个有10个长度的数组
	//s1 := arr[:5]		// 基于arr的前5个元素创建一个切片
	//fmt.Println(len(s1), cap(s1))		// 查看其长度和容量

	// example2
	//s2 := make([]int, 5)	// 创建一个长度和容量都为5的切片，Go内部会创建一个大小为5，内容都是0的数组
	//fmt.Println(len(s2), cap(s2), s2)
	//
	//s3 := make([]int, 5, 10)	// 创建一个长度和容量都为5的切片，Go内部会创建一个大小为10，内容都是0的数组
	//fmt.Println(len(s3), cap(s3), s3)

	// example3
	//s4 := []int {1,2,3,4,5}		// 初始化一个切片
	//fmt.Println(len(s4), cap(s4), s4)	// 5 5 [1 2 3 4 5]
	//
	//// 注意，这里和初始化数组很想
	//arr2 := [5]int {1,2,3,4,5}	// 初始化一个数组
	//arr3 := [...]int {1,2,3,4,5}	// 初始化一个数组,和上一句没有区别
	//fmt.Println(arr2, arr3)

	// example4
	//s5 := []int {1,2,3,4,5,6,7,8,9,10}	// 创建一个有10个长度和容量的切片
	//s6 := s5[:5]
	//fmt.Println(len(s6), cap(s6), s6)	//	5 10 [1 2 3 4 5]
	//s7 := s5[3:5:8]
	//fmt.Println(len(s7), cap(s7), s7)	//	2 5 [4 5]

	// example5
	//s8 := make([]int, 3, 5)		// 创建一个长度为3，容量为5的切片，Go内部会创建一个大小为5，内容都是0的数组
	//s9 := s8
	//s8 = append(s8, 100, 100)
	//fmt.Println(len(s8), cap(s8), s8)	// 5 5 [0 0 0 100 100]
	//fmt.Println(len(s9), cap(s9), s9)	// 3 5 [0 0 0]
	//
	//
	//s8 = append(s8, 100)
	//fmt.Println(len(s8), cap(s8), s8)	// 6 10 [0 0 0 100 100 100]
	//fmt.Println(len(s9), cap(s9), s9)	// 3 5 [0 0 0]
	//
	//s8[1] = 10
	//fmt.Println(s8)	// [0 10 0 100 100 100]
	//fmt.Println(s9) // [0 0 0]

	// example6
	//s8 := make([]int, 3, 5)		// 创建一个长度为3，容量为5的切片，Go内部会创建一个大小为5，内容都是0的数组
	//s9 := s8
	//s8 = append(s8, 100, 100)
	//fmt.Println(len(s8), cap(s8), s8)	// 5 5 [0 0 0 100 100]
	//
	//s9 = append(s9, 50)
	//fmt.Println(len(s9), cap(s9), s9)	// 4 5 [0 0 0 50]
	//fmt.Println(len(s8), cap(s8), s8)	// 5 5 [0 0 0 50 100]

	// example7
	//s5 := []int {1,2,3,4,5,6,7,8,9,10}	// 创建一个有10个长度和容量的切片
	//s7 := s5[3:5:8]
	//fmt.Println(len(s7), cap(s7), s7)	//	2 5 [4 5]
	//s7 = append(s7, 100, 100, 100, 100)
	//fmt.Println(len(s7), cap(s7), s7)	//	 6 10 [4 5 100 100 100 100]
	//fmt.Println(s5)	// [1 2 3 4 5 6 7 8 9 10]

	// example8
	//s5 := []int {1,2,3,4,5,6,7,8,9,10}	// 创建一个有10个长度和容量的切片
	//s7 := s5[3:5:8]
	//fmt.Println(len(s7), cap(s7), s7)	//	2 5 [4 5]
	//s7 = append(s7, 100, 100, 100)
	//s7 = append(s7, 100)
	//fmt.Println(len(s7), cap(s7), s7)	//	 6 10 [4 5 100 100 100 100]
	//fmt.Println(s5)	// [1 2 3 4 5 100 100 100 9 10]

	// example9
	//arr := [5]int {1,2,3,4,5}
	//s10 := arr[:3]
	//s11 := s10[:4]
	//fmt.Println(len(s11), cap(s11), s11)	// 4 5 [1 2 3 4]

	//s12 := s10[:6]		// 报错
	//fmt.Println(len(s11), cap(s11), s11)

	//s13 := s10[:3:3]
	//fmt.Println(len(s13), cap(s13), s13)	// 3 3 [1 2 3]
	//s14 := s13[:4]		// 报错
	//fmt.Println(len(s14), cap(s14), s14)	// 4 5 [1 2 3 4]

	// example10
	//s11 := []int {7, 8, 9}
	//s12 := []int {1,2,3,4,5}
	//copy(s12, s11)
	//fmt.Println(cap(s11), s11, cap(s12), s12)	// 3 [7 8 9] 5 [7 8 9 4 5]
	//s12[1] = 100
	//fmt.Println(cap(s11), s11, cap(s12), s12)	// 3 [7 8 9] 5 [7 100 9 4 5]

	//s13 := []int {1,2,3,4,5}
	//s14 := s13[2:]
	//fmt.Println(cap(s13), s13, cap(s14), s14)		// 5 [1 2 3 4 5] 3 [3 4 5]
	//copy(s13, s14)
	//fmt.Println(cap(s13), s13, cap(s14), s14)		// 5 [3 4 5 4 5] 3 [5 4 5]

	// example11
	//var s0 []string
	//fmt.Println(len(s0), cap(s0), s0)

	// example12
	//s14 := []int {1,2,3,4,5}

	// example13


	//s15 := []int {1,2,3,4,5}
	//reverse(s15)
	//fmt.Println(s15)

	//s16 := []int {1,2,3,4,5}
	//s17 := remove(s16, 2)
	//fmt.Println(s17, cap(s17))
	//s17[0] = 100
	//fmt.Println(s16)
	//s18 := make([]int, 2, 4)
	//print(s18[1], cap(s18))

	//s19 := []int {1,2,3,4,5}
	//s20 := s19[:2]
	//fmt.Println(cap(s20))
	//fmt.Println(s20[4])

	s21 := []int {1,2,3,4,5}
	s22 := []int {1,2,31,4,5}
	s23 := []int {1,2,3,4,5}
	s24 := s21[:]
	fmt.Println(equal(s21, s22))
	fmt.Println(equal(s21, s23))
	fmt.Println(equal(s21, s24))

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 删除s中的第i个元素，要求不影响原切片
// 思路很简单，只需将第i+1到最后一个元素copy或者append到第i个元素到最后一个元素即可
//func remove(s []int, i int) (new_s []int) {
//	new_s = append(new_s, s[:i]...)
//	new_s = append(new_s, s[i+1:]...)
//	return new_s
//}

//func remove(s []int, i int) []int {
//	copy(s[i:], s[i+1:])
//	return s[:len(s)-1]
//}

func remove(s []int, i int) []int {
	s[i] = s[len(s) -1]
	return s[:len(s) - 1]
}

func equal(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i]{
			return false
		}
	}

	return true
}