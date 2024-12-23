package leetcode

import (
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{3, 2, 2, 3}, 3}, []int{2, 2}},
		{"test2", args{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2}, []int{0, 1, 3, 0, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.args.nums, tt.args.val)
			if got != len(tt.want) {
				t.Errorf("removeElement() = %v, want %v", got, len(tt.want))
			}
			for i := 0; i < got; i++ {
				if tt.args.nums[i] != tt.want[i] {
					t.Errorf("removeElement() = %v, want %v", tt.args.nums, tt.want)
				}
			}

		})
	}
}
