package isogram
import "unicode"

func IsIsogram(word string) bool {
    store := map[rune]int{}
	for i, v := range word {
        
        if v == ' ' || v == '-'{
          continue
        }
    
         _ , ok := store[unicode.ToLower(v)] 
        if ok {
    		return false
        } else {
            store[unicode.ToLower(v)] = i 
        }  
    }
    return true
}
