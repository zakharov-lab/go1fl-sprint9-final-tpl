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
			assert.Equal(t, false, containsBadNumber, "Слайс содержит значения выходящие из диапазона [1, %d]", Size)
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
		{name: "Максимум в середине", data: []int{44, 55, 22, 999, 10, 22, 11}, want: 999},
		{name: "Слайс из одинаковых элементов", data: []int{4, 4, 4, 4, 4}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			assert.Equal(t, tt.want, got, "Неправильно найден максимальный элемент. Найден %d, правильный %d", got, tt.want)
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil слайс", data: nil, want: 0},
		{name: "Пустой слайс", data: []int{}, want: 0},
		{name: "Один элемент", data: []int{32}, want: 32},
		{name: "Мало элементов", data: []int{55, 22, 13, 32, 23}, want: 55},
		{name: "Ровно 8 элементов, макимум в начале", data: []int{88, 77, 22, 33, 44, 55, 66, 11}, want: 88},
		{name: "Ровно 8 элементов, макимум в конце", data: []int{11, 77, 22, 33, 44, 55, 66, 88}, want: 88},
		{name: "Ровно 8 элементов, макимум в середине", data: []int{44, 77, 22, 33, 88, 55, 66, 11}, want: 88},
		{name: "9 элементов, максимум в конце", data: []int{88, 77, 22, 33, 44, 55, 66, 11, 99}, want: 99},
		{name: "16 элементов, максимум в середине", data: []int{193, 227, 351, 479, 523, 643, 727, 825, 991699, 100, 111, 182, 137, 14, 145, 129}, want: 991699},
		{name: "10 одинаковых элементов", data: []int{4, 4, 4, 4, 4, 4, 4, 4, 4, 4}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)
			assert.Equal(t, tt.want, got, "Неправильно найден максимальный элемент. Найден %d, правильный %d", got, tt.want)
		})
	}
}
