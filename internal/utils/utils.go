package utils

import (
	"errors"
	"math"
)

func SliceToChanks(sliceIn []int, chankSize int) [][]int {
	if sliceIn == nil ||chankSize >= len(sliceIn) || chankSize < 1{
		res := make([][]int,1)
		res[0] = sliceIn
		return res
	} else {
		numChanks := int(math.Ceil(float64(len(sliceIn)) / float64(chankSize)))
		res := make([][]int,numChanks)
		k := -1
		for i:=0;i<len(sliceIn);i++ {
			if i % chankSize == 0 {
				k++
				res[k] = make([]int,0)
			}
			res[k] = append(res[k],sliceIn[i])
		}
		return res
	}
}

func InverseMap(mapIn *map[int]string) (*map[string]int,error){
	res := make(map[string]int)
	for key,val := range *mapIn {
		if _,ok := res[val];ok {
			return nil, errors.New("cant inverse, not uniq values at map")
		}
		res[val] = key
	}
	return &res,nil
}

func SliceFilter(sliceIn []int, filter []int) []int {
	tmpset := make(map[int]bool)
	for _,elem := range filter {
		tmpset[elem] = true
	}
	var newSlice []int
	for _,elem := range sliceIn {
		if _,ok := tmpset[elem]; !ok {
			newSlice = append(newSlice,elem)
		}
	}
	return newSlice
}