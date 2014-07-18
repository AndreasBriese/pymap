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
	for _, v := range *a {
		s += v
	}
	return s
}

func sumUintArray(a *[]uint) (s uint) {
	for _, v := range *a {
		s += v
	}
	return s
}

func sumInt64Array(a *[]int64) (s int64) {
	for _, v := range *a {
		s += v
	}
	return s
}

func sumUint64Array(a *[]uint64) (s uint64) {
	for _, v := range *a {
		s += v
	}
	return s
}

func Each(a *[]int, fn func(*[]int, int, int, int)) time.Duration {
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
