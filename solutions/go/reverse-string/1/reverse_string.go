package reverse

func Reverse(input string) string {
	b := []rune{}
    for _, v := range input {
        b = append(b, v)
    }
    reversed := ""
    for i := len(b) - 1; i >= 0; i-- {
        reversed += string(b[i])
    } 

    
    return reversed
}
