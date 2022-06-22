package level_02

import (
	"log"
	"testing"
)

func Test_generateSingleList(t *testing.T) {

	s := []int{1, 2, 3, 4, 5}

	head := generateSingleList(s)

	for head != nil {
		log.Println(head.value)
		head = head.next
	}

}

func Test_singleListReverse(t *testing.T)  {
	l := generateSingleList([]int{1, 2, 3, 4, 5})

	for _l := l; _l != nil; {
		log.Println(_l.value)
		_l = _l.next
	}

	log.Println("========")

	m := singleListReverse(l)

	for m != nil {
		log.Println(m.value)
		m = m.next
	}
}

func Test_getMax(t *testing.T) {
	type args struct {
		s     []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{s: []int{0, 1, 2, 3, 4}, left: 0, right: 4}, 4},
		{"", args{s: []int{2, 1, -1, 0, 1}, left: 0, right: 4}, 2},
		{"", args{s: []int{-10, 1, 2, 3, 5, 4}, left: 0, right: 5}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMax(tt.args.s, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("getMax() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_queue(t *testing.T) {
	q := NewQueue(&Node{value: 0})

	for i := 1; i <= 10; i += 2 {
		q.enqueue(&Node{value: i})
	}
	for !q.isEmpty() {
		log.Println(q.dequeue().value)
	}

	log.Println("===============")

	for i := 2; i <= 10; i += 2 {
		q.enqueue(&Node{value: i})
	}
	for !q.isEmpty() {
		log.Println(q.dequeue().value)
	}
}

func Test_stack(t *testing.T) {
	s := NewStack(&Node{value: 0})

	log.Println("push: ")
	for i := 1; i <= 10; i += 2 {
		log.Println(i)
		s.push(&Node{value: i})
	}
	log.Println("pop: ")
	for !s.isEmpty() {
		log.Println(s.pop().value)
	}

	log.Println("===============")

	log.Println("push: ")
	for i := 2; i <= 10; i += 2 {
		log.Println(i)
		s.push(&Node{value: i})
	}
	log.Println("pop: ")
	for !s.isEmpty() {
		log.Println(s.pop().value)
	}
}