package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "adds two numbers",
			args: args{
				x: 1,
				y: 2,
			},
			want: 3,
		},
		{
			name: "adds two negative numbers",
			args: args{
				x: -1,
				y: -2,
			},
			want: -3,
		},
		{
			name: "adds zeros",
			args: args{
				x: 0,
				y: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Adder(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Adder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleAdder() {
	sum := Adder(5, 5)
	fmt.Println(sum)
	// Output: 10
}

func TestExampleAdder(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ExampleAdder()
		})
	}
}
