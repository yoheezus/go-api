package wait

import (
	"fmt"
	"math/rand"
	"time"
)

// WaitDurationFixed waits for a set amount of seconds
func WaitDurationFixed(cur int64) int64 {
	time.Sleep(time.Duration(cur) * time.Second)
	return time.Now().Unix()
}

// WaitDurationRandom waits for a time between 0 and 30s
func WaitDurationRandom() int64 {
	length := fmt.Sprintf("%vms", rand.Intn(30000))
	dur, _ := time.ParseDuration(length)
	time.Sleep(dur)
	return time.Now().Unix()
}
