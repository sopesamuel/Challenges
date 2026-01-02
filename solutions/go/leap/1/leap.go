// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap
// 2100 false
// 1996 true
// 2000 2400 true
// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
    
	if year % 4 == 0 {
        if year % 400 == 0 {
            return true
        } else if year % 100 == 0 {
            return false
        }
        return true
    }
    return false
    
}
