package scrabble
import "strings"
func Score(word string) int {
    count := 0
    
	for _, v := range word {
        v := strings.ToUpper(string(v))
        switch {
            case v == "D" || v == "G" :
            	count += 2
            case v == "B" || v == "C" || v == "M" || v == "P" :
            	count += 3
            case v == "F" || v == "H" || v == "V" || v == "W" || v == "Y" :
            	count += 4
            case v == "K" :
            	count += 5
            case v == "J" || v == "X" :
            	count += 8
            case v == "Q" || v == "Z" :
            	count += 10
            default :
            	count += 1
        }
    }
    return count
}
