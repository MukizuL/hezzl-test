package helpers

import (
	"github.com/MukizuL/hezzl-test/internal/errs"
	"github.com/MukizuL/hezzl-test/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var timeNow = time.Date(2025, 6, 14, 0, 0, 0, 0, time.UTC)

func TestApplication_Reprioritize(t *testing.T) {
	startData := []models.Goods{
		{ID: 1, ProjectID: 1, Name: "", Description: "", Priority: 1, Removed: false, CreatedAt: timeNow},
		{ID: 2, ProjectID: 1, Name: "", Description: "", Priority: 2, Removed: false, CreatedAt: timeNow},
		{ID: 3, ProjectID: 1, Name: "", Description: "", Priority: 3, Removed: false, CreatedAt: timeNow},
		{ID: 4, ProjectID: 1, Name: "", Description: "", Priority: 4, Removed: false, CreatedAt: timeNow},
		{ID: 5, ProjectID: 1, Name: "", Description: "", Priority: 5, Removed: false, CreatedAt: timeNow},
		{ID: 6, ProjectID: 1, Name: "", Description: "", Priority: 6, Removed: false, CreatedAt: timeNow},
		{ID: 7, ProjectID: 1, Name: "", Description: "", Priority: 7, Removed: false, CreatedAt: timeNow},
		{ID: 8, ProjectID: 1, Name: "", Description: "", Priority: 8, Removed: false, CreatedAt: timeNow},
		{ID: 9, ProjectID: 1, Name: "", Description: "", Priority: 9, Removed: false, CreatedAt: timeNow},
		{ID: 10, ProjectID: 1, Name: "", Description: "", Priority: 10, Removed: false, CreatedAt: timeNow},
	}

	tests := []struct {
		name        string
		goods       []models.Goods
		id          int
		projectId   int
		newPriority int
		expectError bool
		error       error
		want        []models.Goods
	}{
		{
			name:        "Same priority 1",
			goods:       startData,
			id:          1,
			projectId:   1,
			newPriority: 1,
			expectError: false,
			error:       nil,
			want: []models.Goods{
				{ID: 1, ProjectID: 1, Name: "", Description: "", Priority: 1, Removed: false, CreatedAt: timeNow},
				{ID: 2, ProjectID: 1, Name: "", Description: "", Priority: 2, Removed: false, CreatedAt: timeNow},
				{ID: 3, ProjectID: 1, Name: "", Description: "", Priority: 3, Removed: false, CreatedAt: timeNow},
				{ID: 4, ProjectID: 1, Name: "", Description: "", Priority: 4, Removed: false, CreatedAt: timeNow},
				{ID: 5, ProjectID: 1, Name: "", Description: "", Priority: 5, Removed: false, CreatedAt: timeNow},
				{ID: 6, ProjectID: 1, Name: "", Description: "", Priority: 6, Removed: false, CreatedAt: timeNow},
				{ID: 7, ProjectID: 1, Name: "", Description: "", Priority: 7, Removed: false, CreatedAt: timeNow},
				{ID: 8, ProjectID: 1, Name: "", Description: "", Priority: 8, Removed: false, CreatedAt: timeNow},
				{ID: 9, ProjectID: 1, Name: "", Description: "", Priority: 9, Removed: false, CreatedAt: timeNow},
				{ID: 10, ProjectID: 1, Name: "", Description: "", Priority: 10, Removed: false, CreatedAt: timeNow},
			},
		},
		{
			name:        "Same priority 2",
			goods:       startData,
			id:          5,
			projectId:   1,
			newPriority: 5,
			expectError: false,
			error:       nil,
			want: []models.Goods{
				{ID: 1, ProjectID: 1, Name: "", Description: "", Priority: 1, Removed: false, CreatedAt: timeNow},
				{ID: 2, ProjectID: 1, Name: "", Description: "", Priority: 2, Removed: false, CreatedAt: timeNow},
				{ID: 3, ProjectID: 1, Name: "", Description: "", Priority: 3, Removed: false, CreatedAt: timeNow},
				{ID: 4, ProjectID: 1, Name: "", Description: "", Priority: 4, Removed: false, CreatedAt: timeNow},
				{ID: 5, ProjectID: 1, Name: "", Description: "", Priority: 5, Removed: false, CreatedAt: timeNow},
				{ID: 6, ProjectID: 1, Name: "", Description: "", Priority: 6, Removed: false, CreatedAt: timeNow},
				{ID: 7, ProjectID: 1, Name: "", Description: "", Priority: 7, Removed: false, CreatedAt: timeNow},
				{ID: 8, ProjectID: 1, Name: "", Description: "", Priority: 8, Removed: false, CreatedAt: timeNow},
				{ID: 9, ProjectID: 1, Name: "", Description: "", Priority: 9, Removed: false, CreatedAt: timeNow},
				{ID: 10, ProjectID: 1, Name: "", Description: "", Priority: 10, Removed: false, CreatedAt: timeNow},
			},
		},
		{
			name:        "Good not found",
			goods:       startData,
			id:          11,
			projectId:   1,
			newPriority: 11,
			expectError: true,
			error:       errs.ErrGoodsNotFound,
			want: []models.Goods{
				{ID: 1, ProjectID: 1, Name: "", Description: "", Priority: 1, Removed: false, CreatedAt: timeNow},
				{ID: 2, ProjectID: 1, Name: "", Description: "", Priority: 2, Removed: false, CreatedAt: timeNow},
				{ID: 3, ProjectID: 1, Name: "", Description: "", Priority: 3, Removed: false, CreatedAt: timeNow},
				{ID: 4, ProjectID: 1, Name: "", Description: "", Priority: 4, Removed: false, CreatedAt: timeNow},
				{ID: 5, ProjectID: 1, Name: "", Description: "", Priority: 5, Removed: false, CreatedAt: timeNow},
				{ID: 6, ProjectID: 1, Name: "", Description: "", Priority: 6, Removed: false, CreatedAt: timeNow},
				{ID: 7, ProjectID: 1, Name: "", Description: "", Priority: 7, Removed: false, CreatedAt: timeNow},
				{ID: 8, ProjectID: 1, Name: "", Description: "", Priority: 8, Removed: false, CreatedAt: timeNow},
				{ID: 9, ProjectID: 1, Name: "", Description: "", Priority: 9, Removed: false, CreatedAt: timeNow},
				{ID: 10, ProjectID: 1, Name: "", Description: "", Priority: 10, Removed: false, CreatedAt: timeNow},
			},
		},
		{
			name:        "Test 1",
			goods:       startData,
			id:          5,
			projectId:   1,
			newPriority: 1,
			expectError: false,
			error:       nil,
			want: []models.Goods{
				{ID: 1, ProjectID: 1, Name: "", Description: "", Priority: 2, Removed: false, CreatedAt: timeNow},
				{ID: 2, ProjectID: 1, Name: "", Description: "", Priority: 3, Removed: false, CreatedAt: timeNow},
				{ID: 3, ProjectID: 1, Name: "", Description: "", Priority: 4, Removed: false, CreatedAt: timeNow},
				{ID: 4, ProjectID: 1, Name: "", Description: "", Priority: 5, Removed: false, CreatedAt: timeNow},
				{ID: 5, ProjectID: 1, Name: "", Description: "", Priority: 1, Removed: false, CreatedAt: timeNow},
				{ID: 6, ProjectID: 1, Name: "", Description: "", Priority: 6, Removed: false, CreatedAt: timeNow},
				{ID: 7, ProjectID: 1, Name: "", Description: "", Priority: 7, Removed: false, CreatedAt: timeNow},
				{ID: 8, ProjectID: 1, Name: "", Description: "", Priority: 8, Removed: false, CreatedAt: timeNow},
				{ID: 9, ProjectID: 1, Name: "", Description: "", Priority: 9, Removed: false, CreatedAt: timeNow},
				{ID: 10, ProjectID: 1, Name: "", Description: "", Priority: 10, Removed: false, CreatedAt: timeNow},
			},
		},
		{
			name:        "Test 2",
			goods:       startData,
			id:          5,
			projectId:   1,
			newPriority: 10,
			expectError: false,
			error:       nil,
			want: []models.Goods{
				{ID: 1, ProjectID: 1, Name: "", Description: "", Priority: 1, Removed: false, CreatedAt: timeNow},
				{ID: 2, ProjectID: 1, Name: "", Description: "", Priority: 2, Removed: false, CreatedAt: timeNow},
				{ID: 3, ProjectID: 1, Name: "", Description: "", Priority: 3, Removed: false, CreatedAt: timeNow},
				{ID: 4, ProjectID: 1, Name: "", Description: "", Priority: 4, Removed: false, CreatedAt: timeNow},
				{ID: 5, ProjectID: 1, Name: "", Description: "", Priority: 10, Removed: false, CreatedAt: timeNow},
				{ID: 6, ProjectID: 1, Name: "", Description: "", Priority: 5, Removed: false, CreatedAt: timeNow},
				{ID: 7, ProjectID: 1, Name: "", Description: "", Priority: 6, Removed: false, CreatedAt: timeNow},
				{ID: 8, ProjectID: 1, Name: "", Description: "", Priority: 7, Removed: false, CreatedAt: timeNow},
				{ID: 9, ProjectID: 1, Name: "", Description: "", Priority: 8, Removed: false, CreatedAt: timeNow},
				{ID: 10, ProjectID: 1, Name: "", Description: "", Priority: 9, Removed: false, CreatedAt: timeNow},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Reprioritize(tt.goods, tt.id, tt.projectId, tt.newPriority)
			if !tt.expectError {
				require.NoError(t, err)
			} else {
				if assert.Error(t, err) {
					assert.Equal(t, tt.error, err)
				}
			}

			assert.Equal(t, tt.want, result)
		})
	}
}
