package memberservice

import (
	membermodel "pharmacy-pos/pkg/db/models/member"
	memberrepo "pharmacy-pos/pkg/db/repository/member"

	"gorm.io/gorm"
)

// MemberService 提供会员相关的服务
type MemberService struct {
	DB *gorm.DB
}

// NewMemberService 创建一个新的 MemberService 实例
func NewMemberService(db *gorm.DB) *MemberService {
	return &MemberService{DB: db}
}

// CreateMember 创建新会员
func (ms *MemberService) CreateMember(member *membermodel.Member) error {
	return memberrepo.CreateMember(ms.DB, member)
}

// GetMemberByID 根据会员ID获取会员信息
func (ms *MemberService) GetMemberByID(id uint) (*membermodel.Member, error) {
	return memberrepo.GetMemberByID(ms.DB, id)
}

// UpdateMember 更新会员信息
func (ms *MemberService) UpdateMember(member *membermodel.Member) error {
	return memberrepo.UpdateMember(ms.DB, member)
}

// DeleteMemberByID 根据ID删除会员
func (ms *MemberService) DeleteMemberByID(id uint) error {
	return memberrepo.DeleteMemberByID(ms.DB, id)
}

// GetAllMembers 获取所有会员信息
func (ms *MemberService) GetAllMembers() ([]membermodel.Member, error) {
	return memberrepo.GetAllMembers(ms.DB)
}

// CreateCoupon 为会员创建优惠券
func (ms *MemberService) CreateCoupon(coupon *membermodel.Coupon) error {
	return memberrepo.CreateCoupon(ms.DB, coupon)
}

// GetCouponByID 根据优惠券ID获取优惠券
func (ms *MemberService) GetCouponByID(couponID uint) (*membermodel.Coupon, error) {
	return memberrepo.GetCouponByID(ms.DB, couponID)
}

// UpdateCoupon 更新优惠券信息
func (ms *MemberService) UpdateCoupon(coupon *membermodel.Coupon) error {
	return memberrepo.UpdateCoupon(ms.DB, coupon)
}

// DeleteCoupon 删除优惠券
func (ms *MemberService) DeleteCoupon(couponID uint) error {
	return memberrepo.DeleteCoupon(ms.DB, couponID)
}

// GetAllCoupons 获取所有优惠券
func (ms *MemberService) GetAllCoupons() ([]membermodel.Coupon, error) {
	return memberrepo.GetAllCoupons(ms.DB)
}
