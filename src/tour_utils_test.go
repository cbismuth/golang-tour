package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Floor(t *testing.T) {
	assert.Equal(t, 1.23, toFixed(1.234, 2))
}

func Test_DeeCopy(t *testing.T) {
	gen := NewRndGen()
	x := gen.Float64()
	y := gen.Float64()

	v1 := Vertex{x, y}

	var v2 Vertex
	err := DeepCopy(&v1, &v2)

	assert.Nil(t, err)
	assert.False(t, v1 != v2)
	assert.Equal(t, v1, v2)
	assert.Equal(t, x, v1.X)
	assert.Equal(t, y, v1.Y)
	assert.Equal(t, x, v1.X)
	assert.Equal(t, y, v1.Y)
}
