package orderrepo

import (
	ordermodels "pharmacy-pos/pkg/db/models/order"
	logger "pharmacy-pos/pkg/util/logger"

	"gorm.io/gorm"
)

var logs = logger.GetLogger()

// CreateOrder 创建新订单
func CreateOrder(db *gorm.DB, order *ordermodels.Order) error {
	result := db.Create(order)
	if result.Error != nil {
		logs.Errorf("创建新订单失败, CustomerName: %s, error: %v", order.CustomerName, result.Error)
		return result.Error
	}
	logs.Infof("创建新订单成功, CustomerName: %s, OrderID: %d", order.CustomerName, order.ID)
	return nil
}

// GetOrderByID 根据订单ID获取订单信息
func GetOrderByID(db *gorm.DB, id uint) (*ordermodels.Order, error) {
	order := &ordermodels.Order{}
	result := db.Preload("OrderItems").Preload("Member").First(order, id)
	if result.Error != nil {
		logs.Errorf("根据订单ID获取订单信息失败, ID: %d, error: %v", id, result.Error)
		return nil, result.Error
	}
	logs.Infof("根据订单ID获取订单信息成功, ID: %d", id)
	return order, nil
}

// UpdateOrder 更新订单信息
func UpdateOrder(db *gorm.DB, order *ordermodels.Order) error {
	result := db.Save(order)
	if result.Error != nil {
		logs.Errorf("更新订单信息失败, OrderID: %d, error: %v", order.ID, result.Error)
		return result.Error
	}
	logs.Infof("更新订单信息成功, OrderID: %d", order.ID)
	return nil
}

// DeleteOrderByID 根据ID删除订单
func DeleteOrderByID(db *gorm.DB, id uint) error {
	result := db.Delete(&ordermodels.Order{}, id)
	if result.Error != nil {
		logs.Errorf("根据ID删除订单失败, OrderID: %d, error: %v", id, result.Error)
		return result.Error
	}
	logs.Infof("根据ID删除订单成功, OrderID: %d", id)
	return nil
}

// GetAllOrders 获取所有订单信息
func GetAllOrders(db *gorm.DB) ([]ordermodels.Order, error) {
	var orders []ordermodels.Order
	result := db.Preload("OrderItems").Preload("Member").Find(&orders)
	if result.Error != nil {
		logs.Errorf("获取所有订单信息失败: %v", result.Error)
		return nil, result.Error
	}
	logs.Infof("获取所有订单信息成功")
	return orders, nil
}

// CreateOrderItem 创建新订单项
func CreateOrderItem(db *gorm.DB, orderItem *ordermodels.OrderItem) error {
	result := db.Create(orderItem)
	if result.Error != nil {
		logs.Errorf("创建新订单项失败, OrderID: %d, error: %v", orderItem.OrderID, result.Error)
		return result.Error
	}
	logs.Infof("创建新订单项成功, OrderID: %d, OrderItemID: %d", orderItem.OrderID, orderItem.ID)
	return nil
}

// UpdateOrderItem 更新订单项
func UpdateOrderItem(db *gorm.DB, orderItem *ordermodels.OrderItem) error {
	result := db.Save(orderItem)
	if result.Error != nil {
		logs.Errorf("更新订单项失败, OrderItemID: %d, error: %v", orderItem.ID, result.Error)
		return result.Error
	}
	logs.Infof("更新订单项成功, OrderItemID: %d", orderItem.ID)
	return nil
}

// DeleteOrderItemByID 根据ID删除订单项
func DeleteOrderItemByID(db *gorm.DB, id uint) error {
	result := db.Delete(&ordermodels.OrderItem{}, id)
	if result.Error != nil {
		logs.Errorf("根据ID删除订单项失败, OrderItemID: %d, error: %v", id, result.Error)
		return result.Error
	}
	logs.Infof("根据ID删除订单项成功, OrderItemID: %d", id)
	return nil
}

// GetOrderItemsByOrderID 根据订单ID获取所有订单项
func GetOrderItemsByOrderID(db *gorm.DB, orderID uint) ([]ordermodels.OrderItem, error) {
	var orderItems []ordermodels.OrderItem
	result := db.Where("order_id = ?", orderID).Find(&orderItems)
	if result.Error != nil {
		logs.Errorf("根据订单ID获取订单项失败, OrderID: %d, error: %v", orderID, result.Error)
		return nil, result.Error
	}
	logs.Infof("根据订单ID获取订单项成功, OrderID: %d", orderID)
	return orderItems, nil
}
