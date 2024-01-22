package constants

const (
	//Trạng thái chờ thanh toán
	OrderStatusPending = "pending_payment"
	//Trạng thái đang xử lý
	OrderStatusProcessing = "processing"
	//Trạng thái chờ xác nhận đơn hàng
	OrderStatusConfirmed = "pending_confirmation"
	//Trạng thái hoàn thành
	OrderStatusCompleted = "completed"
	//Trạng thái hủy
	OrderStatusCancelled = "cancelled"
	//Trạng thái trả lại
	OrderStatusRefunded = "refunded"
)
