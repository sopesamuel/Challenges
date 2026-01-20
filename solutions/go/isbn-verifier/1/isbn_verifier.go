package isbn

func IsValidISBN(isbn string) bool {

    if isbn == ""{
        return false
    }

    pre := 10
    sum := 0
    
    for _, v := range isbn {

        switch {
            case (v >= '0' && v <= '9'):
            	sum += int(v - '0')  * pre
            	pre--
            case v == '-':
            	continue
            case (v == 'x' || v == 'X') && pre == 1 :
            	sum += 10
            	pre--
            default:
            	return false
        }
    }

    if pre != 0 {
        return false
    }
    return sum % 11 == 0
}
