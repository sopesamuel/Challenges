package sorting
import (
    "fmt"
    "strconv"
)
// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
    
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	str, ok := fnb.(FancyNumber)
    if !ok {
        return 0
    }
    value, _ := strconv.Atoi(str.Value())
    return value
    
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	result := ExtractFancyNumber(fnb)
    if result == 0 {
        return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(result))
    }
    return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(result))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
	switch v := i.(type){
        case int:
        	descNum := DescribeNumber(float64(v))
        	return descNum 
        case float64:
        	descNum := DescribeNumber(v)
        	return descNum 
        case NumberBox:
        	descNumBox := DescribeNumberBox(v)
        	return descNumBox
        case FancyNumberBox:
        	fancyBox := DescribeFancyNumberBox(v)
        	return fancyBox
        default:
        return "Return to sender"
    }
}
