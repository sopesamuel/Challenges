package diffsquares

func SquareOfSum(n int) int {
	c := 0
    for i := 1; i < n + 1 ; i++ {
        c += i
    }
   return c * c
}

func SumOfSquares(n int) int {
    c := 0
	for i := 1; i < n + 1 ; i++ {
        c += i * i
    }
    return c
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
