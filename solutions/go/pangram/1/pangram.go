package pangram
import "unicode"
func IsPangram(input string) bool {
    filt := map[rune]int{}
	for i, v := range input {
        filt[unicode.ToLower(v)] = i
    }
	alpha := "abcdefghijklmnopqrstuvwxqyz"
    for _, v := range alpha {
        _, ok := filt[v]
        if !ok{
            return false
        }
    }
    return true
}
