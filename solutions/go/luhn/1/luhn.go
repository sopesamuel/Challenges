package luhn

func Valid(id string) bool {
     if len(id) <= 1 {
            return false
        }
	extractList := []int{}
    for _, v := range id {
        if (v >= '0' && v <= '9'){
        	extractList = append(extractList, int(v - '0'))
        } else if v == ' '{
            continue
        } else {
            return false
        }
    }

   
  	if len(extractList) <= 1  {
        	return false
    	}
    for i := len(extractList) - 2; i >= 0 ; i -= 2 {
       
        if (extractList[i] * 2 > 9) {
            extractList[i] = (extractList[i] * 2) - 9
        } else {
            extractList[i] = extractList[i] * 2
        }
        
    }

    sum := 0
    for _, v := range extractList {
        sum += v 
    }

    return sum % 10 == 0

}
