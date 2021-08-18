package utils

import (
	"errors"
	"math"

	f "github.com/ozonva/ova-food-api/pkg/food"
)

func SliceToChanks(sliceIn []int, chankSize int) [][]int {
	if sliceIn == nil || chankSize >= len(sliceIn) || chankSize < 1 {
		res := make([][]int, 1)
		res[0] = sliceIn
		return res
	} else {
		numChanks := int(math.Ceil(float64(len(sliceIn)) / float64(chankSize)))
		res := make([][]int, numChanks)
		k := -1
		for i := 0; i < len(sliceIn); i++ {
			if i%chankSize == 0 {
				k++
				res[k] = make([]int, 0)
			}
			res[k] = append(res[k], sliceIn[i])
		}
		return res
	}
}

func SplitToBulks(sliceIn []f.Food, chankSize int) [][]f.Food {
	if sliceIn == nil || chankSize < 1 {
		return nil
	} else if chankSize >= len(sliceIn) {
		res := make([][]f.Food, 1)
		res[0] = sliceIn
		return res
	} else {
		numChanks := int(math.Ceil(float64(len(sliceIn)) / float64(chankSize)))
		res := make([][]f.Food, numChanks)
		k := -1
		for i := 0; i < len(sliceIn); i++ {
			if i%chankSize == 0 {
				k++
				res[k] = make([]f.Food, 0)
			}
			res[k] = append(res[k], sliceIn[i])
		}
		return res
	}
}

func FoodsToMap(foods []f.Food) (map[uint64]f.Food, error) {
	res := make(map[uint64]f.Food)
	for _, elem := range foods {
		if _, ok := res[elem.Id]; ok {
			return nil, errors.New("cant create map, not uniq id at foods")
		}
		res[elem.Id] = elem
	}
	return res, nil
}

func InverseMap(mapIn map[int]string) (map[string]int, error) {
	res := make(map[string]int)
	for key, val := range mapIn {
		if _, ok := res[val]; ok {
			return nil, errors.New("cant inverse, not uniq values at map")
		}
		res[val] = key
	}
	return res, nil
}

func SliceFilter(sliceIn []int, filter []int) []int {
	tmpset := make(map[int]bool)
	for _, elem := range filter {
		tmpset[elem] = true
	}
	newSlice := make([]int, 0)
	for _, elem := range sliceIn {
		if _, ok := tmpset[elem]; !ok {
			newSlice = append(newSlice, elem)
		}
	}
	return newSlice
}
