package main

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Basics1(t *testing.T) {
	gen := NewRndGen()

	value := gen.Uint64()
	typeOf := reflect.TypeOf(value)

	fmt.Printf("Here is a random [%s] number: [%d]\n", typeOf, value)
}

func Test_Basics2(t *testing.T) {
	gen := NewRndGen()

	value := gen.Float64()
	sqrt := math.Sqrt(value)

	fmt.Printf("The root square of [%.4f] number is [%.4f]\n", value, sqrt)
}

func Test_Basics3(t *testing.T) {
	fmt.Printf("The first eight digits of Pi are [%.8f]\n", math.Pi)
}

func Test_Basics4(t *testing.T) {
	gen := NewRndGen()

	x := gen.Int()
	y := gen.Int()

	assert.Equal(t, x+y, AddBasics4(x, y))
}

func AddBasics4(x int, y int) int {
	return x + y
}

func Test_Basics5(t *testing.T) {
	gen := NewRndGen()

	x := gen.Int()
	y := gen.Int()

	assert.Equal(t, x+y, AddBasics5(x, y))
}

func AddBasics5(x, y int) int {
	return x + y
}

func Test_Basics6(t *testing.T) {
	const s1 string = "Hello"
	const s2 string = "World"

	s3, s4 := SwapBasics6(s1, s2)

	assert.Equal(t, s2, s3)
	assert.Equal(t, s1, s4)
}

func SwapBasics6(s1, s2 string) (string, string) {
	return s2, s1
}

func Test_Basics7(t *testing.T) {
	gen := NewRndGen()

	sum := gen.Int()
	x, y := SplitBasics7(sum)

	assert.Equal(t, sum*4/9, x)
	assert.Equal(t, sum-(sum*4/9), y)
}

func SplitBasics7(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func Test_Basics8(t *testing.T) {
	var a, b, c uint32

	assert.Equal(t, reflect.Uint32, reflect.TypeOf(a).Kind())
	assert.Equal(t, reflect.TypeOf(a), reflect.TypeOf(b))
	assert.Equal(t, reflect.TypeOf(b), reflect.TypeOf(c))
}

func Test_Basics9(t *testing.T) {
	const a, b = true, "test"

	assert.Equal(t, reflect.Bool, reflect.TypeOf(a).Kind())
	assert.Equal(t, reflect.String, reflect.TypeOf(b).Kind())
}

func Test_Basics10(t *testing.T) {
	a, b := 42, "test"

	assert.Equal(t, reflect.Int, reflect.TypeOf(a).Kind())
	assert.Equal(t, reflect.String, reflect.TypeOf(b).Kind())
}

func Test_Basics11(t *testing.T) {
	var (
		a = 42
		b = "test"
	)

	assert.Equal(t, reflect.Int, reflect.TypeOf(a).Kind())
	assert.Equal(t, reflect.String, reflect.TypeOf(b).Kind())
}

func Test_Basics12(t *testing.T) {
	var (
		a int
		b string
		c bool
	)

	assert.Equal(t, 0, a)
	assert.Equal(t, "", b)
	assert.Equal(t, false, c)
}

func Test_Basics13(t *testing.T) {
	var f = math.Pi
	var i = int(f)

	assert.Equal(t, 3, i)
}

func Test_Basics14(t *testing.T) {
	const f1 float64 = math.Pi
	const f2 = f1

	assert.Equal(t, reflect.Float64, reflect.TypeOf(f1).Kind())
	assert.Equal(t, reflect.TypeOf(f1), reflect.TypeOf(f2))
}

func Test_Basics15(t *testing.T) {
	const f float64 = math.Pi

	assert.Equal(t, reflect.Float64, reflect.TypeOf(f).Kind())
}

func Test_Basics16(t *testing.T) {
	const i uint8 = (1 << 128) >> 127

	assert.Equal(t, reflect.Uint8, reflect.TypeOf(i).Kind())
}
