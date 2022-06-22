package level_01

import (
	"log"
	"math/rand"
	"testing"
)

func TestBitMap(t *testing.T) {
	max := 10000
	bitMap := NewBitMap(max)
	set := make(map[int]struct{})
	testTimes := 10000000
	for i := 0; i < testTimes; i++ {
		num := rand.Intn(max + 1)
		decide := rand.Float32()
		if decide < 3.33 {
			bitMap.add(num)
			set[num] = struct{}{}
		} else if decide < 6.66 {
			bitMap.delete(num)
			delete(set, num)
		} else {
			if _, ok := set[num]; ok != bitMap.contain(num) {
				log.Fatal("Oops!")
			}
		}
	}
}