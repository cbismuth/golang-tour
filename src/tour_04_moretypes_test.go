package main

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

type Vertex struct {
	X, Y int
}

type LatLong struct {
	Lat, Long float64
}

const pin1 string = "Bell Labs"
const lat1 float64 = 40.68433
const long1 float64 = -74.39967

const pin2 string = "Google"
const lat2 float64 = 37.42202
const long2 float64 = -122.08408

func Test_MoreTypes1(t *testing.T) {
	gen := NewRndGen()

	i := gen.Int()

	p := &i
	assert.Equal(t, i, *p)
	assert.NotEqual(t, i, p)

	*p = 42
	assert.Equal(t, i, 42)
}

func Test_MoreTypes2(t *testing.T) {
	v := Vertex{1, 2}

	assert.Equal(t, 1, v.X)
	assert.Equal(t, 2, v.Y)
}

func Test_MoreTypes3(t *testing.T) {
	v := Vertex{1, 2}

	v.X = 3
	v.Y = 4

	assert.Equal(t, 3, v.X)
	assert.Equal(t, 4, v.Y)
}

func Test_MoreTypes4(t *testing.T) {
	v := Vertex{1, 2}

	p := &v

	assert.Equal(t, 1, (*p).X)
	assert.Equal(t, 2, (*p).Y)

	assert.Equal(t, 1, p.X)
	assert.Equal(t, 2, p.Y)
}

func Test_MoreTypes5(t *testing.T) {
	v1 := Vertex{1, 2}
	assert.Equal(t, 1, v1.X)
	assert.Equal(t, 2, v1.Y)

	v2 := Vertex{X: 1}
	assert.Equal(t, 1, v2.X)
	assert.Equal(t, 0, v2.Y)

	v3 := Vertex{X: 1, Y: 2}
	assert.Equal(t, 1, v3.X)
	assert.Equal(t, 2, v3.Y)

	v4 := Vertex{}
	assert.Equal(t, 0, v4.X)
	assert.Equal(t, 0, v4.Y)

	p := &Vertex{}
	assert.Equal(t, 0, (*p).X)
	assert.Equal(t, 0, (*p).Y)
	assert.Equal(t, 0, p.X)
	assert.Equal(t, 0, p.Y)
	assert.Equal(t, "*main.Vertex", reflect.TypeOf(p).String())
	assert.Equal(t, "main.Vertex", reflect.TypeOf(*p).String())
}

func Test_MoreTypes6(t *testing.T) {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	assert.Equal(t, "Hello World", strings.Join(a[:], " "))
}

func Test_MoreTypes7(t *testing.T) {
	a := [2]string{"Hello", "World"}

	assert.Equal(t, "Hello", a[0])
	assert.Equal(t, "World", a[1])
}

func Test_MoreTypes8(t *testing.T) {
	a := []string{"a", "b", "c", "d"}

	assert.Equal(t, []string{"c"}, a[2:3])
	assert.Equal(t, []string{"a", "b", "c", "d"}, a)

	assert.True(t, reflect.DeepEqual([]string{"a", "b", "c", "d"}, a))
}

func Test_MoreTypes9(t *testing.T) {
	a := []string{"Hello", "World"}

	assert.Equal(t, "Hello World", strings.Join(a[:], " "))
}

func Test_MoreTypes10(t *testing.T) {
	a := []string{"a", "b", "c", "d"}

	assert.Equal(t, []string{"a", "b"}, a[:2])
	assert.Equal(t, []string{"c", "d"}, a[2:])
}

func Test_MoreTypes11(t *testing.T) {
	a := []string{"a", "b", "c", "d"}
	assert.Equal(t, 4, len(a))
	assert.GreaterOrEqual(t, cap(a), 4)
	assert.Equal(t, 1, len(a[:1]))
	assert.GreaterOrEqual(t, cap(a[:1]), 4)

	a = a[:1]
	assert.Equal(t, []string{"a"}, a)

	a = a[:4]
	assert.Equal(t, []string{"a", "b", "c", "d"}, a)
}

func Test_MoreTypes12(t *testing.T) {
	var a []int
	assert.Nil(t, a)
	assert.Equal(t, 0, len(a))
	assert.Equal(t, 0, cap(a))
}

func Test_MoreTypes13(t *testing.T) {
	gen := NewRndGen()

	l := gen.Intn(1000)
	a := make([]int, l)

	assert.Equal(t, l, len(a))
	assert.GreaterOrEqual(t, cap(a), l)
}

func Test_MoreTypes14(t *testing.T) {
	a := [][]string{
		{"a", "b"},
		{"c", "d", "e"},
	}

	assert.Equal(t, "a", a[0][0])
	assert.Equal(t, "b", a[0][1])
	assert.Equal(t, "c", a[1][0])
	assert.Equal(t, "d", a[1][1])
	assert.Equal(t, "e", a[1][2])
}

func Test_MoreTypes15(t *testing.T) {
	a := []string{"a", "b"}
	assert.Equal(t, "a", a[0])
	assert.Equal(t, "b", a[1])
	assert.Equal(t, 2, len(a))
	assert.GreaterOrEqual(t, cap(a), 2)

	a = append(a, "c")
	assert.Equal(t, "a", a[0])
	assert.Equal(t, "b", a[1])
	assert.Equal(t, "c", a[2])
	assert.Equal(t, 3, len(a))
	assert.GreaterOrEqual(t, cap(a), 4)
}

func Test_MoreTypes16(t *testing.T) {
	gen := NewRndGen()

	delta := gen.Intn(1000)

	var a []int
	for i := 0; i < gen.Intn(100); i++ {
		a = append(a, i+delta)
	}

	acc := 0
	for i, v := range a {
		assert.Equal(t, acc, i)
		assert.Equal(t, acc+delta, v)
		acc++
	}
}

func Test_MoreTypes17(t *testing.T) {
	gen := NewRndGen()

	delta := gen.Intn(1000)

	var a []int
	for i := 0; i < gen.Intn(100); i++ {
		a = append(a, i+delta)
	}

	acc := 0
	for _, v := range a {
		assert.Equal(t, acc+delta, v)
		acc++
	}
}

func Test_MoreTypes18(t *testing.T) {
	pic.Show(FnMoreTypes18)
}

func FnMoreTypes18(dx, dy int) [][]uint8 {
	var a [][]uint8

	a = make([][]uint8, dx)
	for x := range dx {
		a[x] = make([]uint8, dy)
		for y := range dy {
			a[x][y] = uint8((x + y) / 2)
		}
	}

	return a
}

func Test_MoreTypes19(t *testing.T) {
	m := make(map[string]LatLong)

	m[pin1] = LatLong{lat1, long1}
	assert.Equal(t, lat1, m[pin1].Lat)
	assert.Equal(t, long1, m[pin1].Long)
}

func Test_MoreTypes20(t *testing.T) {
	m := map[string]LatLong{
		pin1: {lat1, long1},
		pin2: {lat2, long2},
	}

	assert.Equal(t, lat1, m[pin1].Lat)
	assert.Equal(t, long1, m[pin1].Long)

	assert.Equal(t, lat2, m[pin2].Lat)
	assert.Equal(t, long2, m[pin2].Long)
}

func Test_MoreTypes21(t *testing.T) {
	// Value type already omitted from map entry declarations

	assert.True(t, true)
}

func Test_MoreTypes22(t *testing.T) {
	m := make(map[string]LatLong)

	m[pin1] = LatLong{lat1, long1}
	assert.Equal(t, lat1, m[pin1].Lat)
	assert.Equal(t, long1, m[pin1].Long)

	e, ok := m[pin1]
	assert.True(t, ok)
	assert.Equal(t, lat1, e.Lat)
	assert.Equal(t, long1, e.Long)

	delete(m, pin1)
	assert.Equal(t, 0.0, m[pin1].Lat)
	assert.Equal(t, 0.0, m[pin1].Long)
}

func Test_MoreTypes23(t *testing.T) {
	wc.Test(WordCountMoreTypes23)
}

func WordCountMoreTypes23(s string) map[string]int {
	m := make(map[string]int)

	for _, field := range strings.Fields(s) {
		m[field] = m[field] + 1
	}

	return m
}

func Test_MoreTypes24(t *testing.T) {
	multiply := func(x, y int) int { return x * y }

	assert.Equal(t, 12, ExecuteMoreTypes24(multiply, 3, 4))
}

func ExecuteMoreTypes24(fn func(int, int) int, x, y int) int {
	return fn(x, y)
}

func Test_MoreTypes25(t *testing.T) {
	fn := AddMoreTypes25()

	for i := 0; i < 10; i++ {
		assert.Equal(t, i+1, fn(1))
	}
}

func AddMoreTypes25() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func Test_MoreTypes26(t *testing.T) {
	fn := FibonacciMoreTypes26()

	a := make([]int, 10)
	for i := 0; i < 10; i++ {
		a[i] = fn()
	}

	assert.Equal(t, []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}, a)
}

func FibonacciMoreTypes26() func() int {
	n := 0
	fn2 := 0
	fn1 := 1

	return func() int {
		var fn int
		switch n {
		case 0:
			fn = fn2
		case 1:
			fn = fn1
		default:
			fn = fn2 + fn1
			fn2 = fn1
			fn1 = fn
		}
		n++
		return fn
	}
}
