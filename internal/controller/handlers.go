package controller

import (
	"errors"
	"github.com/MukizuL/hezzl-test/internal/dto"
	"github.com/MukizuL/hezzl-test/internal/errs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (c *Controller) Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{})
}

func (c *Controller) CreateGoods(ctx *gin.Context) {
	rawProjectId := ctx.Query("projectId")
	projectId, err := strconv.Atoi(rawProjectId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "projectId must be an integer",
			"details": err.Error(),
		})
		return
	}

	if projectId <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "projectId must be greater than zero",
			"details": "",
		})
		return
	}

	var data dto.CreateGoodsRequest
	err = ctx.BindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Incorrect body",
			"details": err.Error(),
		})
		return
	}

	goods, err := c.services.CreateGoods(ctx.Request.Context(), projectId, data.Name)
	if err != nil {
		if errors.Is(err, errs.ErrProjectNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": "Project with this ID does not exist",
				"details": err.Error(),
			})
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Internal server error",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, goods)
}

func (c *Controller) UpdateGoods(ctx *gin.Context) {
	rawId := ctx.Query("id")
	rawProjectId := ctx.Query("projectId")

	id, err := strconv.Atoi(rawId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "id must be an integer",
			"details": err.Error(),
		})
		return
	}

	projectId, err := strconv.Atoi(rawProjectId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "projectId must be an integer",
			"details": err.Error(),
		})
		return
	}

	if id <= 0 || projectId <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "id and projectId must be greater than zero",
			"details": "",
		})
		return
	}

	var data dto.UpdateGoodsRequest
	err = ctx.BindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Incorrect body",
			"details": err.Error(),
		})
		return
	}

	goods, err := c.services.UpdateGoods(ctx.Request.Context(), id, projectId, data.Name, data.Description)
	if err != nil {
		if errors.Is(err, errs.ErrGoodsNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"code":    3,
				"message": "errors.common.notFound",
				"details": "",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Internal server error",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, goods)
}

func (c *Controller) RemoveGoods(ctx *gin.Context) {
	rawId := ctx.Query("id")
	rawProjectId := ctx.Query("projectId")

	id, err := strconv.Atoi(rawId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "id must be an integer",
			"details": err.Error(),
		})
		return
	}

	projectId, err := strconv.Atoi(rawProjectId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "projectId must be an integer",
			"details": err.Error(),
		})
		return
	}

	if id <= 0 || projectId <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "id and projectId must be greater than zero",
			"details": "",
		})
		return
	}

	resp, err := c.services.RemoveGoods(ctx.Request.Context(), id, projectId)
	if err != nil {
		if errors.Is(err, errs.ErrGoodsNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"code":    3,
				"message": "errors.common.notFound",
				"details": "",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Internal server error",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) GetGoods(ctx *gin.Context) {
	ctx.JSON(200, gin.H{})
}

func (c *Controller) ReprioritizeGoods(ctx *gin.Context) {
	ctx.JSON(200, gin.H{})
}
