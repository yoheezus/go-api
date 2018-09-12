package prime

// GetNextPrime gets the next prime after the given number
func GetNextPrime(cur int64) int64 {
	next := cur + 2
	if cur == 2 {
		next = cur + 1
	}
	trynum := int64(3)
	for trynum < next {
		if next%trynum == 0 {
			// Next is not a prime
			next++
			trynum = int64(3)
			continue
		}
		trynum++
	}
	return next
}
