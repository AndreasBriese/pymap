package pymap

import (
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(4)
}

func sumIntArray(a *[]int) (s int) {
	// sumIntArray 
	// takes a pointer to an array/slice of type []int 
	// (will be int32/int64 depending on OS & compiler settings)
	// returns the sum of the values`
	for _, v := range *a {
		s += v
	}
	return s
}

func sumUintArray(a *[]uint) (s uint) {
	// sumUintArray 
	// takes a pointer to an array/slice of type []uint 
	// (will be uint32/uint64 depending on OS & compiler settings)
	// returns the sum of the values`
	for _, v := range *a {
		s += v
	}
	return s
}

func sumInt64Array(a *[]int64) (s int64) {
	// sumInt64Array 
	// takes a pointer to an array/slice of type []int64
	// (allows casting of int64 values if int was compiled as int32)
	// returns the sum of the values`
	for _, v := range *a {
		s += v
	}
	return s
}

func sumUint64Array(a *[]uint64) (s uint64) {
	// sumUint64Array 
	// takes a pointer to an array/slice of type []uint64
	// (allows casting of uint64 values if uint was compiled as uint32)
	// returns the sum of the values`
	for _, v := range *a {
		s += v
	}
	return s
}

func EqualArrays(a1, a2 *[]int) bool {
	// EqualArrays takes 
	// 	 two pointer to arrays/slices 'a1', 'a2' of type []int 
	// returns true if both have equal field content -> bool 
	if len(*a1) != len(*a2) {
		return false
	}
	for i, v := range *a1 {
		if v != (*a2)[i] {
			return false
		}
	}
	return true
}

func Each(a *[]int, fn func(*[]int, int, int, int)) {
	// Each takes 
	// 	 a pointer to an array/slice 'a' of type []int
	//      Attention!  The array/slice 'a' will become modified by holding the results
	//					If you want to preserve the array, use func EachNew instead 
	// 	 a func to process each value in the slice/array 'a'
	// 		func takes 
	// 			a pointer to 'a'
	// 			an index 'i' 
	// 			two ints to be used within func
	// 				i.e. a[i] = x */-+ y
	for i, v := range *a {
		fn(a, i, v, v)
	}
}

func EachNew(a *[]int, fn func(int, int) int) (e []int) {
	// EachNew takes 
	// 	 a pointer to an array/slice 'a' of type []int 
	//		the results will be hold in 'e' []int
	//   	'a' will be preserved - if you want the results beeing stored within 'a' use func Each instead 
	// 	 a func to process each value in the slice/array 'a'
	// 		func takes 
	// 			two ints to be used within func
	// 				i.e. e[i] = v/x */-+ y
	// returns 'e' []int containing the results
	e = make([]int, len(*a))
	for i, v := range *a {
		e[i] = fn(v, v)
	}
	return e
}

func Each2Seq(a1, a2 *[]int, fn func(*[]int, *[]int, int)) {
	// Each2Seq takes 
	// 	 two pointer to arrays/slices 'a1', 'a2' of type []int 
	//	 Attention!  The array/slice 'a1' will become modified by holding the results
	//				If you want to preserve the array 'a', use func EachNew2Seq instead 
	// 	 a func to process each value in the slice/array 'a1' with each corresponding value in 'a2'
	// 		func takes 
	//			two pointers to 'a1', 'a2'
	// 			an index 'i' 
	for i, _ := range *a1 {
		fn(a1, a2, i)
	}
}
