package prime

import (
	"os"
)

// GetNextPrime gets the next prime after the given number
func GetNextPrime(cur int64) int64 {
	if _, err := os.Stat("/tmp/stop"); err == nil {
		print("Not replying as stopping soon")
		return 0
	} else if os.IsNotExist(err) {
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
	} else {
		panic(err)
	}
	return -1
}
