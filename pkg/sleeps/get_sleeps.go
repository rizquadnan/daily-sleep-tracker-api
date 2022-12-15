package sleeps

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/models"
	"gorm.io/gorm"
)

type Pagination struct {
	getTotalPage func(totalRows int64) int
	totalRows    int64
}

func SetupPagination(c *gin.Context, h handler) (func(db *gorm.DB) *gorm.DB, Pagination) {
	if c.Query("page") == "" || c.Query("page_size") == "" {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}, Pagination{totalRows: 0}
	}

	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 5
	}

	offset := (page - 1) * pageSize

	paginator := func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(pageSize)
	}

	return paginator, Pagination{getTotalPage: func(totalRows int64) int {
		return int(math.Ceil(float64(totalRows) / float64(pageSize)))
	}, totalRows: 0}
}

func (h handler) GetSleeps(c *gin.Context) {
	var sleeps []models.Sleep
	userId := c.Query("user")

	paginator, pagination := SetupPagination(c, h)

	if userId == "" {
		h.DB.Model(models.Sleep{}).Count(&pagination.totalRows)
		if result := h.DB.Scopes(paginator).Find(&sleeps); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}
	} else {
		userIdInt, err := strconv.Atoi(userId)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		}

		h.DB.Model(models.Sleep{}).Where(models.Sleep{UserID: uint(userIdInt)}).Count(&pagination.totalRows)
		if result := h.DB.Where(models.Sleep{UserID: uint(userIdInt)}).Scopes(paginator).Find(&sleeps); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"totalRows": pagination.totalRows,
		"totalPage": pagination.getTotalPage(pagination.totalRows),
		"rows":      SleepsToSleepsResponse(sleeps)},
	)
}
