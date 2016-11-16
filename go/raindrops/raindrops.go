package raindrops

import (
	"bytes"
	"strconv"
)

const testVersion = 2

//Convert accepts an int and does some ridiculous stuff to it.
func Convert(n int) string {
	var reply bytes.Buffer

	if n%3 == 0 {
		reply.WriteString("Pling")
	}

	if n%5 == 0 {
		reply.WriteString("Plang")
	}

	if n%7 == 0 {
		reply.WriteString("Plong")
	}

	if reply.Len() != 0 {
		return reply.String()
	}

	return strconv.Itoa(n)
}