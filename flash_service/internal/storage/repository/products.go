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

type ProductsRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductsRepo {
	return &ProductsRepo{
		db: db,
	}
}

func (p *ProductsRepo) CreateProduct(req *pb.CreateProductReq) (*pb.Void, error) {
	id := uuid.NewString()

	query := `INSERT INTO 
		products 
			(id, 
			name, 
			description,
			price,
			image_url,
			stock_quantity) 
		VALUES 
			($1, $2, $3, $4, $5, $6)`

	_, err := p.db.Exec(query, id, req.Name, req.Description, req.Price, req.ImageUrl, req.StockQuantity)

	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (p *ProductsRepo) UpdateProduct(req *pb.UpdateProductReq) (*pb.Void, error) {
	var args []interface{}
	var conditions []string

	if req.Body.Name != "" && req.Body.Name != "string" {
		args = append(args, req.Body.Name)
		conditions = append(conditions, fmt.Sprintf("name = $%d", len(args)))
	}

	if req.Body.Price != 0.0 {
		args = append(args, req.Body.Price)
		conditions = append(conditions, fmt.Sprintf("price = $%d", len(args)))
	}

	args = append(args, time.Now())
	conditions = append(conditions, fmt.Sprintf("updated_at = $%d", len(args)))

	if len(conditions) > 0 {
		query := `UPDATE products SET ` + strings.Join(conditions, ", ") + ` WHERE id = $` + fmt.Sprintf("%d", len(args)+1)
		args = append(args, req.Id)

		_, err := p.db.Exec(query, args...)
		if err != nil {
			log.Println("Error while updating products", err)
			return nil, err
		}
	}

	return &pb.Void{}, nil

}
func (p *ProductsRepo) GetProduct(req *pb.GetById) (*pb.Products, error) {
	query := `
		SELECT 
			id,
			name,
			description,
			price,
			image_url,
			stock_quantity,
			created_at
		FROM
			products 
		WHERE
			deleted_at = 0 
		AND 
			id = $1`

	res := &pb.Products{}

	err := p.db.QueryRow(query, req.Id).
		Scan(
			&res.Id,
			&res.Name,
			&res.Description,
			&res.Price,
			&res.ImageUrl,
			&res.StockQuantity,
			&res.CreatedAt)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (p *ProductsRepo) ListAllProducts(req *pb.ListAllProductsReq) (*pb.ListAllProductsRes, error) {

	query := `
		SELECT 
			id,
			name,
			description,
			price,
			image_url,
			stock_quantity
		FROM
			products 
		WHERE
			deleted_at = 0
			`

	var args []interface{}
	argCount := 1
	filters := []string{}

	if req.Name != "" && req.Name != "string" {
		filters = append(filters, fmt.Sprintf("name = $%d", len(args)+1))
		args = append(args, req.Name)
	}

	if req.Price != 0 && req.Price != 0.0 {
		filters = append(filters, fmt.Sprintf("price = $%d", len(args)+1))
		args = append(args, req.Price)
	}

	if len(filters) > 0 {
		query += " WHERE " + strings.Join(filters, " AND ")
	}

	query += " ORDER BY created_at DESC"

	if req.Filter != nil {
		if req.Filter.Limit > 0 {
			query += fmt.Sprintf(" LIMIT $%d", argCount)
			args = append(args, req.Filter.Limit)
			argCount++
		}
		if req.Filter.Offset > 0 {
			query += fmt.Sprintf(" OFFSET $%d", argCount)
			args = append(args, req.Filter.Offset)
			argCount++
		}
	}

	rows, err := p.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []*pb.Products{}

	for rows.Next() {
		var product pb.Products

		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Price,        
			&product.ImageUrl,
			&product.StockQuantity,
		)
		if err != nil {
			log.Println("Scan error: ", err)
			return nil, err
		}
		products = append(products, &product)
	}

	totalCount := len(products)
	return &pb.ListAllProductsRes{
		Products:   products,
		TotalCount: int64(totalCount),
		Limit:      req.Filter.Limit,
		Offset:     req.Filter.Offset,
	}, nil
}


func (p *ProductsRepo) DeleteProduct(req *pb.GetById) (*pb.Void, error) {
	query := `
	UPDATE
		products
	SET
		deleted_at = extract(epoch from now())
	WHERE
		id = $1`

	_, err := p.db.Exec(query, req.Id)

	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
