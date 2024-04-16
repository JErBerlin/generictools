package channel

import (
	"runtime"
	"sort"
	"sync"
)

// Map applies the function f to each element received from the input channel in and
// sends the result to the output channel out.
// Closing the out channel is responsability of the caller.
func Map[T any](in <-chan T, out chan<- T, f func(T) T) {
	for e := range in {
		out <- f(e)
	}
}

type indexedInput[T any] struct {
	Index int
	Value T
}

type indexedOutput[T any] struct {
	Index int
	Value T
}

// MapParallel applies the function f to each element received from the input channel in and
// sends the result to the output channel out. It uses workers to process the elements concurrently.
// Closing the out channel is responsability of the caller.
func MapParallel[T any](in <-chan T, out chan<- T, f func(T) T) {
	numWorkers := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Create a channel for indexed inputs and outputs.
	// We can tune the lenght of the buffered out channel.
	indexedIn := make(chan indexedInput[T])
	indexedOut := make(chan indexedOutput[T], 100) // Buffered channel for outputs. Size of 100 is arbitrary.

	// Distribute work among workers.
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			for input := range indexedIn {
				indexedOut <- indexedOutput[T]{Index: input.Index, Value: f(input.Value)}
			}
		}()
	}

	// Send inputs with indices.
	go func() {
		idx := 0
		for val := range in {
			indexedIn <- indexedInput[T]{Index: idx, Value: val}
			idx++
		}
		close(indexedIn)
	}()

	// Collect outputs and ensure order.
	go func() {
		wg.Wait()
		close(indexedOut)
	}()

	// Reorder results based on indices.
	var outputs []indexedOutput[T]
	for outVal := range indexedOut {
		outputs = append(outputs, outVal)
	}

	// Sort the outputs slice based on the index.
	sort.Slice(outputs, func(i, j int) bool {
		return outputs[i].Index < outputs[j].Index
	})

	// Send the sorted outputs to the out channel.
	for _, output := range outputs {
		out <- output.Value
	}

	// Closing the out channel is responsability of the caller.
}
