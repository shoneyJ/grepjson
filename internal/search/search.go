package search

import (
	"fmt"
	"runtime"
	"sync"
)

type Semaphore chan struct{}

func (s Semaphore) Acquire() { s <- struct{}{} }
func (s Semaphore) Release() { <-s }

func searchRecursive(data interface{}, keyPattern string, currentPath string, results *[]MatchResult, maxDistance int, mu *sync.Mutex, wg *sync.WaitGroup, sem Semaphore) {

	defer wg.Done()

	switch v := data.(type) {
	case map[string]interface{}:
		for k, val := range v {
			newPath := buildPath(currentPath, k)

			if fuzzyMatch(k, keyPattern, maxDistance) {
				mu.Lock()
				*results = append(*results, MatchResult{
					newPath,
					val,
				})
				mu.Unlock()
			}
			wg.Add(1)
			sem.Acquire()
			go func(val interface{}, path string) {
				defer sem.Release()
				searchRecursive(val, keyPattern, path, results, maxDistance, mu, wg, sem)
			}(val, newPath)
		}
	case []interface{}:
		for i, val := range v {
			newPath := buildPath(currentPath, fmt.Sprintf("[%d]", i))
			wg.Add(1)
			sem.Acquire()
			go func(val interface{}, path string) {
				defer sem.Release()
				searchRecursive(val, keyPattern, path, results, maxDistance, mu, wg, sem)
			}(val, newPath)
		}

	}
}

func searchConcurrent(data interface{}, keyPattern string, maxDistance int) []MatchResult {

	var results []MatchResult

	var mu sync.Mutex
	var wg sync.WaitGroup

	numCores := runtime.NumCPU()

	if numCores < 1 {
		numCores = 1
	}

	runtime.GOMAXPROCS(numCores)

	sem := make(Semaphore, numCores)

	wg.Add(1)

	go func() {

		searchRecursive(data, keyPattern, "", &results, maxDistance, &mu, &wg, sem)
	}()

	wg.Wait()

	return results

}
