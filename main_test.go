package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			assert.Equal(t, containsBadNumber, false, "Слайс содержит значения выходящие из диапазона [1, 1_000_000_000]")
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil слайс", data: nil, want: 0},
		{name: "Пустой слайс", data: []int{}, want: 0},
		{name: "Слайс из одного элемента", data: []int{42}, want: 42},
		{name: "Максимум в начале", data: []int{100, 53, 11, 22, 66}, want: 100},
		{name: "Максимум в конце", data: []int{44, 33, 22, 55, 11, 123}, want: 123},
		{name: "Макмимум в середине", data: []int{44, 55, 22, 999, 10, 22, 11}, want: 999},
		{name: "Слайс из одинаковых элементов", data: []int{4, 4, 4, 4, 4}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			assert.Equal(t, tt.want, got, "Неправильно найден максимальный элемент. Найден %d, правильный %d", got, tt.want)
		})
	}
}
