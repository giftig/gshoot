package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbsPositive(t *testing.T) {
	expected := float32(10.12)
	actual := Abs(10.12)

	assert.Equal(t, actual, expected)
}

func TestAbsZero(t *testing.T) {
	expected := float32(0.0)
	actual := Abs(0.0)

	assert.Equal(t, actual, expected)
}

func TestAbsNegative(t *testing.T) {
	expected := float32(7.777)
	actual := Abs(-7.777)

	assert.Equal(t, actual, expected)
}

func TestMinLeft(t *testing.T) {
	assert.Equal(t, Min(-10.0, 10.0), float32(-10.0))
}

func TestMinEqual(t *testing.T) {
	assert.Equal(t, Min(0.0, 0.0), float32(0.0))
}

func TestMinRight(t *testing.T) {
	assert.Equal(t, Min(20.0, -20.0), float32(-20.0))
}
