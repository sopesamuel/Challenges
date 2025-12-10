package thefarm
import 
("fmt" 
 "errors")


// TODO: define the 'DivideFood' function

func DivideFood(fc FodderCalculator, numberCows int) (float64, error){
    totalFodder, err := fc.FodderAmount(numberCows)
    if err != nil{
       
        return 0, err
    }
    multFact, err := fc.FatteningFactor()
    if err != nil{
         
        return 0, err
    }
    sumFodder := totalFodder/ float64(numberCows)
    amtFoodPerCow := sumFodder * multFact
    return amtFoodPerCow, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, numberCows int) (float64, error){
    if numberCows > 0{
      validateInput, err := DivideFood(fc, numberCows)
		if err != nil{
    		 return 0, err
        }
        return validateInput, nil
    }
    err := errors.New("invalid number of cows")
    return 0, err
}


// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct{
    numberCows int
    message string
}

func (inErr *InvalidCowsError) Error() string{
    return fmt.Sprintf("%d cows are invalid: %s", inErr.numberCows, inErr.message)
}

func ValidateNumberOfCows(numberCows int) error{
    switch{
        case numberCows < 0:
        return &InvalidCowsError{
             numberCows : numberCows,
             message : "there are no negative cows" ,
         }
        case numberCows == 0:
         return &InvalidCowsError{
             numberCows : numberCows,
             message : "no cows don't need food",
         }
    }
    return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
