package main

import (
	"github.com/ozonva/ova-food-api/internal/flusher"
	"github.com/ozonva/ova-food-api/internal/repo"
	f "github.com/ozonva/ova-food-api/pkg/food"
)

func main() {

	coffee := f.Food{Id: 0, UserId: 0, Type: f.Drinks, Name: "Coffee", PortionSize: 60}
	pizza := f.Food{Id: 1, UserId: 0, Type: f.Foods, Name: "Pizza", PortionSize: 300}
	tea := f.Food{Id: 2, UserId: 1, Type: f.Drinks, Name: "Tea", PortionSize: 100}
	bounty := f.Food{Id: 3, UserId: 2, Type: f.Foods, Name: "Bounty", PortionSize: 100}
	cola := f.Food{Id: 4, UserId: 3, Type: f.Drinks, Name: "Cola", PortionSize: 200}
	slice := []f.Food{coffee, pizza, tea, bounty, cola}

	foodRepo := repo.NewRepo("repoFile.txt")
	fl := flusher.NewFlusher(3, foodRepo)
	fl.Flush(slice)

}
