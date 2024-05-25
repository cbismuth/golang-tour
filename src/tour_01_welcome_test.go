package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_Welcome1(t *testing.T) {
	fmt.Println("Hello, 世界!")
}

func Test_Welcome4(t *testing.T) {
	now := time.Now()
	format := time.DateTime

	fmt.Println("The time is", now.Format(format))
}
