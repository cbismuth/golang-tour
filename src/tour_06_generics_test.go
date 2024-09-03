package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Generics1(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}

	assert.Equal(t, -1, Index(a, 0))

	for i := 0; i < 8; i++ {
		assert.Equal(t, i, Index(a, i+1))
	}
}

func Test_Generics2(t *testing.T) {
	l := &List[int]{nil, 0}

	for i := 1; i < 10; i++ {
		l.AddFirst(i)
	}

	assert.Equal(t, "9 8 7 6 5 4 3 2 1 0", l.Sprintf("%d"))
}
