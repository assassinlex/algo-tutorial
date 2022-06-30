package _7_binary_tree

import (
	"fmt"
	"testing"
)

func TestGetMinLightCnt(t *testing.T) {
	times := 10
	length := 20

	t.Log("方案一: 暴力模式")
	for i := 0; i < times; i++ {
		str := randomString(length)
		fmt.Printf("%20s街道: %-21s 最少需要 %d 盏灯\n", "", str, GetMinLightCnt01(str))
	}

	t.Log("方案二: 贪心模式")
	for i := 0; i < times; i++ {
		str := randomString(length)
		fmt.Printf("%20s街道: %-21s 最少需要 %d 盏灯\n", "", str, GetMinLightCnt02(str))
	}
}

func TestDetector(t *testing.T) {
	times := 1000000
	length := 20

	if err := detector(length, times); err != nil {
		t.Log(err.Error())
	}
}
