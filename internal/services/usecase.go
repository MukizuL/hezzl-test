package services

import (
	"context"
	"github.com/MukizuL/hezzl-test/internal/dto"
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

	return resp, nil
}

func (s *Services) RemoveGoods(ctx context.Context, id, projectId int) (*dto.RemoveGoodsResponse, error) {
	err := s.storage.RemoveGoods(ctx, id, projectId)
	if err != nil {
		return nil, err
	}

	resp := &dto.RemoveGoodsResponse{
		ID:        id,
		ProjectID: projectId,
		Removed:   true,
	}

	return resp, nil
}
