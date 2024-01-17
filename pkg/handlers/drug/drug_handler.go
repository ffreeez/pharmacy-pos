package drughandler

import (
	"strconv"

	drugmodel "pharmacy-pos/pkg/db/models/drug"
	drugservice "pharmacy-pos/pkg/service/drug"
	"pharmacy-pos/pkg/util/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DrugHandler 处理药品相关的 HTTP 请求
type DrugHandler struct {
	DrugService *drugservice.DrugService
}

// NewDrugHandler 创建一个新的 DrugHandler 实例
func NewDrugHandler(db *gorm.DB) *DrugHandler {
	return &DrugHandler{
		DrugService: drugservice.NewDrugService(db),
	}
}

// GetDrugByID 根据药品ID获取药品信息
func (dh *DrugHandler) GetDrugByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid drug ID")
		return
	}

	drugID := uint(id)
	drug, err := dh.DrugService.GetDrugByID(drugID)
	if err != nil {
		response.InternalServerError(c, "Failed to get drug")
		return
	}

	response.OK(c, drug, "success")
}

// CreateDrug 创建新药品
func (dh *DrugHandler) CreateDrug(c *gin.Context) {
	var drug drugmodel.Drug
	if err := c.ShouldBindJSON(&drug); err != nil {
		response.BadRequest(c, "Invalid input")
		return
	}

	err := dh.DrugService.CreateDrug(&drug)
	if err != nil {
		response.InternalServerError(c, "Failed to create drug")
		return
	}

	response.Created(c, drug, "success")
}

// UpdateDrug 更新药品信息
func (dh *DrugHandler) UpdateDrug(c *gin.Context) {
	var drug drugmodel.Drug
	if err := c.ShouldBindJSON(&drug); err != nil {
		response.BadRequest(c, "Invalid input")
		return
	}

	err := dh.DrugService.UpdateDrug(&drug)
	if err != nil {
		response.InternalServerError(c, "Failed to update drug")
		return
	}

	response.OK(c, drug, "success")
}

// DeleteDrugByID 根据ID删除药品
func (dh *DrugHandler) DeleteDrugByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid drug ID")
		return
	}

	drugID := uint(id)
	err = dh.DrugService.DeleteDrugByID(drugID)
	if err != nil {
		response.InternalServerError(c, "Failed to delete drug")
		return
	}

	response.OK(c, gin.H{"message": "Drug deleted successfully"}, "success")
}

// GetAllDrugs 获取所有药品信息
func (dh *DrugHandler) GetAllDrugs(c *gin.Context) {
	drugs, err := dh.DrugService.GetAllDrugs()
	if err != nil {
		response.InternalServerError(c, "Failed to get all drugs")
		return
	}

	response.OK(c, drugs, "success")
}

// GetCategoryByID 根据分类ID获取分类信息
func (dh *DrugHandler) GetCategoryByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid drug ID")
		return
	}

	drugID := uint(id)
	drug, err := dh.DrugService.GetCategoryByID(drugID)
	if err != nil {
		response.InternalServerError(c, "Failed to get category")
		return
	}

	response.OK(c, drug, "success")
}

// CreateCategory 创建新分类
func (dh *DrugHandler) CreateCategory(c *gin.Context) {
	var category drugmodel.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		response.BadRequest(c, "Invalid input")
		return
	}

	err := dh.DrugService.CreateCategory(&category)
	if err != nil {
		response.InternalServerError(c, "Failed to create category")
		return
	}

	response.Created(c, category, "success")
}

// UpdateCategory 更新分类信息
func (dh *DrugHandler) UpdateCategory(c *gin.Context) {
	var category drugmodel.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		response.BadRequest(c, "Invalid input")
		return
	}

	err := dh.DrugService.UpdateCategory(&category)
	if err != nil {
		response.InternalServerError(c, "Failed to update category")
		return
	}

	response.OK(c, category, "success")
}

// DeleteCategoryByID 根据ID删除分类
func (dh *DrugHandler) DeleteCategoryByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid drug ID")
		return
	}

	drugID := uint(id)
	err = dh.DrugService.DeleteCategoryByID(drugID)
	if err != nil {
		response.InternalServerError(c, "Failed to delete category")
		return
	}

	response.OK(c, gin.H{"message": "Drug deleted successfully"}, "success")
}

// GetAllCategory 获取所有分类信息
func (dh *DrugHandler) GetAllCategorys(c *gin.Context) {
	drugs, err := dh.DrugService.GetAllCategories()
	if err != nil {
		response.InternalServerError(c, "Failed to get all categorys")
		return
	}

	response.OK(c, drugs, "success")
}
