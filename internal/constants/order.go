package constants

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusPaid       OrderStatus = "paid"
	OrderStatusDelivering OrderStatus = "delivering"
	OrderStatusCompleted  OrderStatus = "completed"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

type OrderItemStatus string

const (
	OrderItemStatusPending    OrderItemStatus = "pending"
	OrderItemStatusPaid       OrderItemStatus = "paid"
	OrderItemStatusRefund     OrderItemStatus = "refund"
	OrderItemStatusDelivering OrderItemStatus = "delivering"
	OrderItemStatusCompleted  OrderItemStatus = "completed"
	OrderItemStatusCancelled  OrderItemStatus = "cancelled"
)
