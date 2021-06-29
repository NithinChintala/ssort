package ssort

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"sync"
)

const (
	DefaultProcs = 4
)

// Float64s sorts a slice of float64s in increasing order.
// Not-a-number (NaN) values are ordered before other values.
func Float64s(data []float64) { Float64sProcs(data, DefaultProcs) }

// Float64s sorts a slice of float64s in increasing order spawning
// numProcs goroutines.
// Not-a-number (NaN) values are ordered before other values.
func Float64sProcs(data []float64, numProcs int) {
	sizes := make([]int, numProcs)
	samps := sample(data, numProcs)

	var sortWg, doneWg sync.WaitGroup
	sortWg.Add(numProcs)
	doneWg.Add(numProcs)

	for i := 0; i < numProcs; i++ {
		go ssortWorker(data, samps, sizes, i, &sortWg, &doneWg)
	}
	doneWg.Wait()
}

// sample takes samples from the input data to create numProcs+1 buckets.
// The sampling process is as follows.
//   1. Randomly select 3*(numProcs-1) items
//   2. Sort the randomly select items
//   3. Take the median of each group of three items in the sorted selection
//   4. Add 0 to the front and MaxFloat64 to the end to get numProcs+1 items
// The result should be numProcs buckets with their ranges in ascending order.
func sample(data []float64, numProcs int) []float64 {
	size := len(data)
	sampleSize := 3 * (numProcs - 1)
	var idx int

	samples := make([]float64, sampleSize)
	for i := 0; i < sampleSize; i++ {
		idx = rand.Intn(size)
		samples[i] = data[idx]
	}
	sort.Float64s(samples)

	numGroups := sampleSize / 3
	buckets := make([]float64, 0)

	buckets = append(buckets, 0)
	for i := 0; i < numGroups; i++ {
		idx = 3*i + 1
		buckets = append(buckets, samples[idx])
	}
	buckets = append(buckets, math.MaxFloat64)

	return buckets
}

// ssortWorker TODO
func ssortWorker(data, samps []float64, sizes []int, procNum int, sortWg, doneWg *sync.WaitGroup) {
	minVal := samps[procNum]
	maxVal := samps[procNum+1]
	binned := make([]float64, 0)
	for i := 0; i < len(data); i++ {
		if procNum == 0 && data[i] == minVal {
			binned = append(binned, data[i])
		} else if data[i] > minVal && data[i] <= maxVal {
			binned = append(binned, data[i])
		}
	}

	sort.Float64s(binned)
	sizes[procNum] = len(binned)

	var start int
	sortWg.Done()
	sortWg.Wait()
	for i := 0; i < procNum; i++ {
		start += sizes[i]
	}
	
	for i, val := range binned {
		data[start+i] = val
	}

	doneWg.Done()
}

func main() {
	test := randSlice(10)
	fmt.Println(test)
	Float64s(test)
	fmt.Println(test)
}

