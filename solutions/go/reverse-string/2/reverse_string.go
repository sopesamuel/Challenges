package reverse

func Reverse(input string) string {
    rune1 := []rune(input)
	b := []rune{}
    
    for i := len(rune1) - 1; i >= 0; i-- {
      	b = append(b, rune1[i])
    } 

    
    return string(b)
}
