package level_02

import (
	"log"
	"testing"
)

func Test_swap(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"", args{1, 3}, 3, 1},
		{"", args{2, 4}, 4, 2},
		{"", args{3, 6}, 6, 3},
		{"", args{1, 2}, 2, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := swap(tt.args.x, tt.args.y)
			if got != tt.want {
				t.Errorf("swap() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("swap() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getNum(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{s: []int{1, 3, 2, 4, 4, 2, 3}}, 1},
		{"", args{s: []int{3, 3, 3, 4, 4, 2, 3}}, 2},
		{"", args{s: []int{3, 3, 2, 4, 4, 2, 3}}, 3},
		{"", args{s: []int{1, 1, 2, 3, 4, 2, 3}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNum(tt.args.s); got != tt.want {
				t.Errorf("getNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAnswer(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// 0110 0000 => 96 | 0010 0000 => 32
		// 0011 0000 => 48 | 0001 0000 => 16
		// 0001 1100 => 28 | 0000 0100 => 4
		// 0001 1000 => 24 | 0000 0100 => 8
		{"", args{x: 96}, 32},
		{"", args{x: 48}, 16},
		{"", args{x: 28}, 4},
		{"", args{x: 24}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAnswer(tt.args.x); got != tt.want {
				t.Errorf("getAnswer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_demo(t *testing.T) {
	m := []map[string]string{
		{"receiver": "prometheus-alert-center", "status": "firing"},
		{"receiver": "prometheus-alert-center", "status": "firing"},
	}

	for _, item := range m {
		log.Println(item["status"])
	}
}