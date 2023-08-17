package main

import (
	"fmt"
	"math"
	"time"
)

var count uint64

func solveLuckyTicketRecursive(n, sumA, sumB uint64) {
	if n == 0 {
		if sumA == sumB {
			count++
		}
		return
	}

	for a := uint64(0); a <= 9; a++ {
		for b := uint64(0); b <= 9; b++ {
			solveLuckyTicketRecursive(n-1, sumA+a, sumB+b)
		}
	}
}

func sum_d(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

func solveLuckyTicketSum(n uint64) uint64 {
	var (
		count int
		comb  = make(map[int][]string, n*9+1)
	)

	for i := 0; i < int(math.Pow10(int(n))); i++ {
		comb[sum_d(i)] = append(comb[sum_d(i)], fmt.Sprintf("%0*d", n, i))
	}

	for i := 0; i < len(comb); i++ {
		count += len(comb[i]) * len(comb[i])
	}
	return uint64(count)
}

func solveLuckyTicketTable(n uint64) uint64 {
	if n <= 0 {
		return 0
	}
	var (
		count      uint64
		luckyTable = make([]uint64, n*9+1)
	)

	copy(luckyTable, []uint64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1})

	for i := uint64(0); i < n-1; i++ {
		tempTable := make([]uint64, n*9+1)
		tempTable[0] = 1
		for j := 1; j < len(tempTable); j++ {
			if j < 10 {
				tempTable[j] = tempTable[j-1] + luckyTable[j]
			} else {
				for s := j - 9; s <= j; s++ {
					tempTable[j] += luckyTable[s]
				}
			}
		}
		copy(luckyTable, tempTable)
	}

	for i := 0; i < len(luckyTable); i++ {
		count += luckyTable[i] * luckyTable[i]
	}

	return count
}

// func solveLuckyTicketApprox(n float64) float64 {
// 	return math.Pow(10, 2*n) / math.Sqrt(33*n*math.Pi)
// }

func main() {
	start := time.Now()
	fmt.Printf("Lucky N: %d  took %vns\n", solveLuckyTicketTable(10), time.Since(start).Nanoseconds())
	start = time.Now()
	fmt.Printf("Lucky N: %d  took %vns\n", solveLuckyTicketSum(7), time.Since(start).Nanoseconds())
	start = time.Now()
	solveLuckyTicketRecursive(5, 0, 0)
	fmt.Printf("Lucky N: %d  took %vns\n", count, time.Since(start).Nanoseconds())
}
