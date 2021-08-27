package api

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/ozonva/ova-food-api/internal/food"
	desc "github.com/ozonva/ova-food-api/pkg/ova-food-api"
)

type FoodAPI struct {
	desc.UnimplementedOvaFoodApiServer
}

func NewFoodAPI() desc.OvaFoodApiServer  {
	return &FoodAPI{}
}

func (fa *FoodAPI)CreateFoodV1(ctx context.Context, req *desc.CreateFoodV1Request) (*emptypb.Empty,error)  {
	if err := req.Validate(); err != nil {
		log.Fatal().Msgf("input parameter error: %v",err.Error())
		return nil,status.Error(codes.InvalidArgument, err.Error())
	}
	log.Info().Msgf("new food created: %s",req.GetFood())
	return &emptypb.Empty{},nil
}
func (fa *FoodAPI)DescribeFoodV1(ctx context.Context, req *desc.DescribeFoodV1Request) (*desc.DescribeFoodV1Response, error){
	if err := req.Validate(); err != nil {
		log.Fatal().Msgf("input parameter error: %v",err.Error())
		return nil,status.Error(codes.InvalidArgument, err.Error())
	}
	foodId := req.GetFoodId()
	coffee := food.Food{Id: foodId, Type: food.Drinks, Name: "Coffee", PortionSize: 60}
	log.Info().Msg("return description of coffee")

	return &desc.DescribeFoodV1Response{
		Food: &desc.Food{
			FoodId: coffee.Id,
			UserId: coffee.UserId,
			FoodT: desc.FoodType(coffee.Type),
			Name: coffee.Name,
			PortionSize: coffee.PortionSize,
		},
	},nil
}

func (fa *FoodAPI)ListFoodsV1(ctx context.Context, req *desc.ListFoodsV1Request) (*desc.ListFoodsV1Response, error){
	if err := req.Validate(); err != nil {
		log.Fatal().Msgf("input parameter error: %v",err.Error())
		return nil,status.Error(codes.InvalidArgument, err.Error())
	}
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
			PortionSize: val.PortionSize,
		}
	}
	log.Info().Msg("return description of coffee & pizza")
	return &desc.ListFoodsV1Response{Foods: foods},nil
}

func (fa *FoodAPI)RemoveFoodV1(ctx context.Context,req *desc.RemoveFoodV1Request) (*emptypb.Empty,error){
	if err := req.Validate(); err != nil {
		log.Fatal().Msgf("input parameter error: %v",err.Error())
		return nil,status.Error(codes.InvalidArgument, err.Error())
	}
	log.Info().Msgf("food deleted: %v",req.FoodId)
	return &emptypb.Empty{},nil
}


