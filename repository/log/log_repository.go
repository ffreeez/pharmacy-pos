package repository

import (
	"context"
	"pharmacy-pos/models"

	"gorm.io/gorm"
)

// LogRepository 定义了操作日志表的接口
type LogRepository interface {
	Create(ctx context.Context, log *models.Log) error
	Delete(ctx context.Context, logID uint) error
	Update(ctx context.Context, log *models.Log) error
	FindByID(ctx context.Context, logID uint) (*models.Log, error)
	FindAll(ctx context.Context) ([]*models.Log, error)
}

// LogRepositoryImpl 实现了LogRepository接口
type LogRepositoryImpl struct {
	db *gorm.DB
}

// NewLogRepository 创建一个新的LogRepository实例
func NewLogRepository(db *gorm.DB) LogRepository {
	return &LogRepositoryImpl{db}
}

// Create 添加一个新的日志记录到数据库
func (repo *LogRepositoryImpl) Create(ctx context.Context, log *models.Log) error {
	return repo.db.WithContext(ctx).Create(log).Error
}

// Delete 根据LogID删除一个日志记录
func (repo *LogRepositoryImpl) Delete(ctx context.Context, logID uint) error {
	return repo.db.WithContext(ctx).Delete(&models.Log{}, logID).Error
}

// Update 更新日志信息
func (repo *LogRepositoryImpl) Update(ctx context.Context, log *models.Log) error {
	return repo.db.WithContext(ctx).Save(log).Error
}

// FindByID 根据LogID查找一个日志记录
func (repo *LogRepositoryImpl) FindByID(ctx context.Context, logID uint) (*models.Log, error) {
	var log models.Log
	err := repo.db.WithContext(ctx).First(&log, logID).Error
	return &log, err
}

// FindAll 返回所有日志记录的列表
func (repo *LogRepositoryImpl) FindAll(ctx context.Context) ([]*models.Log, error) {
	var logs []*models.Log
	err := repo.db.WithContext(ctx).Find(&logs).Error
	return logs, err
}
