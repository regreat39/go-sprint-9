package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Пишите тесты в этом файле
func TestGenerateWithSizeZero(t *testing.T) {
	slice := generateRandomElements(0)
	require.Nil(t, slice)
}

// func TestGenerateWith

func TestGenerateWithSizeOne(t *testing.T) {
	slice := generateRandomElements(1)
	require.Len(t, slice, 1)
}

func TestMaximumSliceNil(t *testing.T) {
	var slice []int
	max := maximum(slice)
	require.Equal(t, 0, max)
}

func TestMaximumSliceEmpty(t *testing.T) {
	slice := []int{}
	max := maximum(slice)
	require.Equal(t, 0, max)
}

func TestMaximumSliceOneElement(t *testing.T) {
	slice := []int{1}
	max := maximum(slice)
	require.Equal(t, 1, max)
}

func TestMaxChunkSliceNil(t *testing.T) {
	var slice []int
	max := maxChunks(slice)
	require.Equal(t, 0, max)
}
