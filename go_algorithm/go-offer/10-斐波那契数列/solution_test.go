package _0_斐波那契数列

func fib(n int) int {
	if n <= 1 {
		return n
	}
	fib_one := 0
	fib_two := 1
	fib_n := 0
	for i := 2; i <= n; i++ {
		fib_n = (fib_one + fib_two) % 1000000007
		fib_one = fib_two
		fib_two = fib_n
	}
	return fib_n
}

// 使用动态规划解法

//递归方法：当n=43时时间超出限制
func fib2(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
