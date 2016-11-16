package gigasecond

import "time"

const testVersion = 4 // find the value in gigasecond_test.go
const gigasecond = time.Second * 1000000000

//AddGigasecond takes a date and returns that date plus one gigasecond
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}