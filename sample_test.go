package goSamplePackage

import "testing"

func TestFibonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Fibonacci(0)",
			args{0},
			0,
		},
		{
			"Fibonacci(1)",
			args{1},
			1,
		},
		{
			"Fibonacci(2)",
			args{2},
			1,
		},
		{
			"Fibonacci(3)",
			args{3},
			2,
		},
		{
			"Fibonacci(4)",
			args{4},
			3,
		},
		{
			"Fibonacci(5)",
			args{5},
			5,
		},
		{
			"Fibonacci(6)",
			args{6},
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibonacci(tt.args.n); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
