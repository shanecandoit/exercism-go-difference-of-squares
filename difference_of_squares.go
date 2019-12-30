package diffsquares

/*
The square of the sum of the first ten natural numbers is
(1 + 2 + ... + 10)² = 55² = 3025.

The sum of the squares of the first ten natural numbers is
1² + 2² + ... + 10² = 385.

Hence the difference between the square of the sum of the first
ten natural numbers and the sum of the squares of the first ten
natural numbers is 3025 - 385 = 2640.

--- FAIL: TestSquareOfSum (0.00s)
    difference_of_squares_test.go:14: SquareOfSum(5) = 100, want 225
--- FAIL: TestSumOfSquares (0.00s)
    difference_of_squares_test.go:22: SumOfSquares(5) = 30, want 55
--- FAIL: TestDifference (0.00s)
	difference_of_squares_test.go:31: Difference(5) = 70, want 170

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

/*
func main() {
	five := SquareOfSum(5)
	fmt.Println("sq", five)

	sum5 := SumOfSquares(5)
	fmt.Println("sum5", sum5)

	diff5 := Difference(5)
	fmt.Println("diff5", diff5)
}
*/
