// Leap stub file

// The package name is expected by the test program.
package leap

// testVersion should match the targetTestVersion in the test file.
const testVersion = 2

//IsLeapYear returns true if a given year is a leap year.
func IsLeapYear(year int) bool {
	if year%100 == 0 {
		if year%400 == 0 {
			return true
		}
		return false
	}

	if year%4 == 0 {
		return true
	}

	return false
}
