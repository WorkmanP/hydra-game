package main 

import (
	"fmt"
)

type NegativeHeadsError struct{}

func (*NegativeHeadsError) Error() string {
	return "Cannot have negative heads"
}
func main() {
	treeSize := 0
	for treeSize < 4 {
		heads, stubs := createBeginningTree(treeSize)
		fmt.Println(heads, stubs)
		fmt.Println(processTree(heads, stubs))
		treeSize ++ 
	}
}

func createBeginningTree(treeSize int) ([]int, []int) {
	heads := make([]int, 0)
	stubs := make([]int, 0)
	for i := 0; i <= treeSize; i++ {
		if i != treeSize {
			heads = append(heads, 0)
			stubs = append(stubs, 1)
		} else {
			heads = append(heads, 1)
			stubs = append(stubs, 0)
		}
	}
	return heads, stubs
}

func processTree (heads, stubs []int) (int) {
	count := 1
	heads, stubs, tempCount := removeLowest(heads, stubs, count)

	// If the count is the same, nothing has changed, so the heads list is
	// empty...

	// This is just a while loop as go doesn't have a separate keyword
	for count != tempCount {
		count = tempCount
		heads, stubs, tempCount = removeLowest(heads, stubs, count)
	}

	// As it return the current step, the number required to kill the hydra
	// is one less...
	return count-1
}
func removeLowest (heads, stubs []int, count int) ([]int, []int, int) {
	for i, val := range heads {
		if i == 0 && val < 0 {
			panic(&NegativeHeadsError{})
		}
		if i == 0 && val != 0 {
			count += heads[0]
			heads[0] = 0
			return heads, stubs, count
		}

		if val == 0 {
			continue
		}
		heads[i]--
		heads[i-1] += count
		count++

		headsAbove := 0

		for j, val := range (heads) {
			if j <= i {continue}
			headsAbove += val
		}
		if val == 1 && headsAbove == 0{
			// Add the new stub create to the list of heads
			heads[i-1]++
			stubs[i-1]=0
		}
		fmt.Println(heads, count)
		return heads, stubs, count
	}
	// If there are no heads left, return the input
	return heads, stubs, count
}
