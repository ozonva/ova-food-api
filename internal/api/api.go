package api

import (
	"context"
	"github.com/ozonva/ova-food-api/internal/food"
	desc "github.com/ozonva/ova-food-api/pkg/ova-food-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type FoodAPI struct {
	desc.UnimplementedOvaFoodApiServer
}

func NewFoodAPI() desc.OvaFoodApiServer  {
	return &FoodAPI{}
}

func (fa *FoodAPI)CreateFoodV1(ctx context.Context, req *desc.CreateFoodV1Request) (*emptypb.Empty,error)  {
	log.Info().Msgf("new food created: %v",req.GetFood().String())
	return &emptypb.Empty{},nil
}
func (fa *FoodAPI)DescribeFoodV1(ctx context.Context, req *desc.DescribeFoodV1Request) (*desc.DescribeFoodV1Response, error){
	foodId := req.GetFoodId()
	coffee := food.Food{Id: foodId, UserId: 0, Type: food.Drinks, Name: "Coffee", PortionSize: 60}
	log.Info().Msgf("return description of coffee")

	//TODO:PortionSize type change
	return &desc.DescribeFoodV1Response{
		Food: &desc.Food{
			FoodId: coffee.Id,
			UserId: coffee.UserId,
			FoodT: desc.FoodType(coffee.Type),
			Name: coffee.Name,
			PortionSize: uint32(coffee.PortionSize),
		},
	},nil
}
//TODO:ListFoodsV1
func (fa *FoodAPI)ListEntities(ctx context.Context, req *desc.ListEntitiesV1Request) (*desc.ListEntitiesV1Response, error){
	//foodIds :=req.GetIds()
	coffee := food.Food{Id: 0, UserId: 0, Type: food.Drinks, Name: "Coffee", PortionSize: 60}
	pizza := food.Food{Id: 1, UserId: 0, Type: food.Foods, Name: "Pizza", PortionSize: 300}
	data := map[uint64]food.Food{coffee.Id:coffee,pizza.Id:pizza}
	foods := make(map[uint64]*desc.Food)

	for fId,val := range data{
		foods[fId] = &desc.Food{
			FoodId:fId,
			UserId: val.UserId,
			FoodT: desc.FoodType(val.Type),
			Name: val.Name,
			PortionSize: uint32(val.PortionSize),
		}
	}
	log.Info().Msgf("return description of coffee & pizza")
	return &desc.ListEntitiesV1Response{Foods: foods},nil
}

func (fa *FoodAPI)RemoveFoodV1(ctx context.Context,req *desc.RemoveFoodV1Request) (*emptypb.Empty,error){
	log.Info().Msgf("food deleted: %v",req.FoodId)
	return &emptypb.Empty{},nil
}


