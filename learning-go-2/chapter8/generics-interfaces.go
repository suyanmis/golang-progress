package main

import (
	"errors"
	"fmt"
	"math"
)

type Pair[T fmt.Stringer] struct {
	Val1 T
	Val2 T
}

type Differ[T any] interface {
	fmt.Stringer
	Diff(T) float64
}
type Point2D struct {
	X, Y int
}
type Point3D struct {
	X, Y, Z int
}

func (p2 Point2D) String() string {
	return fmt.Sprintf("{%d, %d}", p2.X, p2.Y)
}
func (p3 Point3D) String() string {
	return fmt.Sprintf("{%d, %d, %d}", p3.X, p3.Y, p3.Z)
}

func (p1 Point2D) Diff(p2 Point2D) float64 {
	x := float64(p1.X - p2.X)
	y := float64(p1.Y - p2.Y)
	return math.Sqrt(x*x + y*y)
}

func (p1 Point3D) Diff(p2 Point3D) float64 {
	x := float64(p1.X - p2.X)
	y := float64(p1.Y - p2.Y)
	z := float64(p1.Z - p2.Z)
	return math.Sqrt(x*x + y*y + z*z)
}
func FindCloser[T Differ[T]](pair1, pair2 Pair[T]) Pair[T] {
	d1 := pair1.Val1.Diff(pair1.Val2)
	d2 := pair2.Val1.Diff(pair2.Val2)
	if d1 < d2 {
		return pair1
	}
	return pair2
}

func main() {
	pair2Da := Pair[Point2D]{Point2D{1, 1}, Point2D{5, 5}}
	pair2Db := Pair[Point2D]{Point2D{10, 10}, Point2D{15, 5}}
	closer := FindCloser(pair2Da, pair2Db)
	fmt.Println(closer)

	pair3Da := Pair[Point3D]{Point3D{1, 1, 10}, Point3D{5, 5, 0}}
	pair3Db := Pair[Point3D]{Point3D{10, 10, 10}, Point3D{11, 5, 0}}
	closer2 := FindCloser(pair3Da, pair3Db)
	fmt.Println(closer2)

	fmt.Println(divideAndReminder(15, 4))

	mi1 := MyInt(50)
	mi2 := MyInt(11)

	fmt.Println(divideAndReminder(mi1, mi2))

	// s := ImpossibleStruct[int]{10} doesn't work missing String method
	// s2 := ImpossibleStruct[MyInt]{10} doesn't work missing ~
	// s3 := PossibleStruct[MyInt]{10} // works because of the ~

	fmt.Printf("Convert[int32, int64](50000000): %v\n", Convert[int32, int8](50000000))
}

// if the type declared with tilde ~ it will accept the custom types generated from those values
type Integer interface {
	~int | uint | uint16 | uint8 | uint32 | uint64 | uintptr | int32 | int8 | int16 | int64
}
type MyInt int

func divideAndReminder[T Integer](num, denom T) (T, T, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}

type PrintableInt interface {
	String() string
	~int
}

type ImposPrintableInt interface {
	int
	String() string
}

type ImpossibleStruct[T ImposPrintableInt] struct {
	val T
}
type PossibleStruct[T PrintableInt] struct {
	val T
}

func (mi MyInt) String() string {
	return fmt.Sprintf("%d", mi)
}

func Convert[T1, T2 Integer](in T1) T2 {
	return T2(in)
}

func PlusMaxPossible[T Integer](in T) T {
	// return in + 1000 fails 1000 > uint8&int8
	return in + 127 // max
	// return in - 127 // min

}
