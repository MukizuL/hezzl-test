package helpers

import (
	"github.com/MukizuL/hezzl-test/internal/dto"
	"github.com/MukizuL/hezzl-test/internal/errs"
	"github.com/MukizuL/hezzl-test/internal/models"
	"sort"
)

func Reprioritize(goods []models.Goods, id, projectID, newPriority int) ([]models.Goods, error) {
	var targetIndex = -1
	for i, g := range goods {
		if g.ProjectID == projectID && g.ID == id {
			targetIndex = i
			break
		}
	}

	if targetIndex == -1 {
		return goods, errs.ErrGoodsNotFound
	}

	oldPriority := goods[targetIndex].Priority

	if oldPriority == newPriority {
		return goods, nil
	}

	result := make([]models.Goods, len(goods))
	copy(result, goods)

	result[targetIndex].Priority = newPriority

	for i := range result {
		if i == targetIndex {
			continue
		}

		if oldPriority < newPriority {
			if result[i].Priority > oldPriority && result[i].Priority <= newPriority {
				result[i].Priority--
			}
		} else {
			if result[i].Priority >= newPriority && result[i].Priority < oldPriority {
				result[i].Priority++
			}
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})

	return result, nil
}

func GetGoodsResponse(goods []models.Goods, limit, offset int) *dto.GetGoodsResponse {
	var respGoods []dto.CreateGoodsResponse
	total, removed := 0, 0

	for _, good := range goods {
		total++
		respGoods = append(respGoods, dto.CreateGoodsResponse{
			ID:          good.ID,
			ProjectID:   good.ProjectID,
			Name:        good.Name,
			Description: good.Description,
			Priority:    good.Priority,
			Removed:     good.Removed,
			CreatedAt:   good.CreatedAt,
		})
		if good.Removed {
			removed++
		}
	}

	meta := &dto.Meta{
		Total:   total,
		Removed: removed,
		Limit:   limit,
		Offset:  offset,
	}

	return &dto.GetGoodsResponse{
		Meta:  meta,
		Goods: respGoods,
	}
}
