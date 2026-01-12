package darts
import "math"
func Score(x, y float64) int {
	sum := (x * x) + (y * y)
    newSum := math.Sqrt(sum)

    score := 0
    switch {
        case newSum > 5.0 && newSum <= 10.0:
        	score += 1
        case newSum > 1.0 && newSum <= 5.0:
        	score += 5
        case newSum >= 0.0 && newSum <= 1.0:
        	score += 10
        default:
        	score += 0
    }

    return score
}
