package arrays_slices

func Sum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums[]int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return 
}

func SumAllTails(numbersToSum ...[]int)(sums []int){

	for _, v := range numbersToSum {
		if len(v) ==0 {
			sums = append(sums, 0)
			continue
		}
		tail := v[1:]
		sums = append(sums, Sum(tail))
	}
	return
}