package main

import (
	"math/rand"
	_ "net/http/pprof"
	"sync"
)

const (
	rows    = 4
	columns = 10000
)

var matrix = [][]float64{}

func prepare() {
	for i := 0; i < rows; i++ {
		matrix = append(matrix, []float64{})
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			matrix[i] = append(matrix[i], rand.Float64())
		}
	}
}

func main() {
	prepare()
	fourThreads()
}

func fourThreads() {
	outputChan := make(chan float64, rows+10)
	wg := sync.WaitGroup{}
	wg.Add(rows)

	for i := 0; i < rows; i++ {
		go func(index int, wg *sync.WaitGroup, outputChan chan<- float64) {
			temp := 0.0
			for j := 0; j < columns; j++ {
				temp += matrix[index][j]
			}
			outputChan <- temp
			wg.Done()
		}(i, &wg, outputChan)
	}
	wg.Wait()
	close(outputChan)

	sum := 0.0
	for out := range outputChan {
		sum += out
	}
}

func singleThread() {
	sum := 0.0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			sum += matrix[i][j]
		}
	}
}
