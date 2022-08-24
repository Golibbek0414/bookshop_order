package storage
import (
	"github.com/jmoiron/sqlx"

	client "github.com/golibbek/bookshop_order/service/grpc_client"
	"github.com/golibbek/bookshop_order/storage/postgres"
	"github.com/golibbek/bookshop_order/storage/repo"
)
// IStorage ...
type IStorage interface {
	Order() repo.OrderStorageI
}
type storagePg struct {
	db        *sqlx.DB
	orderRepo repo.OrderStorageI
}

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB, client client.IServiceManager) *storagePg {
	return &storagePg{
		db:        db,
		orderRepo: postgres.NewOrderRepo(db, client),
	}
}

func (s storagePg) Order() repo.OrderStorageI {
	return s.orderRepo
}