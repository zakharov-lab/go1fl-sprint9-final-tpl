package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
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
	if len(data) == 0 {
		return 0
	}
	maxValue := data[0]
	for _, value := range data {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	chunkSize := len(data) / Chunks
	if chunkSize == 0 {
		return maximum(data)
	}

	maxValues := make([]int, Chunks)
	var wg sync.WaitGroup

	for i := 0; i < Chunks; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == Chunks-1 {
			end = len(data)
		}
		wg.Add(1)
		go func(idx int, partition []int) {
			defer wg.Done()
			maxValues[idx] = maximum(partition)
		}(i, data[start:end])
	}
	wg.Wait()
	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", Size)
	data := generateRandomElements(Size)

	fmt.Println("Ищем максимальное значение в один поток")
	startTime := time.Now()
	oneThreadMaximum := maximum(data)
	elapsed := time.Since(startTime).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d mks\n", oneThreadMaximum, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", Chunks)
	startTime = time.Now()
	multiThreadMaximum := maximum(data)
	elapsed = time.Since(startTime).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d mks\n", multiThreadMaximum, elapsed)
}
