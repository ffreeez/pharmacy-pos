package drughandler

import (
	"fmt"
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

	fmt.Println(drug)
	err := dh.DrugService.CreateDrug(&drug)
	if err != nil {
		response.InternalServerError(c, "Failed to create drug")
		return
	}

	response.Created(c, drug, "success")
}

// UpdateDrug 更新药品信息
func (dh *DrugHandler) UpdateDrug(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid drug ID")
		return
	}
	var drug drugmodel.Drug
	if err := c.ShouldBindJSON(&drug); err != nil {
		response.BadRequest(c, "Invalid input")
		return
	}

	err = dh.DrugService.UpdateDrug(&drug, uint(id))
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

	// SimplifiedUser 用于只返回必要的用户信息
	type SimplifiedDrug struct {
		ID           uint    `json:"id"`
		Name         string  `json:"name"`
		Description  string  `json:"description"`
		CategoryID   uint    `json:"categoryid"`
		CategoryName string  `json:"categoryname"`
		Price        float64 `json:"price"`
		Inventory    float64 `json:"inventory"`
	}

	simplifiedDrugs, err := dh.DrugService.GetAllDrugs()
	if err != nil {
		response.InternalServerError(c, "Failed to get all drugs")
		return
	}

	var drugResponse []SimplifiedDrug
	for _, drug := range simplifiedDrugs {
		drugResponse = append(drugResponse, SimplifiedDrug{
			ID:           drug.ID,
			Name:         drug.Name,
			Description:  drug.Description,
			CategoryID:   drug.CategoryID,
			CategoryName: drug.Category.Name,
			Price:        drug.Price,
			Inventory:    float64(drug.Inventory),
		})
	}
	response.OK(c, drugResponse, "success")
}

// GetCategoryByID 根据分类ID获取分类信息
func (dh *DrugHandler) GetCategoryByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid drug ID")
		return
	}

	categoryID := uint(id)
	category, err := dh.DrugService.GetCategoryByID(categoryID)
	if err != nil {
		response.InternalServerError(c, "Failed to get category")
		return
	}

	response.OK(c, category, "success")
}

// CreateCategory 创建新分类
func (dh *DrugHandler) CreateCategory(c *gin.Context) {

	type CategoryInput struct {
		Name string `json:"name"`
	}

	var input CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, "Invalid input")
		return
	}

	category := drugmodel.Category{
		Name: input.Name,
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

	type CategoryInput struct {
		Name string `json:"name"`
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid category ID")
		return
	}

	var input CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, "Invalid input")
		return
	}

	categoryID := uint(id)
	category, err := dh.DrugService.GetCategoryByID(categoryID)
	if err != nil {
		response.InternalServerError(c, "Failed to get category")
		return
	}

	category.Name = input.Name
	err = dh.DrugService.UpdateCategory(category, uint(id))
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
func (dh *DrugHandler) GetAllCategories(c *gin.Context) {
	type SimplifiedCategory struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	simplifiedCategories, err := dh.DrugService.GetAllCategories()
	if err != nil {
		response.InternalServerError(c, "Failed to get all categories")
		return
	}

	var categoryResponse []SimplifiedCategory
	for _, drug := range simplifiedCategories {
		categoryResponse = append(categoryResponse, SimplifiedCategory{
			ID:   drug.ID,
			Name: drug.Name,
		})
	}
	response.OK(c, categoryResponse, "success")
}

// GetCategoryByName 根据分类名获取分类
func (dh *DrugHandler) GetCategoryByName(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		response.BadRequest(c, "Name is required")
		return
	}

	category, err := dh.DrugService.GetCategoryByName(name)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.NotFound(c, "Category not found")
		} else {
			response.InternalServerError(c, "Failed to get category")
		}
		return
	}

	data := struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}{
		ID:   category.ID,
		Name: category.Name,
	}

	response.OK(c, data, "success")
}
