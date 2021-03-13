package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.EqualValues(t, 4, Sum(2, 2))
}

func TestProduct(t *testing.T) {
	assert.EqualValues(t, 20, Product(4, 5))
}
