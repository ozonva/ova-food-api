package api

import (
	"context"
	"errors"

	"github.com/ozonva/ova-food-api/internal/repo"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/ozonva/ova-food-api/internal/food"
	desc "github.com/ozonva/ova-food-api/pkg/ova-food-api"
)

type FoodAPI struct {
	desc.UnimplementedOvaFoodApiServer
	repo repo.Repo
}

func NewFoodAPI(r repo.Repo) desc.OvaFoodApiServer {
	return &FoodAPI{repo: r}
}

func (fa *FoodAPI) CreateFoodV1(ctx context.Context, req *desc.CreateFoodV1Request) (*emptypb.Empty, error) {
	if err := req.Validate(); err != nil {
		log.Warn().Msgf("input parameter error: %v", err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	err := fa.repo.AddEntity(food.Food{
		Name:        req.GetFood().Name,
		UserId:      req.GetFood().UserId,
		Type:        uint8(req.GetFood().FoodT.Number()),
		PortionSize: req.GetFood().PortionSize})
	if err != nil {
		log.Warn().Msgf("internal database error: %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	log.Info().Msgf("new food created: %s", req.GetFood())
	return &emptypb.Empty{}, nil
}
func (fa *FoodAPI) DescribeFoodV1(ctx context.Context, req *desc.DescribeFoodV1Request) (*desc.DescribeFoodV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Warn().Msgf("input parameter error: %v", err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	foodId := req.GetFoodId()
	description, err := fa.repo.DescribeEntity(foodId)
	if err != nil {
		if errors.Is(err, repo.HaveNotElementErr) {
			log.Info().Msgf("internal db error: %v", err.Error())
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		log.Warn().Msgf("internal db error: %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	log.Info().Msg("return description")
	return &desc.DescribeFoodV1Response{
		Food: &desc.Food{
			FoodId:      description.Id,
			UserId:      description.UserId,
			FoodT:       desc.FoodType(description.Type),
			Name:        description.Name,
			PortionSize: description.PortionSize,
		},
	}, nil
}
func (fa *FoodAPI) ListFoodsV1(ctx context.Context, req *desc.ListFoodsV1Request) (*desc.ListFoodsV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Warn().Msgf("input parameter error: %v", err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	foods := make(map[uint64]*desc.Food)
	ids := req.GetIds()
	for _, id := range ids {
		elem, err := fa.repo.DescribeEntity(id)
		if err != nil {
			if errors.Is(err, repo.HaveNotElementErr) {
				log.Info().Msgf("internal db error: %v", err.Error())
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
			log.Warn().Msgf("internal db error: %v", err.Error())
			return nil, status.Error(codes.Internal, err.Error())
		}
		foods[elem.Id] = &desc.Food{
			FoodId:      elem.Id,
			UserId:      elem.UserId,
			FoodT:       desc.FoodType(elem.Type),
			Name:        elem.Name,
			PortionSize: elem.PortionSize,
		}
	}
	log.Info().Msg("return elements description")
	return &desc.ListFoodsV1Response{Foods: foods}, nil
}
func (fa *FoodAPI) RemoveFoodV1(ctx context.Context, req *desc.RemoveFoodV1Request) (*emptypb.Empty, error) {
	if err := req.Validate(); err != nil {
		log.Warn().Msgf("input parameter error: %v", err.Error())
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, err.Error())
	}
	id := req.GetFoodId()
	err := fa.repo.RemoveEntity(id)
	if err != nil {
		if errors.Is(err, repo.HaveNotElementErr) {
			log.Info().Msgf("internal db error: %v", err.Error())
			return &emptypb.Empty{}, status.Error(codes.InvalidArgument, err.Error())
		}
		log.Warn().Msgf("internal db error: %v", err.Error())
		return &emptypb.Empty{}, status.Error(codes.Internal, err.Error())
	}
	log.Info().Msgf("food deleted: %v", req.FoodId)
	return &emptypb.Empty{}, nil
}
