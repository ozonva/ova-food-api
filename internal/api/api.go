package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/ozonva/ova-food-api/internal/Kafka/producer"
	"github.com/ozonva/ova-food-api/internal/food"
	"github.com/ozonva/ova-food-api/internal/repo"
	"github.com/ozonva/ova-food-api/internal/utils"
	desc "github.com/ozonva/ova-food-api/pkg/ova-food-api"
)

type FoodAPI struct {
	desc.UnimplementedOvaFoodApiServer
	repo      repo.Repo
	chunkSize int
	producer  sarama.SyncProducer
	topic     string
}

func NewFoodAPI(r repo.Repo, cs int, prod sarama.SyncProducer, top string) desc.OvaFoodApiServer {
	return &FoodAPI{
		repo:      r,
		chunkSize: cs,
		producer:  prod,
		topic:     top,
	}
}

func (fa *FoodAPI) CreateFoodV1(ctx context.Context, req *desc.CreateFoodV1Request) (*emptypb.Empty, error) {
	if err := req.Validate(); err != nil {
		log.Warn().Msgf("input parameter error: %v", err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	err := fa.repo.AddEntity(ctx, food.Food{
		Name:        req.GetFood().Name,
		UserId:      req.GetFood().UserId,
		Type:        uint8(req.GetFood().FoodT.Number()),
		PortionSize: req.GetFood().PortionSize})
	if err != nil {
		log.Warn().Msgf("internal database error: %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	msgstr := fmt.Sprintf("new food CREATED: %s", req.GetFood())
	msg := producer.PrepareMessage(fa.topic, msgstr)
	partition, offset, err := fa.producer.SendMessage(msg)
	if err != nil {
		log.Warn().Msgf("sending msg create error: %v", err.Error())
	} else {
		log.Info().Msgf("message %s was send to partition: %d, offset: %d", msgstr, partition, offset)
	}
	return &emptypb.Empty{}, nil
}
func (fa *FoodAPI) DescribeFoodV1(ctx context.Context, req *desc.DescribeFoodV1Request) (*desc.DescribeFoodV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Warn().Msgf("input parameter error: %v", err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	foodId := req.GetFoodId()
	description, err := fa.repo.DescribeEntity(ctx, foodId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
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
		elem, err := fa.repo.DescribeEntity(ctx, id)
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
	err := fa.repo.RemoveEntity(ctx, id)
	if err != nil {
		if errors.Is(err, repo.HaveNotElementErr) {
			log.Info().Msgf("internal db error: %v", err.Error())
			return &emptypb.Empty{}, status.Error(codes.InvalidArgument, err.Error())
		}
		log.Warn().Msgf("internal db error: %v", err.Error())
		return &emptypb.Empty{}, status.Error(codes.Internal, err.Error())
	}
	msgstr := fmt.Sprintf("food DELETED: %v", req.FoodId)
	msg := producer.PrepareMessage(fa.topic, msgstr)
	partition, offset, err := fa.producer.SendMessage(msg)
	if err != nil {
		log.Warn().Msgf("sending msg delete error: %v", err.Error())
	} else {
		log.Info().Msgf("message %s was send to partition: %d, offset: %d", msgstr, partition, offset)
	}
	return &emptypb.Empty{}, nil
}

func (fa *FoodAPI) MultiCreateFoodsV1(ctx context.Context, req *desc.MultiCreateFoodsV1Request) (*emptypb.Empty, error) {
	if err := req.Validate(); err != nil {
		log.Warn().Msgf("input parameter error: %v", err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	foodsReq := req.GetFoods()
	dbFoods := make([]food.Food, 0)
	for _, elem := range foodsReq {
		dbFoods = append(dbFoods, food.Food{
			Id:          elem.UserId,
			UserId:      elem.UserId,
			Type:        uint8(elem.FoodT),
			Name:        elem.Name,
			PortionSize: elem.PortionSize,
		})
	}
	bulks := utils.SplitToBulks(dbFoods, fa.chunkSize)
	err := fa.repo.MultiAddEntity(ctx, bulks)
	if err != nil {
		log.Warn().Msgf("internal database error: %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	log.Info().Msgf("new foods(%v) created", cap(req.GetFoods()))
	return &emptypb.Empty{}, nil
}
func (fa *FoodAPI) PageFoods(ctx context.Context, req *desc.PageFoodsV1Request) (*desc.PageFoodsV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Warn().Msgf("input parameter error: %v", err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	foods := make(map[uint64]*desc.Food)
	dbFoods := []food.Food{}
	dbFoods, err := fa.repo.ListEntities(ctx, req.GetLimit(), req.GetOffset())
	if err != nil {
		if errors.Is(err, repo.HaveNotElementErr) {
			log.Info().Msgf("internal db error: %v", err.Error())
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		log.Warn().Msgf("internal db error: %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	for _, elem := range dbFoods {
		foods[elem.Id] = &desc.Food{
			FoodId:      elem.Id,
			UserId:      elem.UserId,
			FoodT:       desc.FoodType(elem.Type),
			Name:        elem.Name,
			PortionSize: elem.PortionSize,
		}
	}

	log.Info().Msg("return elements description")
	return &desc.PageFoodsV1Response{Foods: foods}, nil
}
func (fa *FoodAPI) UpdateFoodV1(ctx context.Context, req *desc.UpdateFoodV1Request) (*emptypb.Empty, error) {
	if err := req.Validate(); err != nil {
		log.Warn().Msgf("input parameter error: %v", err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	err := fa.repo.UpdateEntity(ctx, food.Food{
		Id:          req.GetFood().FoodId,
		Name:        req.GetFood().Name,
		UserId:      req.GetFood().UserId,
		Type:        uint8(req.GetFood().FoodT.Number()),
		PortionSize: req.GetFood().PortionSize})
	if err != nil {
		if errors.Is(err, repo.HaveNotElementErr) {
			log.Info().Msgf("internal db error: %v", err.Error())
			return &emptypb.Empty{}, status.Error(codes.InvalidArgument, err.Error())
		}
		log.Warn().Msgf("internal database error: %v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	msgstr := fmt.Sprintf("food UPDATED: %s", req.GetFood())
	msg := producer.PrepareMessage(fa.topic, msgstr)
	partition, offset, err := fa.producer.SendMessage(msg)
	if err != nil {
		log.Warn().Msgf("sending msg update error: %v", err.Error())
	} else {
		log.Info().Msgf("message %s was send to partition: %d, offset: %d", msgstr, partition, offset)
	}
	return &emptypb.Empty{}, nil
}
