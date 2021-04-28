package extjson

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

// NewMap .
func NewMap(dst interface{}) error {
	dscValue := reflect.ValueOf(dst)
	if dscValue.Elem().Kind() != reflect.Map {
		return errors.New("dst error")
	}
	result := reflect.MakeMap(reflect.TypeOf(dst).Elem())
	dscValue.Elem().Set(result)
	return nil
}

// InSlice .
func InSlice(array interface{}, item interface{}) bool {
	values := reflect.ValueOf(array)
	if values.Kind() != reflect.Slice {
		return false
	}

	size := values.Len()
	list := make([]interface{}, size)
	slice := values.Slice(0, size)
	for index := 0; index < size; index++ {
		list[index] = slice.Index(index).Interface()
	}

	for index := 0; index < len(list); index++ {
		if list[index] == item {
			return true
		}
	}
	return false
}

// NewSlice 创建数组
func NewSlice(dsc interface{}, len int) error {
	dstv := reflect.ValueOf(dsc)
	if dstv.Elem().Kind() != reflect.Slice {
		return errors.New("dsc error")
	}

	result := reflect.MakeSlice(reflect.TypeOf(dsc).Elem(), len, len)
	dstv.Elem().Set(result)
	return nil
}

//SliceDelete 删除数组指定下标元素
func SliceDelete(arr interface{}, indexArr ...int) error {
	dstv := reflect.ValueOf(arr)
	if dstv.Elem().Kind() != reflect.Slice {
		return errors.New("dsc error")
	}
	result := reflect.MakeSlice(reflect.TypeOf(arr).Elem(), 0, dstv.Elem().Len()-len(indexArr))
	for index := 0; index < dstv.Elem().Len(); index++ {
		if InSlice(indexArr, index) {
			continue
		}
		result = reflect.Append(result, dstv.Elem().Index(index))
	}

	dstv.Elem().Set(result)
	return nil
}

type refectItem []struct {
	data reflect.Value
	x    int
}

// Len .
func (ri refectItem) Len() int {
	return len(ri)
}

// Swap .
func (ri refectItem) Swap(i, j int) {
	ri[i], ri[j] = ri[j], ri[i]
}

// Less .
func (ri refectItem) Less(i, j int) bool {
	return ri[j].x < ri[i].x
}

//SliceSort 降序
func SliceSort(array interface{}, field string, reverse ...bool) {
	srcV := reflect.ValueOf(array)
	if srcV.Kind() != reflect.Ptr {
		panic("array is not ptr")
	}
	if !srcV.IsValid() || srcV.Elem().Kind() != reflect.Slice {
		panic("SliceSort is not slice data or invalid")
	}

	if srcV.Elem().Len() == 0 {
		return
	}

	sortArray := make(refectItem, srcV.Elem().Len())
	for index := 0; index < srcV.Elem().Len(); index++ {
		value := srcV.Elem().Index(index)
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}
		sortFiled := value.FieldByName(field)
		if !sortFiled.IsValid() {
			panic("SliceSort Filed:" + field + " not exist")
		}
		numFiled, err := strconv.Atoi(fmt.Sprint(sortFiled.Interface()))
		if err != nil {
			panic("SliceSort Filed:" + field + " conversion int type failed")
		}
		sortArray[index].x = numFiled
		sortArray[index].data = srcV.Elem().Index(index)
	}
	if len(reverse) > 0 {
		sort.Sort(sort.Reverse(sortArray))
	} else {
		sort.Sort(sortArray)
	}

	result := reflect.MakeSlice(reflect.TypeOf(array).Elem(), 0, srcV.Elem().Len())
	for index := 0; index < len(sortArray); index++ {
		result = reflect.Append(result, sortArray[index].data)
	}
	srcV.Elem().Set(result)
	return
}

//SliceSortReverse 升序
func SliceSortReverse(array interface{}, field string) {
	SliceSort(array, field, true)
}
