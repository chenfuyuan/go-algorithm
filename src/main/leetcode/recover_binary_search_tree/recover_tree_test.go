package recover_binary_search_tree

import (
	"go-algorithm/src/main/datastruct"
	"go-algorithm/src/main/util"
	"testing"
)

func TestRecoverTree(t *testing.T) {
	type args struct {
		root *datastruct.BitTreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test01",
			args{util.GeneratorTree([]int{1, 3, -1, -1, 2})},
		},
		{
			"test02",
			args{util.GeneratorTree([]int{3, 1, 4, -1, -1, 2})},
		},
		{
			"test03",
			args{util.GeneratorTree([]int{4, 6, 5, 1, 3, -1, 2})},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RecoverTree(tt.args.root)
			inorderNums := tt.args.root.InorderGenerate()
			checkResult := util.IsArraySortAsc(inorderNums)
			if !checkResult {
				t.Errorf("args = %v run result is Error", tt.args)
			}
		})
	}
}
