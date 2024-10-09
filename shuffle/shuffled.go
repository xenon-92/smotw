package shuffle

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func ShuffleNames(names []string) {
	length := len(names)

	for i := length - 1; i > 0; i-- {
		// Generate a secure random number between 0 and i (inclusive)
		randIndex, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			fmt.Println("Error generating random number:", err)
			return
		}

		// Convert the result to an integer
		j := int(randIndex.Int64())

		// Swap the elements
		names[i], names[j] = names[j], names[i]
	}
}
