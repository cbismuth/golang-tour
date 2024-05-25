package main

import (
	"fmt"
	"math"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const n int = 10

func Test_FlowControl1(t *testing.T) {
	sum := 0
	for i := 0; i < n; i++ {
		sum++
	}

	assert.Equal(t, n, sum)
}

func Test_FlowControl2(t *testing.T) {
	sum := 0
	for sum < n {
		sum++
	}

	assert.Equal(t, n, sum)
}

func Test_FlowControl3(t *testing.T) {
	// Same as Test_FlowControl2 (extra semicolons removed)

	assert.True(t, true)
}

func Test_FlowControl4(t *testing.T) {
	sum := 0

	for {
		if sum == 10 {
			break
		}

		sum++
	}

	assert.Equal(t, n, sum)
}

func Test_FlowControl5(t *testing.T) {
	gen := NewRndGen()

	for i := 0; i < n; i++ {
		var fact float64
		if i%2 == 0 {
			fact = 1.0
		} else {
			fact = -1.0
		}

		rnd := gen.Float64() * fact
		sqrt := math.Sqrt(rnd)

		if rnd < 0 {
			assert.True(t, math.IsNaN(sqrt))
		} else {
			assert.Positive(t, sqrt)
		}
	}
}

func Test_FlowControl6(t *testing.T) {
	for i := 0; i < n; i++ {
		if pow := math.Pow(float64(i), float64(i)); pow > float64(n) {
			assert.Greater(t, pow, float64(n))
		}
	}
}

func Test_FlowControl7(t *testing.T) {
	for i := 0; i < n; i++ {
		if pow := math.Pow(float64(i), float64(i)); pow > float64(n) {
			assert.Greater(t, pow, float64(n))
		} else {
			assert.LessOrEqual(t, pow, float64(n))
		}
	}
}

func Test_FlowControl8(t *testing.T) {
	const x float64 = 16.123456789

	assert.Equal(t, math.Sqrt(x), SqrtFlowControl8(x))
}

func SqrtFlowControl8(x float64) float64 {
	z := 1.0

	for i := 0; i < 10000; i++ {
		z -= (z*z - x) / (2 * z)
	}

	fmt.Printf("The square root of [%f] is [%f]\n", x, z)

	return z
}

func Test_FlowControl9(t *testing.T) {
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux")
	case "darwin":
		fmt.Println("macOS")
	default:
		fmt.Println("Other")

	}
}

func Test_FlowControl10(t *testing.T) {
	switch weekday := time.Now().Weekday(); time.Saturday {
	case weekday + 0:
		fmt.Println("Weekend is now!")
	case weekday + 1:
		fmt.Println("Weekend is tomorrow!")
	default:
		fmt.Println("Weekend is too far away ...")
	}
}

func Test_FlowControl11(t *testing.T) {
	hour := time.Now().Hour()

	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening ...")
	}
}

func Test_FlowControl12(t *testing.T) {
	defer fmt.Printf(" ... World!\n")

	fmt.Printf("Hello")
}

func Test_FlowControl13(t *testing.T) {
	defer fmt.Printf("d\n")
	defer fmt.Printf("c")
	defer fmt.Printf("b")

	fmt.Printf("a")
}
