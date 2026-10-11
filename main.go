package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size <= 0 {
		return nil
	}

	sl := make([]int, 0, size)

	for i := 0; i < size; i++ {
		sl = append(sl, rand.Int())
	}
	return sl
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]
	for _, v := range data {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	var wg sync.WaitGroup

	maxSl := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS
	remain := len(data) % CHUNKS

	start := 0
	for i := 0; i < CHUNKS; i++ {
		end := start + chunkSize
		if i == CHUNKS-1 {
			end += remain
		}
		chunk := data[start:end]
		start = end

		wg.Add(1)
		go func(chunk []int) {
			defer wg.Done()
			maxSl[i] = maximum(chunk)
		}(chunk)
	}
	wg.Wait()

	return maximum(maxSl)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start1 := time.Now()
	max := maximum(data)
	elapsed := time.Since(start1).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start2 := time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start2).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
