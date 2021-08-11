package utils

import (
	"reflect"
	"testing"

	f "github.com/ozonva/ova-food-api/pkg/food"
	"github.com/stretchr/testify/assert"
)

func TestSliceToChanks(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	testTables := []struct {
		chSize int
		chanks [][]int
	}{
		{0, [][]int{{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}},
		{1, [][]int{{0}, {1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}},
		{2, [][]int{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {8, 9}}},
		{3, [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {9}}},
		{11, [][]int{{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}},
	}
	for _, table := range testTables {
		chanksRes := SliceToChanks(slice, table.chSize)
		if !reflect.DeepEqual(table.chanks, chanksRes) {
			t.Errorf("For slice %v split by %d elements chanks not correct,"+
				" expected: %v, got: %v", slice, table.chSize, table.chanks, chanksRes)
		}
	}

}

func TestSplitToBulks(t *testing.T) {
	coffee := f.Food{0, 0, f.Drinks, "Coffee", 60}
	pizza := f.Food{1, 0, f.Foods, "Pizza", 300}
	tea := f.Food{2, 1, f.Drinks, "Tea", 100}
	bounty := f.Food{3, 2, f.Foods, "Bounty", 100}
	cola := f.Food{4, 3, f.Drinks, "Cola", 200}

	slice := []f.Food{coffee, pizza, tea, bounty, cola}
	testTables := []struct {
		chSize         int
		expextedChanks [][]f.Food
	}{
		{0, [][]f.Food{slice}},
		{1, [][]f.Food{{coffee}, {pizza}, {tea}, {bounty}, {cola}}},
		{2, [][]f.Food{{coffee, pizza}, {tea, bounty}, {cola}}},
		{3, [][]f.Food{{coffee, pizza, tea}, {bounty, cola}}},
		{6, [][]f.Food{slice}},
	}
	for _, test := range testTables {
		assert.Equal(t, test.expextedChanks, SplitToBulks(slice, test.chSize))
	}
}

func TestFoodsToMap(t *testing.T) {
	coffee := f.Food{0, 0, f.Drinks, "Coffee", 60}
	pizza := f.Food{1, 0, f.Foods, "Pizza", 300}
	tea := f.Food{2, 1, f.Drinks, "Tea", 100}
	bounty := f.Food{3, 2, f.Foods, "Bounty", 100}
	cola := f.Food{4, 3, f.Drinks, "Cola", 200}

	slice := []f.Food{coffee, pizza, tea, bounty, cola}
	expMap := map[uint64]f.Food{0: coffee, 1: pizza, 2: tea, 3: bounty, 4: cola}
	res, err := FoodsToMap(slice)
	if err == nil {
		assert.Equal(t, expMap, res)
	} else {
		assert.Error(t, err)
	}
}

func TestInverseMap(t *testing.T) {
	testMap := map[int]string{1: "a", 2: "b", 3: "c"}
	resMap := map[string]int{"a": 1, "b": 2, "c": 3}
	inversed, err := InverseMap(testMap)
	if err != nil {
		t.Errorf("Error:%v", err.Error())
	} else if !reflect.DeepEqual(resMap, inversed) {
		t.Errorf("For map %v "+
			" expected inversed map: %v, got: %v", testMap, resMap, inversed)
	}
}
func TestSliceFilter(t *testing.T) {
	testTables := []struct {
		slice  []int
		filter []int
		expRes []int
	}{
		{[]int{0, 1, 2, 3, 4, 5, 6}, []int{2, 4, 6}, []int{0, 1, 3, 5}},
		{[]int{0, 1, 2, 3, 4, 5, 6}, []int{0, 1, 2, 3, 4, 5, 6, 7}, make([]int, 0)},
		{[]int{0, 1, 2, 3, 4, 5, 6}, []int{}, []int{0, 1, 2, 3, 4, 5, 6}},
	}
	for _, table := range testTables {
		gotRes := SliceFilter(table.slice, table.filter)
		if !reflect.DeepEqual(table.expRes, gotRes) {
			t.Errorf("For slice %v filter by %v "+
				" expected: %v, got: %v", table.slice, table.filter, table.expRes, gotRes)
		}
	}
}
