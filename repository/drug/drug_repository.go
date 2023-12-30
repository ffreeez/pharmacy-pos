package repository

import (
	"context"
	"pharmacy-pos/models"

	"gorm.io/gorm"
)

// DrugRepository 定义了操作药品表的接口
type DrugRepository interface {
	Create(ctx context.Context, drug *models.Drug) error
	Delete(ctx context.Context, drugID uint) error
	Update(ctx context.Context, drug *models.Drug) error
	FindByID(ctx context.Context, drugID uint) (*models.Drug, error)
	FindAll(ctx context.Context) ([]*models.Drug, error)
}

// DrugRepositoryImpl 实现了DrugRepository接口
type DrugRepositoryImpl struct {
	db *gorm.DB
}

// NewDrugRepository 创建一个新的DrugRepository实例
func NewDrugRepository(db *gorm.DB) DrugRepository {
	return &DrugRepositoryImpl{db}
}

// Create 添加一个新的药品到数据库
func (repo *DrugRepositoryImpl) Create(ctx context.Context, drug *models.Drug) error {
	return repo.db.WithContext(ctx).Create(drug).Error
}

// Delete 根据drugID删除一个药品
func (repo *DrugRepositoryImpl) Delete(ctx context.Context, drugID uint) error {
	return repo.db.WithContext(ctx).Delete(&models.Drug{}, drugID).Error
}

// Update 更新药品信息
func (repo *DrugRepositoryImpl) Update(ctx context.Context, drug *models.Drug) error {
	return repo.db.WithContext(ctx).Save(drug).Error
}

// FindByID 根据drugID查找一个药品
func (repo *DrugRepositoryImpl) FindByID(ctx context.Context, drugID uint) (*models.Drug, error) {
	var drug models.Drug
	err := repo.db.WithContext(ctx).First(&drug, drugID).Error
	return &drug, err
}

// FindAll 返回所有药品的列表
func (repo *DrugRepositoryImpl) FindAll(ctx context.Context) ([]*models.Drug, error) {
	var drugs []*models.Drug
	err := repo.db.WithContext(ctx).Find(&drugs).Error
	return drugs, err
}
