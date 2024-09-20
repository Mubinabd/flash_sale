package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	pb "github.com/Mubinabd/flash_sale/internal/pkg/genproto"
	"github.com/google/uuid"
)

type FlashSaleProductsRepo struct {
	db *sql.DB
}

func NewFlashSaleProductsRepo(db *sql.DB) *FlashSaleProductsRepo {
	return &FlashSaleProductsRepo{
		db: db,
	}
}

func (r *FlashSaleProductsRepo) CreateFlashSaleProduct(req *pb.CreateFlashSaleProductReq) (*pb.Void, error) {

	id := uuid.NewString()

	query := `INSERT INTO
		flash_sales_products	
		(id, 
		flash_sale_id, 
		product_id,
		discounted_price,
		available_quantity) 
		VALUES 
		($1, $2, $3, $4, $5)`

	_, err := r.db.Exec(query, id, req.FlashSaleId, req.ProductId, req.AvailableQuantity, req.DiscountedPrice)

	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (r *FlashSaleProductsRepo) UpdateFlashSaleProduct(req *pb.UpdateFlashSaleProductReq) (*pb.Void, error) {

	var args []interface{}
	var conditions []string

	if req.Body.FlashSaleId != "" && req.Body.FlashSaleId != "string" {
		args = append(args, req.Body.FlashSaleId)
		conditions = append(conditions, fmt.Sprintf("flash_sale_id = $%d", len(args)))
	}

	if req.Body.ProductId != "" && req.Body.ProductId != "string" {
		args = append(args, req.Body.ProductId)
		conditions = append(conditions, fmt.Sprintf("product_id = $%d", len(args)))
	}

	if req.Body.AvailableQuantity != 0 {
		args = append(args, req.Body.AvailableQuantity)
		conditions = append(conditions, fmt.Sprintf("available_quantity = $%d", len(args)))
	}
	if req.Body.DiscountedPrice != 0.0 {
		args = append(args, req.Body.DiscountedPrice)
		conditions = append(conditions, fmt.Sprintf("discounted_price = $%d", len(args)))
	}

	args = append(args, time.Now())
	conditions = append(conditions, fmt.Sprintf("updated_at = $%d", len(args)))

	query := fmt.Sprintf("UPDATE flash_sales_products SET %s WHERE id = $%d", strings.Join(conditions, ", "), len(args)+1)

	args = append(args, req.Id)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.Void{}, nil
}

func (r *FlashSaleProductsRepo) DeleteFlashSaleProduct(req *pb.GetById) (*pb.Void, error) {

	query := `UPDATE
		flash_sales_products
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

func (r *FlashSaleProductsRepo) GetFlashSaleProduct(req *pb.GetById) (*pb.FlashSaleProduct, error) {

	query := `
		SELECT 
			f.id,
			f.available_quantity,
			f.discounted_price,
			s.id,
			s.name,
			s.start_time,
			s.end_time,
			s.status,
			p.id,
			p.name,
			p.price,
			p.description,
			p.image_url,
			p.stock_quantity,
		FROM
			flash_sales_products f
		LEFT JOIN
			flash_sales s
		ON
			f.flash_sale_id = s.id
		LEFT JOIN
			products p
		ON
			f.product_id = p.id
		WHERE
			f.id = $1
		AND 
			f.deleted_at = 0`

	row := r.db.QueryRow(query, req.Id)

	res := &pb.FlashSaleProduct{}
	err := row.Scan(
		&res.Id,
		&res.AvailableQuantity,
		&res.DiscountedPrice,
		&res.FlashSale.Id,
		&res.FlashSale.Name,
		&res.FlashSale.StartTime,
		&res.FlashSale.EndTime,
		&res.FlashSale.Status,
		&res.Product.Id,
		&res.Product.Name,
		&res.Product.Price,
		&res.Product.Description,
		&res.Product.ImageUrl,
		&res.Product.StockQuantity,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *FlashSaleProductsRepo) ListAllFlashSaleProducts(req *pb.ListAllFlashSaleProductsReq) (*pb.ListAllFlashSaleProductsRes, error) {

    query := `
    SELECT 
        f.id,
        f.available_quantity,
        f.discounted_price,
        s.id,
        s.name,
        s.start_time,
        s.end_time,
        s.status,
        p.id,
        p.name,
        p.price,
        p.description,
        p.image_url,
        p.stock_quantity
    FROM
        flash_sales_products f
     JOIN
        flash_sales s
    ON
        f.flash_sale_id = s.id
     JOIN
        products p
    ON
        f.product_id = p.id
    WHERE
        f.deleted_at = 0`

    rows, err := r.db.Query(query)
    if err != nil {
        return nil, fmt.Errorf("failed to execute query: %v", err)
    }
    defer rows.Close()

    res := &pb.ListAllFlashSaleProductsRes{
		FlashSaleProducts: make([]*pb.FlashSaleProduct, 0), 
	}
	
	for rows.Next() {
		product := pb.FlashSaleProduct{
			FlashSale: &pb.FlashSale{},
			Product:  &pb.Products{},
		}
		err := rows.Scan(
			&product.Id,
			&product.AvailableQuantity,
			&product.DiscountedPrice,
			&product.FlashSale.Id,
			&product.FlashSale.Name,
			&product.FlashSale.StartTime,
			&product.FlashSale.EndTime,
			&product.FlashSale.Status,
			&product.Product.Id,
			&product.Product.Name,
			&product.Product.Price,
			&product.Product.Description,
			&product.Product.ImageUrl,
			&product.Product.StockQuantity,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
	
		res.FlashSaleProducts = append(res.FlashSaleProducts, &product)
	}
	
	res.TotalCount = int64(len(res.FlashSaleProducts))

    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("rows iteration error: %v", err)
    }

    return res, nil
}

