package main

import "math"

func countSymmetricIntegers(low int, high int) int {
	total := 0
	for i := low; i < high+1; i++ {
		digs := countDigits(i)
		if digs%2 != 0 {
			continue
		}
		digs /= 2
		pow := int(math.Pow(10, float64(digs)))
		left := i / pow
		right := i % pow
		//fmt.Println(left, right)
		if sumDigits(left) == sumDigits(right) {
			//fmt.Println("takin")
			total++
		}
	}
	return total

}

func sumDigits(n int) int {
	total := 0
	for n > 0 {
		total += n % 10
		n /= 10
	}
	return total

}

func countDigits(n int) int {
	total := 0
	for n > 0 {
		total++
		n = n / 10
	}
	return total
}
