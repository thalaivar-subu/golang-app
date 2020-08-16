package primenumber

func FindPrime(n int, primeNumbers chan []int) {
	slice := make([]int, 0)
	for i := 2; i <= n; i++ {
		flag := make(chan bool)
		go isPrime(i, flag)
		boolFlag := <-flag
		if boolFlag {
			slice = append(slice, i)
		}
	}
	primeNumbers <- slice
}

func isPrime(n int, flag chan bool) {
	if n <= 1 {
		flag <- false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			flag <- false
		}
	}
	flag <- true
}
