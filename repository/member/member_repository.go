package repository

import (
	"context"
	"pharmacy-pos/models"

	"gorm.io/gorm"
)

// MemberRepository 定义了操作会员表的接口
type MemberRepository interface {
	Create(ctx context.Context, member *models.Member) error
	Delete(ctx context.Context, memberID uint) error
	Update(ctx context.Context, member *models.Member) error
	FindByID(ctx context.Context, memberID uint) (*models.Member, error)
	FindAll(ctx context.Context) ([]*models.Member, error)
}

// MemberRepositoryImpl 实现了MemberRepository接口
type MemberRepositoryImpl struct {
	db *gorm.DB
}

// NewMemberRepository 创建一个新的MemberRepository实例
func NewMemberRepository(db *gorm.DB) MemberRepository {
	return &MemberRepositoryImpl{db}
}

// Create 添加一个新的会员到数据库
func (repo *MemberRepositoryImpl) Create(ctx context.Context, member *models.Member) error {
	return repo.db.WithContext(ctx).Create(member).Error
}

// Delete 根据MemberID删除一个会员
func (repo *MemberRepositoryImpl) Delete(ctx context.Context, memberID uint) error {
	return repo.db.WithContext(ctx).Delete(&models.Member{}, memberID).Error
}

// Update 更新会员信息
func (repo *MemberRepositoryImpl) Update(ctx context.Context, member *models.Member) error {
	return repo.db.WithContext(ctx).Save(member).Error
}

// FindByID 根据MemberID查找一个会员
func (repo *MemberRepositoryImpl) FindByID(ctx context.Context, memberID uint) (*models.Member, error) {
	var member models.Member
	err := repo.db.WithContext(ctx).First(&member, memberID).Error
	return &member, err
}

// FindAll 返回所有会员的列表
func (repo *MemberRepositoryImpl) FindAll(ctx context.Context) ([]*models.Member, error) {
	var members []*models.Member
	err := repo.db.WithContext(ctx).Find(&members).Error
	return members, err
}
