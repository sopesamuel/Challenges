package rotationalcipher
import "strings"
func RotationalCipher(plain string, shiftKey int) string {
	if plain == ""{
        return ""
    }
    if shiftKey == 0 || shiftKey == 26 {
        return plain
    }
  
    alpha2 := []rune{}
    alpha := "abcdefghijklmnopqrstuvwxyz"
     for _, v := range alpha{
		alpha2 = append(alpha2, v)
    }

    //rotated alpahbet
	final := string(alpha2[shiftKey:]) + string(alpha2[:shiftKey])
    secondF := []rune{}
     for _, v3 := range final{
         secondF = append(secondF, v3)
        }
    

    finalString := ""
    for _, v := range plain {
         if v >= 'A' && v <= 'Z'{
			for i, v1 := range alpha {
                if string(v) == strings.ToUpper(string(v1)){
                    finalString += strings.ToUpper(string(secondF[i]))
                    break
                }
            }
             } else if v >= 'a' && v <= 'z'{
            for i, v2 := range alpha {
                if string(v) == strings.ToLower(string(v2)){
                    finalString += strings.ToLower(string(secondF[i]))
                    break
                }
            }
            } else {
                 finalString += string(v)
            }
        }
    
    return finalString    
}
