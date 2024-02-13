package handler

import (
	"context"
	"net/http"
	"strconv"
	"test/api/models"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateIncome godoc
// @Router       /income [POST]
// @Summary      Creates a new income
// @Description  create a new income
// @Tags         income
// @Accept       json
// @Produce      json
// @Success      201  {object}  models.Income
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateIncome(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := h.services.Income().Create(ctx)
	if err != nil {
		handleResponse(c, "error while creating income", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusCreated, resp)
}

// GetIncome godoc
// @Router       /income/{id} [GET]
// @Summary      Get income by id
// @Description  get income by id
// @Tags         income
// @Accept       json
// @Produce      json
// @Param        id path string true "income_id"
// @Success      201  {object}  models.Income
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetIncome(c *gin.Context) {
	uid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := h.services.Income().Get(ctx, models.PrimaryKey{ID: uid})
	if err != nil {
		handleResponse(c, "error is while getting income by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, resp)
}

// GetIncomeList godoc
// @Router       /incomes [GET]
// @Summary      Get incomes list
// @Description  get incomes list
// @Tags         income
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      201  {object}  models.IncomesResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetIncomeList(c *gin.Context) {
	var (
		page, limit int
		search      string
		err         error
	)

	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, "error is while converting pageStr", http.StatusBadRequest, err)
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c, "error is while converting limitStr", http.StatusBadRequest, err)
		return
	}

	search = c.Query("search")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := h.services.Income().GetList(ctx, models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(c, "error is while getting incomes list", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "", http.StatusOK, resp)
}

// DeleteIncome godoc
// @Router       /income/{id} [Delete]
// @Summary      Delete income
// @Description  delete income
// @Tags         income
// @Accept       json
// @Produce      json
// @Param        id path string true "income_id"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteIncome(c *gin.Context) {
	uid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := h.services.Income().Delete(ctx, models.PrimaryKey{ID: uid}); err != nil {
		handleResponse(c, "error is while deleting basket", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "", http.StatusOK, nil)
}
