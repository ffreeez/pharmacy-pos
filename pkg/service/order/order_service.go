package orderservice

import (
	ordermodels "pharmacy-pos/pkg/db/models/order"
	orderrepo "pharmacy-pos/pkg/db/repository/order"

	"gorm.io/gorm"
)

// OrderService 提供订单相关的服务
type OrderService struct {
	DB *gorm.DB
}

// NewOrderService 创建一个新的 OrderService 实例
func NewOrderService(db *gorm.DB) *OrderService {
	return &OrderService{DB: db}
}

// CreateOrder 创建新订单
func (os *OrderService) CreateOrder(order *ordermodels.Order) error {
	return orderrepo.CreateOrder(os.DB, order)
}

// GetOrderByID 根据订单ID获取订单信息
func (os *OrderService) GetOrderByID(id uint) (*ordermodels.Order, error) {
	return orderrepo.GetOrderByID(os.DB, id)
}

// UpdateOrder 更新订单信息
func (os *OrderService) UpdateOrder(order *ordermodels.Order) error {
	return orderrepo.UpdateOrder(os.DB, order)
}

// DeleteOrderByID 根据ID删除订单
func (os *OrderService) DeleteOrderByID(id uint) error {
	return orderrepo.DeleteOrderByID(os.DB, id)
}

// GetAllOrders 获取所有订单信息
func (os *OrderService) GetAllOrders() ([]ordermodels.Order, error) {
	return orderrepo.GetAllOrders(os.DB)
}

// CreateOrderItem 创建新订单项
func (os *OrderService) CreateOrderItem(orderItem *ordermodels.OrderItem) error {
	return orderrepo.CreateOrderItem(os.DB, orderItem)
}

// GetOrderItemByID 根据订单项ID获取订单项
func (os *OrderService) GetOrderItemByID(id uint) (*ordermodels.OrderItem, error) {
	return orderrepo.GetOrderItemByID(os.DB, id)
}

// UpdateOrderItem 更新订单项
func (os *OrderService) UpdateOrderItem(orderItem *ordermodels.OrderItem) error {
	return orderrepo.UpdateOrderItem(os.DB, orderItem)
}

// DeleteOrderItemByID 根据ID删除订单项
func (os *OrderService) DeleteOrderItemByID(id uint) error {
	return orderrepo.DeleteOrderItemByID(os.DB, id)
}

// GetAllOrderItems 获取所有订单项
func (os *OrderService) GetAllOrderItems() ([]ordermodels.OrderItem, error) {
	return orderrepo.GetAllOrderItems(os.DB)
}