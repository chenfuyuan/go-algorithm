package is_valid_bst

import (
	"go-algorithm/src/main/datastruct"
	"go-algorithm/src/main/util"
	"testing"
)

func Test_isValidBST_LDR_Recursion(t *testing.T) {
	type args struct {
		root *datastruct.BitTreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test_01",
			args{util.GeneratorTree([]int{2, 1, 3})},
			true,
		},
		{
			"test_02",
			args{util.GeneratorTree([]int{5, 1, 4, -1, -1, 3, 6})},
			false,
		},
		{
			"test_03",
			args{util.GeneratorTree([]int{5, 4, 6, -1, -1, 3, 7})},
			false,
		},
		{
			"test_04",
			args{util.GeneratorTree([]int{0})},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidBSTLdrRecursion(tt.args.root); got != tt.want {
				t.Errorf("IsValidBSTLdrRecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}
