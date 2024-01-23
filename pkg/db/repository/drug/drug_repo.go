package drugrepo

import (
	"errors"
	drugmodel "pharmacy-pos/pkg/db/models/drug" // 更正导入路径为药品模型
	logger "pharmacy-pos/pkg/util/logger"

	"pharmacy-pos/pkg/util/e"

	"gorm.io/gorm"
)

var logs = logger.GetLogger()

// GetDrugByID 根据药品ID获取药品信息
func GetDrugByID(db *gorm.DB, id uint) (*drugmodel.Drug, error) {
	drug := &drugmodel.Drug{}
	result := db.First(drug, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		logs.Errorf("根据药品ID获取药品信息失败, ID: %d", id)
		return nil, e.ErrNotFound
	}
	if result.Error != nil {
		logs.Errorf("查询药品时发生错误, ID: %d, error: %v", id, result.Error)
		return nil, result.Error
	}
	logs.Infof("根据药品ID获取药品信息成功, ID: %d", id)
	return drug, nil
}

// CreateDrug 创建新的药品记录
func CreateDrug(db *gorm.DB, drug *drugmodel.Drug) error {
	result := db.Create(drug)
	if result.Error != nil {
		logs.Errorf("创建新药品记录失败, Name: %s, error: %v", drug.Name, result.Error)
		return result.Error
	}
	logs.Infof("创建新药品记录成功, Name: %s", drug.Name)
	return nil
}

// UpdateDrug 更新药品信息
func UpdateDrug(db *gorm.DB, drug *drugmodel.Drug, id uint) error {
	var existDrug drugmodel.Drug
	if err := db.First(&existDrug, id).Error; err != nil {
		logs.Errorf("未找到药品: %d, error: %v", id, err)
		return err
	}

	result := db.Save(drug)
	if result.Error != nil {
		logs.Errorf("更新药品信息失败, ID: %d, error: %v", drug.ID, result.Error)
		return result.Error
	}
	logs.Infof("更新药品信息成功, ID: %d", drug.ID)
	return nil
}

// DeleteDrugByID 根据ID删除药品记录
func DeleteDrugByID(db *gorm.DB, id uint) error {
	result := db.Delete(&drugmodel.Drug{}, id)
	if result.Error != nil {
		logs.Errorf("根据ID删除药品记录失败, ID: %d, error: %v", id, result.Error)
		return result.Error
	}
	logs.Infof("根据ID删除药品记录成功, ID: %d", id)
	return nil
}

// GetAllDrugs 获取所有药品的信息
func GetAllDrugs(db *gorm.DB) ([]drugmodel.Drug, error) {
	var drugs []drugmodel.Drug

	result := db.Preload("Category").Find(&drugs)

	if result.Error != nil {
		logs.Errorf("获取所有药品信息失败: %v", result.Error)
		return nil, result.Error
	}
	logs.Infof("获取所有药品信息成功")
	return drugs, nil
}

func GetDrugByDrugName(db *gorm.DB, drugname string) (*drugmodel.Drug, error) {
	drug := &drugmodel.Drug{}
	result := db.Where("name = ?", drugname).Preload("Category").First(drug)
	if result.Error != nil {
		logs.Errorf("根据药品名查找药品失败, drugname: %s", drugname)
		return nil, result.Error
	}
	logs.Infof("根据药品名查找药品成功, drugname: %s", drugname)
	return drug, nil
}

// GetCategoryByID 根据分类ID获取分类信息
func GetCategoryByID(db *gorm.DB, id uint) (*drugmodel.Category, error) {
	category := &drugmodel.Category{}
	result := db.First(category, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		logs.Errorf("根据分类ID获取分类信息失败, ID: %d", id)
		return nil, e.ErrNotFound
	}
	if result.Error != nil {
		logs.Errorf("查询分类时发生错误, ID: %d, error: %v", id, result.Error)
		return nil, result.Error
	}
	logs.Infof("根据分类ID获取分类信息成功, ID: %d", id)
	return category, nil
}

// CreateCategory 创建新的分类记录
func CreateCategory(db *gorm.DB, category *drugmodel.Category) error {
	result := db.Create(category)
	if result.Error != nil {
		logs.Errorf("创建新分类记录失败, Name: %s, error: %v", category.Name, result.Error)
		return result.Error
	}
	logs.Infof("创建新分类记录成功, Name: %s", category.Name)
	return nil
}

// UpdateCategory 更新分类信息
func UpdateCategory(db *gorm.DB, category *drugmodel.Category, id uint) error {
	var existCategory drugmodel.Category
	if err := db.First(&existCategory, id).Error; err != nil {
		logs.Errorf("未找到分类: %d, error: %v", id, err)
		return err
	}

	result := db.Save(category)
	if result.Error != nil {
		logs.Errorf("更新分类信息失败, ID: %d, error: %v", category.ID, result.Error)
		return result.Error
	}
	logs.Infof("更新分类信息成功, ID: %d", category.ID)
	return nil
}

// DeleteCategoryByID 根据ID删除分类记录
func DeleteCategoryByID(db *gorm.DB, id uint) error {
	result := db.Delete(&drugmodel.Category{}, id)
	if result.Error != nil {
		logs.Errorf("根据ID删除分类记录失败, ID: %d, error: %v", id, result.Error)
		return result.Error
	}
	logs.Infof("根据ID删除分类记录成功, ID: %d", id)
	return nil
}

// GetAllCategorys 获取所有分类的信息
func GetAllCategorys(db *gorm.DB) ([]drugmodel.Category, error) {
	var category []drugmodel.Category
	result := db.Preload("Drugs").Find(&category)
	if result.Error != nil {
		logs.Errorf("获取所有分类信息失败: %v", result.Error)
		return nil, result.Error
	}
	logs.Infof("获取所有分类信息成功")
	return category, nil
}

// GetCategoryByName 根据分类名获取分类信息
func GetCategoryByName(db *gorm.DB, name string) (*drugmodel.Category, error) {
	category := &drugmodel.Category{}
	result := db.Where("name = ?", name).First(category)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		logs.Errorf("根据分类名获取分类信息失败, Name: %s", name)
		return nil, e.ErrNotFound
	}
	if result.Error != nil {
		logs.Errorf("查询分类时发生错误, Name: %s, error: %v", name, result.Error)
		return nil, result.Error
	}
	logs.Infof("根据分类名获取分类信息成功, Name: %s", name)
	return category, nil
}
