package services

import (
	"context"
	"errors"
	"github.com/MukizuL/hezzl-test/internal/dto"
	"github.com/MukizuL/hezzl-test/internal/errs"
	"github.com/MukizuL/hezzl-test/internal/helpers"
	"github.com/MukizuL/hezzl-test/internal/models"
	"go.uber.org/zap"
)

func (s *Services) CreateGoods(ctx context.Context, projectId int, name string) (*dto.CreateGoodsResponse, error) {
	id, err := s.storage.CreateGoods(ctx, projectId, name)
	if err != nil {
		return nil, err
	}

	goods, err := s.storage.GetGood(ctx, id)
	if err != nil {
		return nil, err
	}

	s.logger.InfoNats(goods.ID, goods.ProjectID, goods.Name, goods.Description, goods.Priority, goods.Removed)

	resp := &dto.CreateGoodsResponse{
		ID:          goods.ID,
		ProjectID:   goods.ProjectID,
		Name:        goods.Name,
		Description: goods.Description,
		Priority:    goods.Priority,
		Removed:     goods.Removed,
		CreatedAt:   goods.CreatedAt,
	}

	return resp, nil
}

func (s *Services) UpdateGoods(ctx context.Context, id, projectId int, name, description string) (*dto.CreateGoodsResponse, error) {
	err := s.storage.UpdateGood(ctx, id, projectId, name, description)
	if err != nil {
		return nil, err
	}

	goods, err := s.storage.GetGood(ctx, id)
	if err != nil {
		return nil, err
	}

	s.logger.InfoNats(goods.ID, goods.ProjectID, goods.Name, goods.Description, goods.Priority, goods.Removed)

	resp := &dto.CreateGoodsResponse{
		ID:          goods.ID,
		ProjectID:   goods.ProjectID,
		Name:        goods.Name,
		Description: goods.Description,
		Priority:    goods.Priority,
		Removed:     goods.Removed,
		CreatedAt:   goods.CreatedAt,
	}

	s.storage.Invalidate(ctx)

	return resp, nil
}

func (s *Services) RemoveGoods(ctx context.Context, id, projectId int) (*dto.RemoveGoodsResponse, error) {
	goods, err := s.storage.GetGood(ctx, id)
	if err != nil {
		return nil, err
	}

	err = s.storage.RemoveGoods(ctx, id, projectId)
	if err != nil {
		return nil, err
	}

	s.logger.InfoNats(goods.ID, goods.ProjectID, goods.Name, goods.Description, goods.Priority, true)

	resp := &dto.RemoveGoodsResponse{
		ID:        id,
		ProjectID: projectId,
		Removed:   true,
	}

	s.storage.Invalidate(ctx)

	return resp, nil
}

func (s *Services) Reprioritize(ctx context.Context, id, projectId, newPriority int) ([]dto.ReprioritizeResponse, error) {
	goods, err := s.storage.GetGoodsSortPriority(ctx)
	if err != nil {
		return nil, err
	}

	newGoods, err := helpers.Reprioritize(goods, id, projectId, newPriority)
	if err != nil {
		return nil, err
	}

	for _, good := range newGoods {
		s.logger.InfoNats(good.ID, good.ProjectID, good.Name, good.Description, good.Priority, good.Removed)
	}

	var resp []dto.ReprioritizeResponse
	for _, g := range newGoods {
		resp = append(resp, dto.ReprioritizeResponse{
			ID:       g.ID,
			Priority: g.Priority,
		})
	}

	s.storage.Invalidate(ctx)

	return resp, nil
}

func (s *Services) GetGoods(ctx context.Context, limit, offset int) (*dto.GetGoodsResponse, error) {
	var goods []models.Goods
	goods, err := s.storage.Get(ctx, limit, offset)
	if err != nil {
		if !errors.Is(err, errs.ErrCacheMiss) {
			s.logger.ZapLogger.Error("Redis error", zap.Error(err))
		}
	} else {
		s.logger.ZapLogger.Info("Took from Redis")
		return helpers.GetGoodsResponse(goods, limit, offset), nil
	}

	err = s.storage.Set(ctx)
	if err != nil {
		s.logger.ZapLogger.Error("Redis error", zap.Error(err))
	}

	goods, err = s.storage.GetGoodsWithLimit(ctx, limit, offset)
	if err != nil {
		s.logger.ZapLogger.Info("Took from Postgres")
		return nil, err
	}

	return helpers.GetGoodsResponse(goods, limit, offset), nil
}
