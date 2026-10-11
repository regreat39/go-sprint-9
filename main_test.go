package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Пишите тесты в этом файле
func TestGenerateWithSizeZero(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantLen int
		wantNil bool
	}{
		{"отрицательный размер слайса", -1, 0, true},
		{"нулевой размер слайса", 0, 0, true},
		{"размер слайса равен 1", 1, 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sl := generateRandomElements(tt.size)
			require.Len(t, sl, tt.wantLen)
			if tt.wantNil {
				require.Nil(t, sl)
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"пустой слайс", []int{}, 0},
		{"непустой слайс с нулем", []int{0}, 0},
		{"слайс с одним элементом", []int{1}, 1},
		{"одинаковые элементы", []int{7, 7, 7}, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maxValue := maximum(tt.data)
			require.Equal(t, tt.want, maxValue)

		})
	}
}
