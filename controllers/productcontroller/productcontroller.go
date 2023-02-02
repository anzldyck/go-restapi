package productcontroller

import (
	"net/http"
	"github.com/anzldyck/go-restapi/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": products})

}
func Show(c *gin.Context) {
	var product models.Product
	product_id := c.Param("id")

	if err := models.DB.First(&product, product_id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"}
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"products": product})
}
func Create(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	} 

	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"products": product})
}
func Update(c *gin.Context) {
	var product models.Product
	product_id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	} 

	if models.DB.Model(&product).Where("product_id = ?", product_id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "tidak dapat mengupdate product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil diupdate"})
}
func Delete(c *gin.Context) {
	var product models.Product
	var input struct {
		product_id json.Number
	} 

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	} 

	product_id, _ := input.product_id.Int64()
	if models.DB.Delete(&product, product_id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat menghapus"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil dihapus"})
}