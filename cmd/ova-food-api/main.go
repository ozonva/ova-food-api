package main

import (
	"fmt"

	"github.com/ozonva/ova-food-api/internal/utils"
	f "github.com/ozonva/ova-food-api/pkg/food"
)

func main() {

	utils.GetConfigFromFiles("config.txt")

	coffee := f.Food{0, 0, f.Drinks, "Coffee", 60.0}
	fmt.Println(coffee.String())

	pizza := f.Food{1, 0, f.Foods, "Pizza", 300}
	pizzaObj := f.CreateFood([]byte(pizza.String()))
	fmt.Println(pizzaObj.String())
}
