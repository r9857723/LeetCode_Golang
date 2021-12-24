package main

import "testing"

func Test_(t *testing.T) {
	type args struct {
		nums []int
		k    int
		t    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			//Input: nums = [1,2,3,1], k = 3, t = 0
			//Output: true
			name: "test 1",
			args: args{
				nums: []int{1, 2, 3, 1},
				k:    3,
				t:    2,
			},
			want: true,
		}, {
			//Input: nums = [1,0,1,1], k = 1, t = 2
			//Output: true
			name: "test 2",
			args: args{
				nums: []int{1, 0, 1, 1},
				k:    1,
				t:    2,
			},
			want: true,
		}, {

			//Input: nums = [1,5,9,1,5,9], k = 2, t = 3
			//Output: false
			name: "test 3",
			args: args{
				nums: []int{1, 5, 9, 1, 5, 9},
				k:    2,
				t:    3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyAlmostDuplicate(tt.args.nums, tt.args.k, tt.args.t); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
