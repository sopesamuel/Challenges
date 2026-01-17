package romannumerals
import "errors"

func ToRomanNumeral(input int) (string, error) {
    if input < 1 || input > 3999 {
        return "", errors.New("Invalid number")
    }

    num := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    values := []string{"M","CM", "D", "CD", "C", "XC", "L","XL", "X", "IX","V","IV", "I"}

	roman := ""
    sum := input
    for i := 0; i < len(num); i++ {
        if sum >= num[i]{
            for sum >= num[i]{
                sum -= num[i]
                roman += values[i]
            }
        }
    }
    return roman, nil
}
