package util

import "testing"

func Test_isArraySortAsc(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test01",
			args{[]int{1, 2, 3}},
			true,
		},
		{
			"test02",
			args{[]int{10, 30, 40, 20, 15, 7}},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsArraySortAsc(tt.args.array); got != tt.want {
				t.Errorf("isArraySort() = %v,want %v", got, tt.want)
			}
		})
	}
}
