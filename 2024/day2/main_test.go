package main

import "testing"

func Test_isSafe(t *testing.T) {
	type args struct {
		report []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Safe without removing any level",
			args: args{report: []int{7, 6, 4, 2, 1}},
			want: true,
		},
		{
			name: "Unsafe regardless of which level is removed",
			args: args{report: []int{1, 2, 7, 8, 9}},
			want: false,
		},
		{
			name: "Unsafe regardless of which level is removed",
			args: args{report: []int{9, 7, 6, 2, 1}},
			want: false,
		},
		{
			name: "Safe by removing the second level, 3",
			args: args{report: []int{1, 3, 2, 4, 5}},
			want: true,
		},
		{
			name: "Safe by removing the third level, 4",
			args: args{report: []int{8, 6, 4, 4, 1}},
			want: true,
		},
		{
			name: "Safe without removing any level",
			args: args{report: []int{1, 3, 6, 7, 9}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSafe(tt.args.report); got != tt.want {
				t.Errorf("isSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}
