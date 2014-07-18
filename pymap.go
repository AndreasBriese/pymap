package pymap

import (
	// "log"
	"runtime"
	"time"
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

func Each(a *[]int, fn func(*[]int, int, int, int)) time.Duration {
	`Each 
	takes 
		a pointer to an array/slice 'a' of type []int
		a func to process each value in the slice/array 'a'
			func takes 
				a pointer to 'a'
				an index 'i' 
				two ints to be used within func
					i.e. a[i] = x */-+ y
	returns the time.Duration of processing`
	st := time.Now()
	for i, v := range *a {
		fn(a, i, v, v)
	}
	return (time.Since(st))
}

func EachNew(a *[]int, fn func(int, int) int) (e []int) {
	e = make([]int, len(*a))
	for i, v := range *a {
		e[i] = fn(v, v)
	}
	return e
}

func EqualArrays(a1, a2 *[]int) bool {
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

func Each2Seq(s1, s2 *[]int, fn func(*[]int, *[]int, int)) time.Duration {
	st := time.Now()
	for i, _ := range *s1 {
		fn(s1, s2, i)
	}
	return (time.Since(st))
}
