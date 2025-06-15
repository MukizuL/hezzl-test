package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/MukizuL/hezzl-test/internal/errs"
	"github.com/MukizuL/hezzl-test/internal/models"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"time"
)

func (s *Storage) Get(ctx context.Context, limit, offset int) ([]models.Goods, error) {
	rawData, err := s.redis.Get(ctx, "goods").Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, errs.ErrCacheMiss
		}

		s.logger.Error("Failed to get from cache",
			zap.String("method", "Get"),
			zap.Error(err))

		return nil, errs.ErrInternalServerError
	}

	data := bytes.NewBufferString(rawData)

	var goods []models.Goods
	err = json.NewDecoder(data).Decode(&goods)
	if err != nil {
		s.logger.Error("Failed to unmarshal data",
			zap.String("method", "Get"),
			zap.Error(err))

		return nil, errs.ErrInternalServerError
	}

	if offset >= len(goods) {
		return []models.Goods{}, nil
	}

	if offset+limit > len(goods) {
		return goods[offset:], nil
	}

	return goods[offset : offset+limit+1], nil
}

func (s *Storage) Set(ctx context.Context) error {
	goods, err := s.GetGoodsSortId(ctx)
	if err != nil {
		return err
	}

	data := new(bytes.Buffer)

	err = json.NewEncoder(data).Encode(goods)
	if err != nil {
		s.logger.Error("Failed to marshal data",
			zap.String("method", "Set"),
			zap.Error(err))

		return errs.ErrInternalServerError
	}

	err = s.redis.Set(ctx, "goods", data.String(), time.Minute*1).Err()
	if err != nil {
		s.logger.Error("Failed to set cache",
			zap.String("method", "Set"),
			zap.Error(err))

		return errs.ErrInternalServerError
	}

	return nil
}

func (s *Storage) Invalidate(ctx context.Context) {
	s.redis.Del(ctx, "goods")
}
