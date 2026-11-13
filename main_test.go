package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantLen int
	}{
		{name: "Отрицательный размер", size: -100, wantLen: 0},
		{name: "Нулевой размер", size: 0, wantLen: 0},
		{name: "Положительный размер", size: 100, wantLen: 100},
		{name: "Большой размер", size: 10000, wantLen: 10000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			assert.Equal(t, tt.wantLen, len(got), "Размер слайса не соответствует ожидаемому")
			containsBadNumber := false
			for _, value := range got {
				if value <= 0 || value > MaxRandNum {
					containsBadNumber = true
					break
				}
			}
			assert.Equal(t, containsBadNumber, false, "Слайс содержит значения меньше или равные 0 или больше максимально допустимого")
		})
	}
}
