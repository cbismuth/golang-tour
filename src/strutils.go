package main

import (
	"math/rand"
	"time"
)

const LETTERS string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func NewRandomString() string {
	length := 64

	return NewRandomStringWithLen(length)
}

func NewRandomStringWithLen(length int) string {
	gen := rand.New(rand.NewSource(time.Now().UnixNano()))

	return NewRandomStringWithGen(length, gen)
}

func NewRandomStringWithGen(length int, gen *rand.Rand) string {
	// Go string are UTF-8 and ASCII is a subset of UTF-8,
	// ASCII is 7-bit character encoding system to represent
	// 128 unique characters.
	b := make([]byte, length)
	for i := range b {
		b[i] = LETTERS[gen.Intn(len(LETTERS))]
	}
	return string(b)
}

func SwapStringTuple(s1, s2 string) (string, string) {
	return s2, s1
}
