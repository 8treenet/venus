package utils

import "testing"

func TestSliceUp(t *testing.T) {
	var numberList []int
	for i := 0; i < 50; i++ {
		numberList = append(numberList, i)
	}

	var numbeResult [][]int
	SliceUp(numberList, &numbeResult, 13) //把numberList以13长度切割成二维数组
	for _, subNumberList := range numbeResult {
		t.Log("subNumberList", subNumberList)
	}

	strList := []string{"a", "b", "c", "d", "e", "f", "g", "h", "fuck"}
	var strResult [][]string
	SliceUp(strList, &strResult, 2) //把strList以2长度切割成二维数组
	for _, v := range strResult {
		t.Log("subStrList", v)
	}
}

func TestInSlice(t *testing.T) {
	if !InSlice([]string{"1", "2", "3", "4"}, "1") {
		panic("")
	}

	if !InSlice([]int{1, 2, 3, 4, 5}, 3) {
		panic("")
	}
}
