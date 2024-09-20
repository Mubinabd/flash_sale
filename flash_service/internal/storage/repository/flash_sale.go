package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	pb "github.com/Mubinabd/flash_sale/internal/pkg/genproto"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FlashSaleRepo struct {
	db *sql.DB
}

func NewFlashSaleRepo(db *sql.DB) *FlashSaleRepo {
	return &FlashSaleRepo{
		db: db,
	}
}

func (r *FlashSaleRepo) CreateFlashSale(req *pb.CreateFlashSalesReq) (*pb.Void, error) {
	id := uuid.NewString()

	query := `INSERT INTO flash_sales 
			(id, 
			name, 
			product_id,
			start_time, 
			end_time, 
			status, 
			latitude, 
			longitude, 
			address) 
			VALUES 
			($1, $2, $3, $4, $5, $6, $7, $8,$9)`

	_, err := r.db.Exec(query, id, req.Name, req.StartTime, req.EndTime, req.Status,req.ProductID,nil,nil,nil)

	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (r *FlashSaleRepo) UpdateFlashSale(req *pb.UpdateFlashSalesReq) (*pb.Void, error) {
	var args []interface{}
	var conditions []string

	if req.Body.Name != "" {
		args = append(args, req.Body.Name)
		conditions = append(conditions, fmt.Sprintf("name = $%d", len(args)))
	}

	if req.Body.StartTime != "" {
		args = append(args, req.Body.StartTime)
		conditions = append(conditions, fmt.Sprintf("start_time = $%d", len(args)))
	}

	if req.Body.EndTime != "" {
		args = append(args, req.Body.EndTime)
		conditions = append(conditions, fmt.Sprintf("end_time = $%d", len(args)))
	}

	if req.Body.Status != "" {
		args = append(args, req.Body.Status)
		conditions = append(conditions, fmt.Sprintf("status = $%d", len(args)))
	}

	if len(conditions) > 0 {
		args = append(args, time.Now())
		conditions = append(conditions, fmt.Sprintf("updated_at = $%d", len(args)))
		args = append(args, req.Id)

		query := `UPDATE flash_sales SET ` + strings.Join(conditions, ", ") + ` WHERE id = $` + fmt.Sprintf("%d", len(args))
		_, err := r.db.Exec(query, args...)
		if err != nil {
			log.Println("Error while updating flash_sales", err)
			return nil, err
		}
	}

	// Update products in flash sale
	for _, product := range req.Body.Products {
		_, err := r.db.Exec(`INSERT INTO flash_sale_products (flash_sale_id, product_id, added_at) VALUES ($1, $2, NOW())`,
			req.Id, product.Id)
		if err != nil {
			return nil, err
		}
	}

	return &pb.Void{}, nil
}

func (r *FlashSaleRepo) GetFlashSale(req *pb.GetById) (*pb.FlashSale, error) {
	query := `
			SELECT 
				id,
				name,
				start_time,
				end_time,
				status,
				created_at
			FROM
				flash_sales
			WHERE
				id = $1
			AND 
				deleted_at = 0`

	row := r.db.QueryRow(query, req.Id)

	res := &pb.FlashSale{}
	err := row.Scan(
		&res.Id,
		&res.Name,
		&res.StartTime,
		&res.EndTime,
		&res.Status,
		&res.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	// Get products for the flash sale
	rows, err := r.db.Query(`SELECT p.id, p.name, p.description, p.image_url, p.price, p.quantity_available
							  FROM flash_sale_products fsp
							  JOIN products p ON p.id = fsp.product_id
							  WHERE fsp.flash_sale_id = $1`, req.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*pb.Product
	for rows.Next() {
		var product pb.Product
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.ImageUrl,
			&product.Price,
			&product.QuantityAvailable,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	res.Products = products

	return res, nil
}

func (r *FlashSaleRepo) DeleteFlashSale(req *pb.GetById) (*pb.Void, error) {
	query := `
			UPDATE
				flash_sales
			SET
				deleted_at = extract(epoch from now())
			WHERE
				id = $1`

	_, err := r.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (r *FlashSaleRepo) ListAllFlashSales(req *pb.ListAllFlashSalesReq) (*pb.ListAllFlashSalesRes, error) {
	query := `
			SELECT 
				id,
				name,
				start_time,
				end_time,
				status,
				created_at
			FROM
				flash_sales
			WHERE
				deleted_at = 0`

	var args []interface{}
	if req.Name != "" {
		args = append(args, req.Name)
		query += ` AND name = $1`
	}
	if req.Status != "" {
		args = append(args, req.Status)
		query += ` AND status = $2`
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flashSales []*pb.FlashSale
	for rows.Next() {
		var flashSale pb.FlashSale
		err := rows.Scan(
			&flashSale.Id,
			&flashSale.Name,
			&flashSale.StartTime,
			&flashSale.EndTime,
			&flashSale.Status,
			&flashSale.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		flashSales = append(flashSales, &flashSale)
	}
	totalCount := int64(len(flashSales))

	return &pb.ListAllFlashSalesRes{
		FlashSales: flashSales,
		TotalCount: totalCount,
		Limit:      req.Filter.Limit,
		Offset:     req.Filter.Offset,
	}, nil
}

func (r *FlashSaleRepo) AddProductToFlashSale(req *pb.AddProductReq) (*pb.Void, error) {
	_, err := r.db.Exec(`INSERT INTO flash_sale_products (flash_sale_id, product_id, added_at) VALUES ($1, $2, NOW())`,
		req.FlashSaleId, req.Product.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (r *FlashSaleRepo) RemoveProductFromFlashSale(req *pb.RemoveProductReq) (*pb.Void, error) {
	_, err := r.db.Exec(`DELETE FROM flash_sale_products WHERE flash_sale_id = $1 AND product_id = $2`,
		req.FlashSaleId, req.ProductId)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (r *FlashSaleRepo) CancelFlashSale(req *pb.GetById) (*pb.CancelFlashSaleRes, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`UPDATE flash_sales SET status = 'canceled', updated_at = NOW() WHERE id = $1`, req.Id)
	if err != nil {
		return nil, err
	}

	var cancellationID string
	err = tx.QueryRow(`INSERT INTO flash_sale_cancellations (flash_sale_id, cancellation_status, created_at) VALUES ($1, 'canceled', NOW()) RETURNING id`,
		req.Id).Scan(&cancellationID)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &pb.CancelFlashSaleRes{CancellationStatus: cancellationID}, nil
}

func (s *FlashSaleRepo) GetNearbyFlashSales(req *pb.GetNearbyFlashSalesReq) (*pb.NearbyFlashSalesRes, error) {
	rows, err := s.db.Query(`
        SELECT id AS flash_sale_id, name, description, latitude, longitude, address
        FROM flash_sales
        WHERE ST_DistanceSphere(
                ST_MakePoint(longitude, latitude),
                ST_MakePoint($2, $1)
              ) <= $3
        ORDER BY ST_DistanceSphere(
                ST_MakePoint(longitude, latitude),
                ST_MakePoint($2, $1)
              )
        LIMIT $4 OFFSET $5`,
		req.Latitude, req.Longitude, req.Radius)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flashSales []*pb.FlashSaleWithLocation
	for rows.Next() {
		var sale pb.FlashSaleWithLocation
		if err := rows.Scan(&sale.FlashSaleId, &sale.Name, &sale.Description, &sale.Latitude, &sale.Longitude, &sale.Address); err != nil {
			return nil, err
		}
		flashSales = append(flashSales, &sale)
	}

	var totalCount int64
	err = s.db.QueryRow(`
        SELECT COUNT(*)
        FROM flash_sales
        WHERE ST_DistanceSphere(
                ST_MakePoint(longitude, latitude),
                ST_MakePoint($2, $1)
              ) <= $3`,
		req.Latitude, req.Longitude, req.Radius).Scan(&totalCount)
	if err != nil {
		return nil, err
	}

	return &pb.NearbyFlashSalesRes{
		FlashSales: flashSales,
		TotalCount: totalCount,
	}, nil
}

func (s *FlashSaleRepo) GetStoreLocation(req *pb.GetStoreLocationReq) (*pb.StoreLocation, error) {
	var store pb.StoreLocation
	err := s.db.QueryRow(`
        SELECT store_id, name, address, latitude, longitude
        FROM stores
        WHERE store_id = $1`, req.StoreId).Scan(
		&store.StoreId, &store.Name, &store.Address, &store.Latitude, &store.Longitude)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "Store not found")
		}
		return nil, err
	}

	return &store, nil
}
