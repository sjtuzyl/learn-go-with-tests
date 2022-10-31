package sum

func Sum(nums []int) int {
	ret := 0
	for _, val := range nums {
		ret += val
	}
	return ret
}

func SumAll(numss ...[]int) []int {
	rets := make([]int, len(numss))
	for ind, nums := range numss {
		rets[ind] = Sum(nums)
	}
	return rets
}

func SumAllTails(numss ...[]int) []int {
	var rets []int
	for _, nums := range numss {
		if len(nums) == 0 {
			rets = append(rets, 0)
		} else {
			tails := nums[1:]
			rets = append(rets, Sum(tails))
		}
	}
	return rets
}
