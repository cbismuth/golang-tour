package main

import (
	"fmt"
	"image"
	"io"
	"math"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/tour/pic"
)

func Test_Methods1(t *testing.T) {
	input := Google

	expected := 127.69
	actual := toFixed(input.AbsMethods1(), 2)

	assert.Equal(t, expected, actual)
}

//goland:noinspection GoMixedReceiverTypes
func (v Vertex) AbsMethods1() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Test_Methods2(t *testing.T) {
	input := Google

	expected := 127.69
	actual := toFixed(AbsMethods2(input), 2)

	assert.Equal(t, expected, actual)
}

func AbsMethods2(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Test_Methods3(t *testing.T) {
	assert.Equal(t, 42.0, MyFloat64(42.0).Abs())
	assert.Equal(t, 42.0, MyFloat64(-42.0).Abs())
}

func Test_Methods4(t *testing.T) {
	v1 := &Vertex{-2.0, -3.0}
	v1.AbsXYMethods41()
	assert.Equal(t, &Vertex{2.0, 3.0}, v1)

	v2 := Vertex{-2.0, -3.0}
	v2.AbsXYMethods42()
	assert.Equal(t, Vertex{-2.0, -3.0}, v2)
}

//goland:noinspection GoMixedReceiverTypes
func (v *Vertex) AbsXYMethods41() {
	v.X = math.Abs(v.X)
	v.Y = math.Abs(v.Y)
}

//goland:noinspection GoMixedReceiverTypes
func (v Vertex) AbsXYMethods42() {
	v.X = math.Abs(v.X)
	v.Y = math.Abs(v.Y)
}

func Test_Methods5(t *testing.T) {
	v1 := &Vertex{-2.0, -3.0}
	AbsXYMethods51(v1)
	assert.Equal(t, &Vertex{2.0, 3.0}, v1)

	v2 := Vertex{-2.0, -3.0}
	AbsXYMethods52(v2)
	assert.Equal(t, Vertex{-2.0, -3.0}, v2)
}

//goland:noinspection GoMixedReceiverTypes
func AbsXYMethods51(v *Vertex) {
	v.X = math.Abs(v.X)
	v.Y = math.Abs(v.Y)
}

//goland:noinspection GoMixedReceiverTypes
func AbsXYMethods52(v Vertex) {
	v.X = math.Abs(v.X)
	v.Y = math.Abs(v.Y)
}

func Test_Methods6(t *testing.T) {
	v1 := &Vertex{-2.0, -3.0}
	v1.AbsXYMethods6()
	assert.Equal(t, &Vertex{2.0, 3.0}, v1)

	v2 := Vertex{-2.0, -3.0}
	v2.AbsXYMethods6()
	assert.Equal(t, Vertex{2.0, 3.0}, v2)
}

//goland:noinspection GoMixedReceiverTypes
func (v *Vertex) AbsXYMethods6() {
	v.X = math.Abs(v.X)
	v.Y = math.Abs(v.Y)
}

func Test_Methods7(t *testing.T) {
	v1 := &Vertex{-2.0, -3.0}
	AbsXYMethods71(*v1)
	assert.Equal(t, &Vertex{-2.0, -3.0}, v1)

	v2 := Vertex{-2.0, -3.0}
	AbsXYMethods71(v2)
	assert.Equal(t, Vertex{-2.0, -3.0}, v2)

	v3 := &Vertex{-2.0, -3.0}
	v3.AbsXYMethods72()
	assert.Equal(t, &Vertex{2.0, 3.0}, v3)

	v4 := Vertex{-2.0, -3.0}
	v4.AbsXYMethods72()
	assert.Equal(t, Vertex{2.0, 3.0}, v4)
}

//goland:noinspection GoMixedReceiverTypes
func AbsXYMethods71(v Vertex) {
	v.X = math.Abs(v.X)
	v.Y = math.Abs(v.Y)
}

//goland:noinspection GoMixedReceiverTypes
func (v *Vertex) AbsXYMethods72() {
	v.X = math.Abs(v.X)
	v.Y = math.Abs(v.Y)
}

func Test_Methods8(t *testing.T) {
	var v Vertex
	err := DeepCopy(&Google, &v)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	v.Scale(10.0)

	assert.Equal(t, 1276.9, toFixed(v.Abs(), 2))
}

func Test_Methods9(t *testing.T) {
	v := Vertex{}
	absType := reflect.TypeOf((*Abser)(nil)).Elem()
	implements := reflect.TypeOf(&v).Implements(absType)
	assert.True(t, implements)

	assert.Implements(t, (*Abser)(nil), new(Vertex))
}

func Test_Methods10(t *testing.T) {
	v1 := Google
	v2 := &Google

	assert.Equal(t, 127.69, toFixed(v1.Abs(), 2))
	assert.Equal(t, 127.69, toFixed(v2.Abs(), 2))
}

func Test_Methods11(t *testing.T) {
	v1 := Google
	v2 := &Google
	v3 := MyFloat64(-42.0)

	assert.Equal(t, 127.69, toFixed(v1.Abs(), 2))
	assert.Equal(t, 127.69, toFixed(v2.Abs(), 2))
	assert.Equal(t, 42.0, v3.Abs())
}

func Test_Methods12(t *testing.T) {
	var v Vertex

	assert.Equal(t, 0.0, toFixed(v.Abs(), 2))
}

func Test_Methods13(t *testing.T) {
	assert.Panics(t, func() {
		var i I
		i.nop()
	})
}

func Test_Methods14(t *testing.T) {
	// See I interface type declaration

	assert.True(t, true)
}

func Test_Methods15(t *testing.T) {
	var v1 Abser
	v1 = &Google

	v2, ok := v1.(*Vertex)

	assert.True(t, ok)
	assert.Equal(t, &Google, v2)
	assert.Panics(t, func() {
		_ = v1.(I)
	})
}

func Test_Methods16(t *testing.T) {
	var v Abser
	v = &Google

	switch v.(type) {
	case *Vertex:
		// Success
	default:
		assert.Fail(t, "Bad type assertion from switch statement")
	}
}

func Test_Methods17(t *testing.T) {
	assert.Equal(t, "X:37.42, Y:-122.08", Google.String())
}

func Test_Methods18(t *testing.T) {
	assert.Equal(t, "1.2.3.4", (&IPAddr{1, 2, 3, 4}).String())
}

func Test_Methods19(t *testing.T) {
	err := &MyError{time.Now(), "Internal error"}

	assert.Regexp(t, regexp.MustCompile("\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2} Internal error"), fmt.Sprint(err))
}

func Test_Methods20(t *testing.T) {
	sqrt, err := MySqrt(-1.234)

	assert.IsType(t, ErrNegativeSqrt(0.0), err)
	assert.Implements(t, (*error)(nil), err)
	assert.Equal(t, 0.0, sqrt)
	assert.Equal(t, "Cannot sqrt negative number: -1.23", fmt.Sprintf("%s", err))
}

func Test_Methods21(t *testing.T) {
	reader := strings.NewReader(ALPHANUM[10:24])

	bytes := make([]byte, 8)

	n, err := reader.Read(bytes)
	assert.Equal(t, 8, n)
	assert.Equal(t, "abcdefgh", string(bytes[:n]))
	assert.Nil(t, err)

	n, err = reader.Read(bytes)
	assert.Equal(t, 6, n)
	assert.Equal(t, "ijklmn", string(bytes[:n]))
	assert.Nil(t, err)

	n, err = reader.Read(bytes)
	assert.Equal(t, 0, n)
	assert.Equal(t, "", string(bytes[:n]))
	assert.Equal(t, io.EOF, err)
}

func Test_Methods22(t *testing.T) {
	reader := MyReader{14, 0}

	bytes := make([]byte, 8)

	n, err := reader.Read(bytes)
	assert.Equal(t, 8, n)
	assert.Equal(t, "AAAAAAAA", string(bytes[:n]))
	assert.Nil(t, err)

	n, err = reader.Read(bytes)
	assert.Equal(t, 6, n)
	assert.Equal(t, "AAAAAA", string(bytes[:n]))
	assert.Nil(t, err)

	n, err = reader.Read(bytes)
	assert.Equal(t, 0, n)
	assert.Equal(t, "", string(bytes[:n]))
	assert.Equal(t, io.EOF, err)
}

func Test_Methods23(t *testing.T) {
	message := "Lbh penpxrq gur pbqr!"

	r := Rot13Reader{strings.NewReader(message)}
	b := make([]byte, len(message))
	n, err := r.Read(b)

	assert.Nil(t, err)
	assert.Equal(t, len(message), n)
	assert.Equal(t, "You cracked the code!", string(b))
}

func Test_Methods24(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 1024, 768))
	r, g, b, a := img.At(0, 0).RGBA()

	assert.Equal(t, 1024, img.Bounds().Dx())
	assert.Equal(t, 768, img.Bounds().Dy())

	assert.Equal(t, 0, int(r))
	assert.Equal(t, 0, int(g))
	assert.Equal(t, 0, int(b))
	assert.Equal(t, 0, int(a))
}

func Test_Methods25(t *testing.T) {
	i := MyImage{}

	pic.ShowImage(i)
}
