package main

import (
	"time"

	"github.com/ozonva/ova-food-api/internal/flusher"
	"github.com/ozonva/ova-food-api/internal/food"
	"github.com/ozonva/ova-food-api/internal/repo"
	"github.com/ozonva/ova-food-api/internal/saver"
)

func main() {
	//utils.UpdateConfig("configs/config.txt")
	coffee := food.Food{Id: 0, UserId: 0, Type: food.Drinks, Name: "Coffee", PortionSize: 60}
	pizza := food.Food{Id: 1, UserId: 0, Type: food.Foods, Name: "Pizza", PortionSize: 300}
	tea := food.Food{Id: 2, UserId: 1, Type: food.Drinks, Name: "Tea", PortionSize: 100}
	bounty := food.Food{Id: 3, UserId: 2, Type: food.Foods, Name: "Bounty", PortionSize: 100}
	cola := food.Food{Id: 4, UserId: 3, Type: food.Drinks, Name: "Cola", PortionSize: 200}
	//slice := []food.Food{coffee, pizza, tea, bounty, cola}

	foodRepo := repo.NewRepo("repoFile.txt")
	fl := flusher.NewFlusher(2, foodRepo)
	saver := saver.NewSaver(2, fl)

	saver.Save(coffee)
	time.Sleep(1500 * time.Millisecond)
	saver.Save(pizza)
	saver.Save(cola)

	saver.Save(bounty)
	saver.Save(tea)

	saver.Save(coffee)

	saver.Close()
}
