package main

import (
	"fmt"

	myutil "github.com/ozonva/ova-food-api/internal/utils"
)

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	chanks := myutil.SliceToChanks(slice, 3)
	fmt.Println("slice: ", slice)
	fmt.Println("chanks: ", chanks)

	testMap := map[int]string{1: "a", 2: "b", 3: "c"}
	inversed, err := myutil.InverseMap(testMap)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("\nmap: ", testMap)
		fmt.Println("inversed: ", inversed)
	}

	filter := []int{2, 4, 6}
	fmt.Println("\nslice: ", slice)
	fmt.Println("filter: ", filter)
	fmt.Println("filtered slice: ", myutil.SliceFilter(slice, filter))
}
