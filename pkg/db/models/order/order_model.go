package ordermodels

import (
	membermodel "pharmacy-pos/pkg/db/models/member"

	"gorm.io/gorm"
)

// Order 表示订单信息
type Order struct {
	gorm.Model
	CustomerName string              `gorm:"type:varchar(100);not null" json:"customerName"` // 消费者名字
	MemberID     *uint               `json:"memberId,omitempty"`                             // 会员ID，如果是会员的话
	Member       *membermodel.Member `gorm:"foreignKey:MemberID" json:"member,omitempty"`    // 关联的会员信息
	OrderItems   []OrderItem         `gorm:"foreignKey:OrderID" json:"orderItems"`           // 订单项
	TotalAmount  float64             `json:"totalAmount"`                                    // 订单总金额
	Discount     float64             `json:"discount"`                                       // 应用的总折扣金额
	FinalAmount  float64             `json:"finalAmount"`                                    // 折扣后的最终金额
}

// OrderItem 表示订单中的单个药品项
type OrderItem struct {
	gorm.Model
	OrderID   uint    `gorm:"index" json:"orderId"` // 关联的订单ID
	DrugID    uint    `json:"drugId"`               // 关联的药品ID
	Quantity  int     `json:"quantity"`             // 购买数量
	UnitPrice float64 `json:"unitPrice"`            // 当时的药品单价
}
