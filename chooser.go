package rltao

import (
	"crypto/rand"
	"math/big"
)

func GenerateId() int {
	base10_sum := 0 // This will be the base10 sum of the monograms

	// Simple random
	// choose random from [0,81)
	// randBigInt, err := rand.Int(rand.Reader, big.NewInt(81))
	// if err != nil {
	// 	panic(err)
	// }

	// More authentic random to build the tetragram (tetra is four)
	// Choose a random int 0-2 for each "gram" of the tetragram, then sum it to the previous with a weighting
	// To convert from trinary/ternary (base 3) to base 10, multiply each part by it's order of magnitude
	// i.e. The first or bottom line gets a weight of 27, the second, 9, then 3, then 1
	// TODO: Am I doing this backwards? Would it be more authentic to give the last rand trinary the heaviest weight?
	for i := 0; i <= 3; i++ {
		// choose random up from [0,3)
		randBigInt, err := rand.Int(rand.Reader, big.NewInt(3))
		if err != nil {
			panic(err)
		}
		randInt := int(randBigInt.Int64())

		switch i {
		case 0:
			base10_sum = base10_sum + randInt*27
		case 1:
			base10_sum = base10_sum + randInt*9
		case 2:
			base10_sum = base10_sum + randInt*3
		case 3:
			base10_sum = base10_sum + randInt*1
		}
	}

	return base10_sum + 1 // Range of the passages is from 1 to 81
}
