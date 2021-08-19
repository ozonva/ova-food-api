package utils

import (
	"testing"

	"github.com/ozonva/ova-food-api/internal/food"

	"github.com/stretchr/testify/assert"
)

func TestSliceToChanks(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	testTables := []struct {
		chSize         int
		expextedChanks [][]int
	}{
		{0, [][]int{{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}},
		{1, [][]int{{0}, {1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}},
		{2, [][]int{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {8, 9}}},
		{3, [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {9}}},
		{11, [][]int{{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}},
	}
	for _, test := range testTables {
		assert.Equal(t, test.expextedChanks, SliceToChanks(slice, test.chSize))
	}

}

func TestSplitToBulks(t *testing.T) {
	coffee := food.Food{Id: 0, UserId: 0, Type: food.Drinks, Name: "Coffee", PortionSize: 60}
	pizza := food.Food{Id: 1, UserId: 0, Type: food.Foods, Name: "Pizza", PortionSize: 300}
	tea := food.Food{Id: 2, UserId: 1, Type: food.Drinks, Name: "Tea", PortionSize: 100}
	bounty := food.Food{Id: 3, UserId: 2, Type: food.Foods, Name: "Bounty", PortionSize: 100}
	cola := food.Food{Id: 4, UserId: 3, Type: food.Drinks, Name: "Cola", PortionSize: 200}

	slice := []food.Food{coffee, pizza, tea, bounty, cola}
	testTables := []struct {
		chSize         int
		expextedChanks [][]food.Food
	}{
		{0, [][]food.Food{slice}},
		{1, [][]food.Food{{coffee}, {pizza}, {tea}, {bounty}, {cola}}},
		{2, [][]food.Food{{coffee, pizza}, {tea, bounty}, {cola}}},
		{3, [][]food.Food{{coffee, pizza, tea}, {bounty, cola}}},
		{6, [][]food.Food{slice}},
	}
	for _, test := range testTables {
		assert.Equal(t, test.expextedChanks, SplitToBulks(slice, test.chSize))
	}
}

func TestFoodsToMap(t *testing.T) {
	coffee := food.Food{Id: 0, UserId: 0, Type: food.Drinks, Name: "Coffee", PortionSize: 60}
	pizza := food.Food{Id: 1, UserId: 0, Type: food.Foods, Name: "Pizza", PortionSize: 300}
	tea := food.Food{Id: 2, UserId: 1, Type: food.Drinks, Name: "Tea", PortionSize: 100}
	bounty := food.Food{Id: 3, UserId: 2, Type: food.Foods, Name: "Bounty", PortionSize: 100}
	cola := food.Food{Id: 4, UserId: 3, Type: food.Drinks, Name: "Cola", PortionSize: 200}

	slice := []food.Food{coffee, pizza, tea, bounty, cola}
	expMap := map[uint64]food.Food{0: coffee, 1: pizza, 2: tea, 3: bounty, 4: cola}
	res, err := FoodsToMap(slice)
	assert.NoError(t, err)
	assert.Equal(t, expMap, res)
}

func TestInverseMap(t *testing.T) {
	testMap := map[int]string{1: "a", 2: "b", 3: "c"}
	expMap := map[string]int{"a": 1, "b": 2, "c": 3}
	inversedMap, err := InverseMap(testMap)
	assert.NoError(t, err)
	assert.Equal(t, expMap, inversedMap)
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
		assert.Equal(t, table.expRes, gotRes)
	}
}
