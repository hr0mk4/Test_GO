package main

import (
	"fmt"
	"time"
)

func TimeDifference(start, end time.Time) time.Duration {
	return end.Sub(start)
}

func main() {
	start := time.Date(2023, 10, 23, 2, 41, 49, 0, time.UTC)
	end := time.Date(2023, 10, 23, 4, 41, 49, 0, time.UTC)
	diff := TimeDifference(start, end)
	fmt.Println(diff) // Output: 2h0m0s
	t := time.Now()
	fmt.Println(t.Weekday())
	time.
}
