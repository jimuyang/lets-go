package subpack

import "testing"

func TestReverseInt(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"", args{[]int{1, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseInt(tt.args.arr)
		})
	}
}
