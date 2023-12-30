package repository

import (
	"context"
	"pharmacy-pos/models"

	"gorm.io/gorm"
)

// InventoryRepository 定义了操作库存表的接口
type InventoryRepository interface {
	Create(ctx context.Context, inventory *models.Inventory) error
	Delete(ctx context.Context, inventoryID uint) error
	Update(ctx context.Context, inventory *models.Inventory) error
	FindByID(ctx context.Context, inventoryID uint) (*models.Inventory, error)
	FindAll(ctx context.Context) ([]*models.Inventory, error)
}

// InventoryRepositoryImpl 实现了InventoryRepository接口
type InventoryRepositoryImpl struct {
	db *gorm.DB
}

// NewInventoryRepository 创建一个新的InventoryRepository实例
func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	return &InventoryRepositoryImpl{db}
}

// Create 添加一个新的库存记录到数据库
func (repo *InventoryRepositoryImpl) Create(ctx context.Context, inventory *models.Inventory) error {
	return repo.db.WithContext(ctx).Create(inventory).Error
}

// Delete 根据InventoryID删除一个库存记录
func (repo *InventoryRepositoryImpl) Delete(ctx context.Context, inventoryID uint) error {
	return repo.db.WithContext(ctx).Delete(&models.Inventory{}, inventoryID).Error
}

// Update 更新库存信息
func (repo *InventoryRepositoryImpl) Update(ctx context.Context, inventory *models.Inventory) error {
	return repo.db.WithContext(ctx).Save(inventory).Error
}

// FindByID 根据InventoryID查找一个库存记录
func (repo *InventoryRepositoryImpl) FindByID(ctx context.Context, inventoryID uint) (*models.Inventory, error) {
	var inventory models.Inventory
	err := repo.db.WithContext(ctx).First(&inventory, inventoryID).Error
	return &inventory, err
}

// FindAll 返回所有库存记录的列表
func (repo *InventoryRepositoryImpl) FindAll(ctx context.Context) ([]*models.Inventory, error) {
	var inventories []*models.Inventory
	err := repo.db.WithContext(ctx).Find(&inventories).Error
	return inventories, err
}
