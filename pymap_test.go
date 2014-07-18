package pymap

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

var (
	t1, t2 time.Duration
)

func TestSumme_IntArray(t *testing.T) {

	Convey("Given a []int {1,2,3}", t, func() {
		a := []int{1, 2, 3}

		Convey("When summed up", func() {
			s := sumIntArray(&a)

			Convey("The sum should be 6", nil)
			So(s, ShouldEqual, 6)

		})

	})

}

func TestSumme_UintArray(t *testing.T) {

	Convey("Given a []uint {1,2,3}", t, func() {
		a := []uint{1, 2, 3}

		Convey("When summed up", func() {
			s := sumUintArray(&a)

			Convey("The sum should be 6", nil)
			So(s, ShouldEqual, 6)
		})

	})

}

func TestSumme_Int64Array(t *testing.T) {

	Convey("Given a []int64 {1,2,3}", t, func() {
		a := []int64{1, 2, 3}

		Convey("When summed up", func() {
			s := sumInt64Array(&a)

			Convey("The sum should be 6", nil)
			So(s, ShouldEqual, 6)
		})

	})

}

func TestSumme_Uint64Array(t *testing.T) {

	Convey("Given a []uint64 {1,2,3}", t, func() {
		a := []uint64{1, 2, 3}

		Convey("When summed up", func() {
			s := sumUint64Array(&a)

			Convey("The sum should be 6", nil)
			So(s, ShouldEqual, 6)
		})

	})

}

func TestEqualArrays(t *testing.T) {

	Convey("Given a []int {1,2,3}, b []int {1,3,3}, c []int{1,2}, d []int{}", t, func() {
		a := []int{1, 2, 3}
		b := []int{1, 3, 3}
		c := []int{1, 2}
		d := []int{}

		Convey("EqualArrays a, a -> True ", func() {
			equal := EqualArrays(&a, &a)
			Convey("should be true ", nil)
			So(equal, ShouldBeTrue)
		})

		Convey("EqualArrays a, b -> False ", func() {
			equal := EqualArrays(&a, &b)
			Convey("should be false ", nil)
			So(equal, ShouldBeFalse)
		})

		Convey("EqualArrays a, c -> False ", func() {
			equal := EqualArrays(&a, &c)
			Convey("should be false ", nil)
			So(equal, ShouldBeFalse)
		})

		Convey("EqualArrays a, d -> False ", func() {
			equal := EqualArrays(&a, &d)
			Convey("should be false ", nil)
			So(equal, ShouldBeFalse)
		})

	})

}

func TestEach(t *testing.T) {

	Convey("Given a []int {1,2,3}", t, func() {
		a := []int{}
		for i := 0; i < 100000000; i++ {
			a = append(a, i)
		}

		Convey("Each x*x", func() {
			t1 = Each(&a, func(a *[]int, i, x, y int) {
				(*a)[i] = x * y
			})

			Convey("Now a should be 1, 4, 9 ", nil)
			fmt.Println(t1)
			// So(a[0], ShouldEqual, 1)
			// So(a[1], ShouldEqual, 4)
			// So(a[2], ShouldEqual, 9)
		})

	})

}

func TestEachNew(t *testing.T) {

	Convey("Given a [1000]int {1 .. 1000}, fn(y) = x*x", t, func() {
		a := []int{}
		for i := 0; i < 1000; i++ {
			a = append(a, i)
		}

		Convey("return b = a**2", func() {
			b := EachNew(&a, func(x, y int) int {
				return x * y
			})

			Convey("Now b should be Each(a, fn=x*x) ", nil)
			t1 = Each(&a, func(a *[]int, i, x, y int) {
				(*a)[i] = x * y
			})

			So(EqualArrays(&b, &a), ShouldBeTrue)
			// type bi interface {
			// 	Len
			// }
			// bii := bi(b)
			// switch i := bi.(type) {
			// case int:
			// 	fmt.Println(i)
			// default:
			// 	fmt.Println(i)
			// }

			// So(a[1], ShouldEqual, 4)
			// So(a[2], ShouldEqual, 9)
		})

	})

}

func TestEach2Seq(t *testing.T) {

	Convey("Given a []int {1,2,3}", t, func() {
		a := []int{}
		for i := 0; i < 1000; i++ {
			a = append(a, i)
		}
		b := make([]int, len(a))
		copy(b, a)

		Convey("Each s1[i]*s2[i]", func() {
			t2 = Each2Seq(&b, &a, func(s1, s2 *[]int, i int) {
				(*s1)[i] *= (*s2)[i]
			})
			t1 = Each(&a, func(a *[]int, i, x, y int) {
				(*a)[i] = x * y
			})

			Convey("Now a should be 2, 8, 18 ", nil)
			fmt.Println(t2, t1 < t2)
			fmt.Println(EqualArrays(&a, &b))
			// So(a[0], ShouldEqual, 2)
			// So(a[1], ShouldEqual, 8)
			// So(a[2], ShouldEqual, 18)
		})

	})

}
