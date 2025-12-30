package raindrops
import "fmt"

func Convert(number int) string {
    msg := ""
  		if  number % 3 == 0 {
            msg += "Pling"
        }
       
        if number % 5 == 0 {
            msg += "Plang"
        }
    
        if number % 7 == 0 {
           msg += "Plong"
        }
    	if msg == ""{
            return fmt.Sprintf("%d", number)
        }
       return msg
}
