package dto

import "time"

type CreateGoodsRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateGoodsResponse struct {
	ID          int       `json:"id"`
	ProjectID   int       `json:"project_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	Removed     bool      `json:"removed"`
	CreatedAt   time.Time `json:"created_at"`
}

type UpdateGoodsRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type RemoveGoodsResponse struct {
	ID        int  `json:"id"`
	ProjectID int  `json:"project_id"`
	Removed   bool `json:"removed"`
}

type ReprioritizeRequest struct {
	Priority int `json:"newPriority" binding:"required"`
}

type ReprioritizeResponse struct {
	ID       int `json:"id"`
	Priority int `json:"priority"`
}
