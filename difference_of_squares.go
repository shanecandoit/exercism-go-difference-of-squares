package diffsquares

/*
The square of the sum of the first ten natural numbers is
(1 + 2 + ... + 10)² = 55² = 3025.

The sum of the squares of the first ten natural numbers is
1² + 2² + ... + 10² = 385.

SquareOfSum(5)  = 225
SumOfSquares(5) = 55
Difference(5)   = 170
*/

// SquareOfSum sums then square
func SquareOfSum(n int) (sqr int) {

	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	sqr = sum * sum
	return sqr
}

// SumOfSquares squares then sum
func SumOfSquares(n int) (sum int) {
	sum = 0
	for i := 0; i <= n; i++ {
		sum += i * i
	}
	return sum
}

// Difference of SquareOfSum and SumOfSquares
func Difference(n int) (diff int) {
	diff = SquareOfSum(n) - SumOfSquares(n)
	return diff
}
