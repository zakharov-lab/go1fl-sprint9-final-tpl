package main

import (
	"fmt"
	"math/rand"
)

const (
	Size       = 100_000_000
	MaxRandNum = 1_000_000_000
	Chunks     = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}
	generated := make([]int, size)
	for i := range generated {
		generated[i] = rand.Intn(MaxRandNum) + 1
	}
	return generated
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	return 0
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	return 0
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", Size)
	// ваш код здесь

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь

	//fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", Chunks)
	// ваш код здесь

	//fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
