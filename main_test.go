package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func randomSize() int {
	rnd := rand.NewSource(time.Now().UnixNano())
	return rand.New(rnd).Intn(1000)
}

func TestGenerateRandomElements(t *testing.T) {
	size := randomSize()

	res := generateRandomElements(size)
	require.Len(t, res, size)

	testValues := []int{9, 5, 0, 1}

	for _, size := range testValues {
		assert.Len(t, generateRandomElements(size), size)
	}
}

func TestMaximum(t *testing.T) {
	testValues := map[int][]int{
		9:  {3, 5, 1, 8, 9},
		13: {4, 13, 6, 9, 10, 8},
		0:  {},
		3:  {3},
	}

	for expected, slice := range testValues {
		assert.Equal(t, expected, maximum(slice))
	}
}

func TestMaxCHUNKS(t *testing.T) {
	testValues := map[int][]int{
		33:  {3, 5, 1, 8, 9, 3, 5, 1, 3, 5, 1, 3, 5, 1, 4, 3, 4, 6, 7, 23, 4, 5, 7, 1, 33, 4, 6},
		357: {4, 13, 6, 9, 10, 78, 23, 22, 5, 77, 1, 5, 6, 88, 2, 357},
		0:   {},
		3:   {3},
		42:  {11, 42, 5, 1, 6},
	}

	for expected, slice := range testValues {
		assert.Equal(t, expected, maxChunks(slice))
	}

}
