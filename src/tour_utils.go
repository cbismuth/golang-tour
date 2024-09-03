package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"io"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/tour/tree"
)

type I interface {
	nop() int
}

type Abser interface {
	Abs() float64
}

type MyFloat64 float64

func (f MyFloat64) Abs() float64 {
	var abs float64

	if f < 0 {
		abs = float64(-f)
	} else {
		abs = float64(f)
	}

	return abs
}

func (f MyFloat64) String() string {
	return strconv.FormatFloat(float64(f), 'g', -1, 64)
}

type Vertex struct {
	X, Y float64
}

//goland:noinspection GoMixedReceiverTypes
func (v *Vertex) Abs() float64 {
	if v == nil {
		return 0.0
	}

	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//goland:noinspection GoMixedReceiverTypes
func (v *Vertex) Scale(f float64) {
	if v == nil {
		return
	}

	v.X = v.X * f
	v.Y = v.Y * f
}

//goland:noinspection GoMixedReceiverTypes
func (v *Vertex) String() string {
	return fmt.Sprintf("X:%.2f, Y:%.2f", v.X, v.Y)
}

var BellLabs = Vertex{40.68433, -74.39967}
var Google = Vertex{37.42202, -122.08408}

type IPAddr [4]byte

func (a *IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", a[0], a[1], a[2], a[3])
}

type MyError struct {
	timestamp time.Time
	message   string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%s %s", e.timestamp.Format("2006-01-02 15:04:05"), e.message)
}

type ErrNegativeSqrt float64

func (err ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot sqrt negative number: %.2f", float64(err))
}

func MySqrt(f float64) (float64, error) {
	if f < 0 {
		return 0.0, ErrNegativeSqrt(f)
	}

	return math.Sqrt(f), nil
}

type MyReader struct {
	length int
	cursor int
}

func (r *MyReader) Read(b []byte) (int, error) {
	if r.cursor >= r.length {
		return 0, io.EOF
	}

	n := 0
	for ; n < len(b) && r.cursor < r.length; n, r.cursor = n+1, r.cursor+1 {
		b[n] = 'A'
	}

	return n, nil
}

type Rot13Reader struct {
	inner io.Reader
}

func (r *Rot13Reader) Read(b []byte) (int, error) {
	n, err := r.inner.Read(b)
	if err != nil {
		return 0, err
	}

	const a = byte('a')
	const z = byte('z')
	const A = byte('A')
	const Z = byte('Z')

	for i := 0; i < n; i++ {
		switch {
		case a <= b[i] && b[i] <= z:
			b[i] = a + ((b[i]-a)+13)%26
		case A <= b[i] && b[i] <= Z:
			b[i] = A + ((b[i]-A)+13)%26
		}
	}

	return n, nil
}

type MyImage struct{}

func (i MyImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (i MyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, 400, 400)
}

func (i MyImage) At(x int, y int) color.Color {
	const r = uint8(161)
	const g = uint8(221)
	const b = uint8(251)
	const a = uint8(255)

	const w int = 4

	if (x-y)%w == 0 {
		return color.RGBA{R: r, G: g, B: b, A: a}
	}

	return color.White
}

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

type List[T any] struct {
	next  *List[T]
	value T
}

func (l *List[T]) AddFirst(value T) {
	tmp := &List[T]{l.next, l.value}

	l.next = tmp
	l.value = value
}

func (l *List[T]) Sprintf(format string) string {
	s := ""

	i := l
	for i != nil {
		s += " " + fmt.Sprintf(format, i.value)

		i = i.next
	}

	return strings.TrimSpace(s)
}

func SumStream(a []int, ch chan int) {
	sum := 0

	for _, v := range a {
		sum += v
	}

	ch <- sum
}

func FibonacciCallback() func() int {
	x := 0
	y := 1

	return func() int {
		defer func() { x, y = y, x+y }()
		return x
	}
}

func FibonacciStream(ch, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)

	WalkRecursive(t, ch)
}

func WalkRecursive(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	WalkRecursive(t.Left, ch)
	ch <- t.Value
	WalkRecursive(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		x1, ok1 := <-ch1
		x2, ok2 := <-ch2

		if ok1 != ok2 || x1 != x2 {
			return false
		}

		if !ok1 || !ok2 {
			break
		}
	}

	return true
}

type SafeCounter struct {
	lock     sync.Mutex
	counters map[string]int
}

func (sc *SafeCounter) GetValue(key string) int {
	sc.lock.Lock()
	defer sc.lock.Unlock()

	return sc.counters[key]
}

func (sc *SafeCounter) Increment(key string) {
	sc.lock.Lock()
	defer sc.lock.Unlock()

	sc.counters[key]++
}

func Crawl(url string, depth int, fetcher Fetcher, fetchResults *FetchResults) {
	if depth < 0 || len(url) == 0 {
		return
	}

	fetchResults.lock.Lock()
	_, ok := fetchResults.results[url]
	if ok {
		return
	}
	fetchResults.lock.Unlock()

	res, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fetchResults.lock.Lock()
	fetchResults.results[url] = res
	fetchResults.lock.Unlock()

	fmt.Println(url)

	for _, child := range res.urls {
		go Crawl(child, depth-1, fetcher, fetchResults)
	}
}

type Fetcher interface {
	Fetch(url string) (res *FetchResult, err error)
}

type FetchResults struct {
	lock    sync.Mutex
	results map[string]*FetchResult
}

type FetchResult struct {
	body string
	urls []string
}

type FetcherMock struct{}

func (_ *FetcherMock) Fetch(url string) (res *FetchResult, err error) {
	switch url {
	case "https://golang.org/":
		return &FetchResult{
			"The Go Programming Language",
			[]string{
				"https://golang.org/pkg/",
				"https://golang.org/cmd/",
			},
		}, nil
	case "https://golang.org/pkg/":
		return &FetchResult{
			"Packages",
			[]string{
				"https://golang.org/",
				"https://golang.org/cmd/",
				"https://golang.org/pkg/fmt/",
				"https://golang.org/pkg/os/",
			},
		}, nil
	case "https://golang.org/pkg/fmt/":
		return &FetchResult{
			"Package fmt",
			[]string{
				"https://golang.org/",
				"https://golang.org/pkg/",
			},
		}, nil
	case "https://golang.org/pkg/os/":
		return &FetchResult{
			"Package os",
			[]string{
				"https://golang.org/",
				"https://golang.org/pkg/",
			},
		}, nil
	default:
		return &FetchResult{"", []string{}}, nil
	}
}

const ALPHANUM string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func NewRandomString() string {
	length := 64

	return NewRandomStringWithLen(length)
}

func NewRandomStringWithLen(length int) string {
	gen := NewRndGen()

	return NewRandomStringWithGen(length, gen)
}

func NewRandomStringWithGen(length int, gen *rand.Rand) string {
	// Go string are UTF-8 and ASCII is a subset of UTF-8,
	// ASCII is 7-bit character encoding system to represent
	// 128 unique characters.
	b := make([]byte, length)
	for i := range b {
		b[i] = ALPHANUM[gen.Intn(len(ALPHANUM))]
	}

	return string(b)
}

func NewRndGen() *rand.Rand {
	now := time.Now().UnixNano()
	src := rand.NewSource(now)
	gen := rand.New(src)

	return gen
}

func toFixed(f float64, p uint8) float64 {
	pow := math.Pow(10.0, float64(p))

	return math.Floor(f*pow) / pow
}

func DeepCopy(src, dst any) error {
	bytes, err := json.Marshal(src)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, dst)
}
