package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"

	pbc "github.com/golibbek/bookshop_order/genproto/catalog_service"
	pb "github.com/golibbek/bookshop_order/genproto/order_service"
	client "github.com/golibbek/bookshop_order/service/grpc_client"
)

type orderRepo struct {
	db      *sqlx.DB
	catalog client.IServiceManager
}

// NewOrderRepo ...
func NewOrderRepo(db *sqlx.DB, client client.IServiceManager) *orderRepo {
	return &orderRepo{
		db:      db,
		catalog: client,
	}
}

func (r *orderRepo) CreateOrder(order pb.OrderReq) (pb.OrderResp, error) {
	var id string
	err := r.db.QueryRow(`
		INSERT INTO orders(order_id, book_id, description, updated_at)
		VALUES ($1, $2, $3, current_timestamp) RETURNING order_id`,
		order.Id,
		order.BookId,
		order.Description,
	).Scan(&id)
	if err != nil {
		return pb.OrderResp{}, err
	}

	orderResp, err := r.GetOrder(id)

	if err != nil {
		return pb.OrderResp{}, nil
	}

	return orderResp, nil
}

func (r *orderRepo) GetOrder(id string) (pb.OrderResp, error) {
	var order pb.OrderResp
	var bookId string

	err := r.db.QueryRow(`
		SELECT order_id, book_id, description, created_at, updated_at FROM orders
		WHERE order_id=$1 AND deleted_at IS NULL`, id).Scan(
		&order.Id,
		&bookId,
		&order.Description,
		&order.CreatedAt,
		&order.UpdatedAt,
	)
	if err != nil {
		return pb.OrderResp{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()
	book, err := r.catalog.CatalogService().GetBook(ctx, &pbc.ByIdReq{Id: bookId})
	if err != nil {
		return pb.OrderResp{}, err
	}

	order.Book = book.Name

	return order, nil
}

func (r *orderRepo) ListOrders(page, limit int64) ([]*pb.OrderResp, int64, error) {
	offset := (page - 1) * limit

	rows, err := r.db.Queryx(`
		SELECT order_id, book_id, description, created_at FROM orders
		WHERE deleted_at IS NULL
		LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}
	defer rows.Close() // nolint:errcheck

	var (
		orders []*pb.OrderResp
		count  int64
	)

	for rows.Next() {
		var order pb.OrderResp
		var bookId string
		err = rows.Scan(
			&order.Id,
			&bookId,
			&order.Description,
			&order.CreatedAt,
		)
		if err != nil {
			return nil, 0, err
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
		defer cancel()
		book, err := r.catalog.CatalogService().GetBook(ctx, &pbc.ByIdReq{Id: bookId})
		if err != nil {
			return nil, 0, err
		}

		order.Book = book.Name

		orders = append(orders, &order)
	}

	err = r.db.QueryRow(`SELECT count(*) FROM orders WHERE deleted_at IS NULL`).Scan(&count)
	if err != nil {
		return nil, 0, err
	}
	return orders, count, nil
}

func (r *orderRepo) UpdateOrder(order pb.OrderReq) (pb.OrderResp, error) {
	result, err := r.db.Exec(`
		UPDATE orders SET book_id=$1, description=$2, updated_at=current_timestamp
		WHERE order_id=$3 AND deleted_at IS NULL`,
		order.BookId,
		order.Description,
		order.Id,
	)
	if err != nil {
		return pb.OrderResp{}, err
	}
	if i, _ := result.RowsAffected(); i == 0 {
		return pb.OrderResp{}, sql.ErrNoRows
	}

	orderResp, err := r.GetOrder(order.Id)
	if err != nil {
		return pb.OrderResp{}, err
	}

	return orderResp, nil
}

func (r *orderRepo) DeleteOrder(id string) error {
	result, err := r.db.Exec(`
		UPDATE orders SET deleted_at=current_timestamp WHERE order_id=$1 AND deleted_at IS NULL`, id)
	if err != nil {
		return err
	}
	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}
	return nil
}
