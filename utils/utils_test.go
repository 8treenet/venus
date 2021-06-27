package utils

import "testing"

func TestSliceUp(t *testing.T) {
	var numberList []int
	for i := 0; i < 300; i++ {
		numberList = append(numberList, i)
	}

	var numbeResult [][]int
	SliceUp(numberList, &numbeResult, 55) //把numberList以55长度切割成二维数组
	t.Log(len(numbeResult))

	strList := []string{"a", "b", "c", "d", "e", "f", "g", "h", "fuck"}
	var strResult [][]string
	e := SliceUp(strList, &strResult, 2) //把strList以2长度切割成二维数组
	t.Log(strResult, e)
}

func TestInSlice(t *testing.T) {
	if !InSlice([]string{"1", "2", "3", "4"}, "1") {
		panic("")
	}

	if !InSlice([]int{1, 2, 3, 4, 5}, 3) {
		panic("")
	}
}
