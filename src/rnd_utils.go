package main

import (
	"math/rand"
	"time"
)

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
