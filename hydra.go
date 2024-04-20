package main 

import (
	"fmt"
	"math/big"
)

// NegativeHeadsError detects whether a number of heads os negative.
// Protects against underflow
type NegativeHeadsError struct{}

func (*NegativeHeadsError) Error() string {
	// Guards against integer underflow
	return "Cannot have negative heads"
}
func main() {
	treeSize := 0

	for treeSize < 6 {
		heads := createBeginningTree(treeSize)
		fmt.Printf("Beginning Hydra Heads shape: %s\n", heads)
		fmt.Printf("Final number of rounds taken: %s\n\n", processTree(heads))
		treeSize++ 
	}
}

func createBeginningTree(treeSize int) []*big.Int {
	heads := make([]*big.Int, 0)
	for i := 0; i <= treeSize; i++ {
		if i != treeSize {
			heads = append(heads, big.NewInt(0))
		} else {
			heads = append(heads, big.NewInt(1))
		}
	}
	return heads
}

func processTree (heads []*big.Int) (*big.Int) {
	count := big.NewInt(1)
	heads, tempCount := removeLowest(heads, count)
	// If the count is the same, nothing has changed, so the heads list is
	// empty...
	for count.Cmp(tempCount) != 0 {
		// This is just a while loop, as Go doesn't have a separate keyword

		// If you want to see the "moves"
		fmt.Printf("Hydra head shape is: %s, at round: %s.\n", heads, tempCount)
		count.Add(big.NewInt(0), tempCount)
		heads, tempCount = removeLowest(heads, count)
	}

	// As it finds the current step, the number required to kill the hydra
	// is one less...
	count.Sub(count, big.NewInt(1))
	return count
}
func removeLowest (heads []*big.Int, count *big.Int) ([]*big.Int, *big.Int) {
	// Avoiding pointer collision
	newCount := big.NewInt(0)
	newCount.Add(newCount, count)

	// ONLY USE THESE FOR COMPARISON / INCREMENT
	// DO NOT ASSIGN TO LISTS	
	bigOne := big.NewInt(1)
	bigZero := big.NewInt(0)

	for i, val := range heads {
		if val.Cmp(bigZero) == -1 {
			panic(&NegativeHeadsError{})
		}

		if val.Cmp(bigZero) == 0 {
			// No heads, move up the tree
			continue
		}

		if i == 0 {
			// Heads that don't generate others can be killed in
			// one step. Killing them all at once reduces load times
			newCount.Add(newCount, heads[0])

			// Dont set it here as it is a pointer and so could change
			// big Zero
			heads[0] = big.NewInt(0)
			return heads, newCount
		} 

		headsAbove := 0
		for j, val := range (heads) {
			if j <= i {continue}
			if val.Cmp(bigZero) != 0 {
				headsAbove = 1
				break
			}
		}

		if i == 1 {
			// Generic reduction formula for N heads at connected node 1
			newCount = simplifyHeightOne(val, count)
			heads[i] = big.NewInt(0)
			if headsAbove == 0 {
				// Because we destroy the last node that remains
				newCount = big.NewInt(0).Add(newCount, big.NewInt(1))
			}
			return heads, newCount
		}

		// We have found a head so kill it
		heads[i].Sub(heads[i], bigOne)

		// Increase the number of heads below by the step count
		heads[i-1].Add(heads[i-1], count)

		// Increase the count
		newCount = big.NewInt(0).Add(newCount, bigOne)
		
		if heads[i].Cmp(bigZero) == 0 && headsAbove == 0{
			// Add the new stub created to the list of heads
			heads[i-1].Add(heads[i-1], bigOne)
		}
		return heads, newCount
	}

	// If there are no heads left, return the input
	return heads, newCount
}

func simplifyHeightOne(heads *big.Int, count *big.Int) (*big.Int) {
	newCount := big.NewInt(0).Add(big.NewInt(0), count)

	bigOne := big.NewInt(1)	
	bigTwo := big.NewInt(2)

	twoExp := big.NewInt(0).Exp(bigTwo, heads, nil)
	bracket := big.NewInt(0).Add(newCount, bigOne)

	result := big.NewInt(0).Mul(twoExp, bracket)
	
	return big.NewInt(0).Sub(result, bigOne)
}
