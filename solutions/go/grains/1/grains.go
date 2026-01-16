package grains
import "math"
import "errors"
func Square(number int) (uint64, error) {
    if number <= 0 || number > 64 {
        return 0, errors.New("Invalid number")
    }
	return uint64(math.Pow(2, float64(number - 1))), nil
}

func Total() uint64 {
	sum := uint64(0)
    for i := 1; i < 65; i++ {
        re, err := Square(i)
        if err == nil {
            sum += re
        }
    }
    return sum
}
