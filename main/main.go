package main

import (
	"fmt"
)

func RemovNb(m uint64) [][2]uint64 {
	var ans = make([][2]uint64, 0)

	// count all the numbers in sequence
	var count = (m * (m + 1)) / 2

	// check the pairs
	// count - i - j = ij -> (count - i) = ij + j -> (count - i) = j(i + 1) -> j = ( count - i ) / ( i + 1 )
	for i := uint64(1); i <= m; i++ {
		j := (count - i) / (i + 1)

		if (count-i)%(i+1) == 0 && j <= m {
			ans = append(ans, [2]uint64{i, j})
		}
	}

	if len(ans) == 0 {
		return nil
	}

	return ans
}

func main() {
	fmt.Println(RemovNb(26))
}
