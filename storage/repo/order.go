package repo
import (
	pb "github.com/golibbek/bookshop_order/genproto/order_service"
)

// UserStorageI ...
type OrderStorageI interface {
	CreateOrder(pb.OrderReq) (pb.OrderResp, error)
	GetOrder(id string) (pb.OrderResp, error)
	ListOrders(page, limit int64) ([]*pb.OrderResp, int64, error)
	UpdateOrder(pb.OrderReq) (pb.OrderResp, error)
	DeleteOrder(id string) error
}