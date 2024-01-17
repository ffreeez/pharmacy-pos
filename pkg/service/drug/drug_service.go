package drugservice

import (
	drugmodel "pharmacy-pos/pkg/db/models/drug"
	drugrepo "pharmacy-pos/pkg/db/repository/drug"

	"gorm.io/gorm"
)

// DrugService 提供药品相关的服务
type DrugService struct {
	DB *gorm.DB
}

// NewDrugService 创建一个新的 DrugService 实例
func NewDrugService(db *gorm.DB) *DrugService {
	return &DrugService{DB: db}
}

// GetDrugByID 根据药品ID获取药品信息
func (ds *DrugService) GetDrugByID(id uint) (*drugmodel.Drug, error) {
	return drugrepo.GetDrugByID(ds.DB, id)
}

// CreateDrug 创建新的药品记录
func (ds *DrugService) CreateDrug(drug *drugmodel.Drug) error {
	return drugrepo.CreateDrug(ds.DB, drug)
}

// UpdateDrug 更新药品信息
func (ds *DrugService) UpdateDrug(drug *drugmodel.Drug, id uint) error {
	return drugrepo.UpdateDrug(ds.DB, drug, id)
}

// DeleteDrugByID 根据ID删除药品记录
func (ds *DrugService) DeleteDrugByID(id uint) error {
	return drugrepo.DeleteDrugByID(ds.DB, id)
}

// GetAllDrugs 获取所有药品的信息
func (ds *DrugService) GetAllDrugs() ([]drugmodel.Drug, error) {
	return drugrepo.GetAllDrugs(ds.DB)
}

// GetCategoryByID 根据药品ID获取分类信息
func (ds *DrugService) GetCategoryByID(id uint) (*drugmodel.Category, error) {
	return drugrepo.GetCategoryByID(ds.DB, id)
}

// CreateCategory 创建新的分类记录
func (ds *DrugService) CreateCategory(category *drugmodel.Category) error {
	return drugrepo.CreateCategory(ds.DB, category)
}

// UpdateCategory 更新分类信息
func (ds *DrugService) UpdateCategory(drug *drugmodel.Category, id uint) error {
	return drugrepo.UpdateCategory(ds.DB, drug, id)
}

// DeleteCategoryByID 根据ID删除分类记录
func (ds *DrugService) DeleteCategoryByID(id uint) error {
	return drugrepo.DeleteCategoryByID(ds.DB, id)
}

// GetAllCategory 获取所有分类的信息
func (ds *DrugService) GetAllCategories() ([]drugmodel.Category, error) {
	return drugrepo.GetAllCategorys(ds.DB)
}
