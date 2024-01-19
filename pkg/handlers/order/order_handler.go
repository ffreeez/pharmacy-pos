package orderhandler

import (
	"strconv"

	ordermodels "pharmacy-pos/pkg/db/models/order"
	orderservice "pharmacy-pos/pkg/service/order"
	"pharmacy-pos/pkg/util/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// OrderHandler 处理订单相关的 HTTP 请求
type OrderHandler struct {
	OrderService *orderservice.OrderService
}

// NewOrderHandler 创建一个新的 OrderHandler 实例
func NewOrderHandler(db *gorm.DB) *OrderHandler {
	return &OrderHandler{
		OrderService: orderservice.NewOrderService(db),
	}
}

// CreateOrder 创建新订单
func (oh *OrderHandler) CreateOrder(c *gin.Context) {
	var order ordermodels.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		response.BadRequest(c, "Invalid input")
		return
	}

	err := oh.OrderService.CreateOrder(&order)
	if err != nil {
		response.InternalServerError(c, "Failed to create order")
		return
	}

	response.Created(c, order, "success")
}

// GetOrderByID 根据ID获取订单
func (oh *OrderHandler) GetOrderByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid order ID")
		return
	}

	orderID := uint(id)
	order, err := oh.OrderService.GetOrderByID(orderID)
	if err != nil {
		response.InternalServerError(c, "Failed to get order")
		return
	}

	response.OK(c, order, "success")
}

// UpdateOrder 更新订单信息
func (oh *OrderHandler) UpdateOrder(c *gin.Context) {
	var order ordermodels.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		response.BadRequest(c, "Invalid input")
		return
	}

	err := oh.OrderService.UpdateOrder(&order)
	if err != nil {
		response.InternalServerError(c, "Failed to update order")
		return
	}

	response.OK(c, order, "success")
}

// DeleteOrderByID 根据ID删除订单
func (oh *OrderHandler) DeleteOrderByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid order ID")
		return
	}

	orderID := uint(id)
	err = oh.OrderService.DeleteOrderByID(orderID)
	if err != nil {
		response.InternalServerError(c, "Failed to delete order")
		return
	}

	response.OK(c, gin.H{"message": "Order deleted successfully"}, "success")
}

// GetAllOrders 获取所有订单信息
func (oh *OrderHandler) GetAllOrders(c *gin.Context) {
	orders, err := oh.OrderService.GetAllOrders()
	if err != nil {
		response.InternalServerError(c, "Failed to get all orders")
		return
	}

	response.OK(c, orders, "success")
}

// CreateOrderItem 创建新订单项
func (oh *OrderHandler) CreateOrderItem(c *gin.Context) {
	var orderItem ordermodels.OrderItem
	if err := c.ShouldBindJSON(&orderItem); err != nil {
		response.BadRequest(c, "Invalid input")
		return
	}

	err := oh.OrderService.CreateOrderItem(&orderItem)
	if err != nil {
		response.InternalServerError(c, "Failed to create order item")
		return
	}

	response.Created(c, orderItem, "success")
}

// UpdateOrderItem 更新订单项
func (oh *OrderHandler) UpdateOrderItem(c *gin.Context) {
	var orderItem ordermodels.OrderItem
	if err := c.ShouldBindJSON(&orderItem); err != nil {
		response.BadRequest(c, "Invalid input")
		return
	}

	err := oh.OrderService.UpdateOrderItem(&orderItem)
	if err != nil {
		response.InternalServerError(c, "Failed to update order item")
		return
	}

	response.OK(c, orderItem, "success")
}

// DeleteOrderItemByID 根据ID删除订单项
func (oh *OrderHandler) DeleteOrderItemByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid order item ID")
		return
	}

	orderItemID := uint(id)
	err = oh.OrderService.DeleteOrderItemByID(orderItemID)
	if err != nil {
		response.InternalServerError(c, "Failed to delete order item")
		return
	}

	response.OK(c, gin.H{"message": "Order item deleted successfully"}, "success")
}

// GetOrderItemsByOrderID 根据订单ID获取所有订单项
func (oh *OrderHandler) GetOrderItemsByOrderID(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid order ID")
		return
	}

	orderItems, err := oh.OrderService.GetOrderItemsByOrderID(uint(orderID))
	if err != nil {
		response.InternalServerError(c, "Failed to get order items")
		return
	}

	response.OK(c, orderItems, "success")
}
