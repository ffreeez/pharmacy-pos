package memberhandler

import (
	"strconv"

	membermodel "pharmacy-pos/pkg/db/models/member"
	memberservice "pharmacy-pos/pkg/service/member"
	"pharmacy-pos/pkg/util/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// MemberHandler 处理会员相关的 HTTP 请求
type MemberHandler struct {
	MemberService *memberservice.MemberService
}

// NewMemberHandler 创建一个新的 MemberHandler 实例
func NewMemberHandler(db *gorm.DB) *MemberHandler {
	return &MemberHandler{
		MemberService: memberservice.NewMemberService(db),
	}
}

// CreateMember 创建新会员
func (mh *MemberHandler) CreateMember(c *gin.Context) {
	var member membermodel.Member
	if err := c.ShouldBindJSON(&member); err != nil {
		response.BadRequest(c, "Invalid input")
		return
	}

	err := mh.MemberService.CreateMember(&member)
	if err != nil {
		response.InternalServerError(c, "Failed to create member")
		return
	}

	response.Created(c, member, "success")
}

// GetMemberByID 根据ID获取会员
func (mh *MemberHandler) GetMemberByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid member ID")
		return
	}

	memberID := uint(id)
	member, err := mh.MemberService.GetMemberByID(memberID)
	if err != nil {
		response.InternalServerError(c, "Failed to get member")
		return
	}

	response.OK(c, member, "success")
}

// UpdateMember 更新会员信息
func (mh *MemberHandler) UpdateMember(c *gin.Context) {
	var member membermodel.Member
	if err := c.ShouldBindJSON(&member); err != nil {
		response.BadRequest(c, "Invalid input")
		return
	}

	err := mh.MemberService.UpdateMember(&member)
	if err != nil {
		response.InternalServerError(c, "Failed to update member")
		return
	}

	response.OK(c, member, "success")
}

// DeleteMemberByID 根据ID删除会员
func (mh *MemberHandler) DeleteMemberByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid member ID")
		return
	}

	memberID := uint(id)
	err = mh.MemberService.DeleteMemberByID(memberID)
	if err != nil {
		response.InternalServerError(c, "Failed to delete member")
		return
	}

	response.OK(c, gin.H{"message": "Member deleted successfully"}, "success")
}

// GetAllMembers 获取所有会员信息
func (mh *MemberHandler) GetAllMembers(c *gin.Context) {
	members, err := mh.MemberService.GetAllMembers()
	if err != nil {
		response.InternalServerError(c, "Failed to get all members")
		return
	}

	response.OK(c, members, "success")
}

// AddCouponToMember 为会员添加优惠券
func (mh *MemberHandler) AddCouponToMember(c *gin.Context) {
	var coupon membermodel.Coupon
	if err := c.ShouldBindJSON(&coupon); err != nil {
		response.BadRequest(c, "Invalid input")
		return
	}

	err := mh.MemberService.AddCouponToMember(&coupon)
	if err != nil {
		response.InternalServerError(c, "Failed to add coupon to member")
		return
	}

	response.Created(c, coupon, "success")
}

// UseCoupon 使用优惠券
func (mh *MemberHandler) UseCoupon(c *gin.Context) {
	couponID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid coupon ID")
		return
	}

	err = mh.MemberService.UseCoupon(uint(couponID))
	if err != nil {
		response.InternalServerError(c, "Failed to use coupon")
		return
	}

	response.OK(c, gin.H{"message": "Coupon used successfully"}, "success")
}

// DeleteCoupon 删除优惠券
func (mh *MemberHandler) DeleteCoupon(c *gin.Context) {
	couponID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid coupon ID")
		return
	}

	err = mh.MemberService.DeleteCoupon(uint(couponID))
	if err != nil {
		response.InternalServerError(c, "Failed to delete coupon")
		return
	}

	response.OK(c, gin.H{"message": "Coupon deleted successfully"}, "success")
}

// GetCouponsByMemberID 获取会员的所有优惠券
func (mh *MemberHandler) GetCouponsByMemberID(c *gin.Context) {
	memberID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid member ID")
		return
	}

	coupons, err := mh.MemberService.GetCouponsByMemberID(uint(memberID))
	if err != nil {
		response.InternalServerError(c, "Failed to get coupons")
		return
	}

	response.OK(c, coupons, "success")
}
