package sum

func Sum(nums []int) int {
	ret := 0
	for _, val := range nums {
		ret += val
	}
	return ret
}
