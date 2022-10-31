package sum

import "testing"

// 测试覆盖率
// go test -cover
func TestSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	got := Sum(nums)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, nums)
	}

	t.Run("collection of any size", func(t *testing.T) {
		nums := []int{1, 2, 3}

		got := Sum(nums)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, nums)
		}
	})
}
