package gin

import (
	"net/http"

	"example.com/modules/category/categorybiz"
	"example.com/modules/category/categorymodel"
	"example.com/modules/category/categorystorage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data categorymodel.CategoryCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := categorystorage.NewSQLStore(db)
		biz := categorybiz.NewCreateCategoryBiz(store)

		if err := biz.CreateCategory(c.Request.Context(), &data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, data)
	}
}
