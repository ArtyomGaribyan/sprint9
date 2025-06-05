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
	if size < 1 {
		return []int{}
	}
	rnd := rand.NewSource(time.Now().UnixNano())

	res := make([]int, size)
	for i := 0; i < size; i++ {
		res[i] = int(rnd.Int63())
	}
	return res
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, n := range data[1:] {
		if max < n {
			max = n
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) < CHUNKS {
		return maximum(data)
	}

	var wg sync.WaitGroup
	max := make([]int, CHUNKS+1)

	sizeChunk := len(data) / CHUNKS

	f := func(i int, d []int) {
		defer wg.Done()
		max[i] = maximum(d)
	}

	wg.Add(CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		d := data[i*sizeChunk : (i+1)*sizeChunk]
		go f(i, d)
	}

    	max[CHUNKS] = maximum(data[CHUNKS*sizeChunk:])

	wg.Wait()
	return maximum(max) 	
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
