package handlers

import (
	"context"
	pb "flashSale_gateway/internal/pkg/genproto"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateProduct creates a new Product
// @Summary       Create Product
// @Description   Create a new Product
// @Tags          Product
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         Product body pb.CreateProductReq true "Product data"
// @Success       200  {string}  string "Product created successfully"
// @Failure       400  {string}  string "Invalid request"
// @Failure       500  {string}  string "Internal server error"
// @Router        /v1/product/create [post]
func (h *Handler) CreateProduct(c *gin.Context) {
	var req pb.CreateProductReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, gin.H{"message": "Invalid request: " + err.Error()})
		return
	}

	input, err := protojson.Marshal(&req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to marshal request:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	err = h.Producer.ProduceMessages("create-product", input)
	if err != nil {
		h.Logger.ERROR.Println("Failed to produce Kafka message:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "Product created successfully"})

}

// @Summary Get Product
// @Description Get an Product by ID
// @Tags Product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Product ID"
// @Success 200 {object} pb.Products "Product data"
// @Failure 400 {string} string "Invalid request"
// @Failure 404 {string} string "Product not found"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/product/{id} [get]
func (h *Handler) GetProduct(c *gin.Context) {
	req := pb.GetById{}
	id := c.Param("id")

	req.Id = id

	res, err := h.Clients.Product.GetProduct(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)

}

// @Summary Update Product
// @Description Update an existing Product by ID
// @Tags Product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Product body pb.UpdateProductReq true "Product update data"
// @Success 200 {string} string "message":"Product updated successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/product/update/{id} [put]
func (h *Handler) UpdateProduct(c *gin.Context) {
	var req pb.UpdateProductReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	input, err := protojson.Marshal(&req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to marshal request:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	err = h.Producer.ProduceMessages("update-product", input)
	if err != nil {
		h.Logger.ERROR.Println("Failed to produce Kafka message:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "Products updated successfully"})
}

// @Summary List Products
// @Description List Products with filters
// @Tags Product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param name query int false "Name"
// @Param price query int false "Price"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} pb.ListAllProductsRes "List of Products"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/product/list [get]
func (h *Handler) ListProducts(c *gin.Context) {
	var filter pb.ListAllProductsReq
	name := c.Query("name")
	filter.Name = name

	priceStr := c.Query("price")
	if priceStr != "" {
		price, err := strconv.Atoi(priceStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid price value"})
			return
		}
		filter.Price = float32(price)
	}

	f := pb.Pagination{}
	filter.Filter = &f

	if limit := c.Query("limit"); limit != "" {
		if value, err := strconv.Atoi(limit); err == nil {
			filter.Filter.Limit = int32(value)
		} else {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	}

	if offset := c.Query("offset"); offset != "" {
		if value, err := strconv.Atoi(offset); err == nil {
			filter.Filter.Offset = int32(value)
		} else {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	}

	resp, err := h.Clients.Product.ListAllProducts(context.Background(), &filter)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}

// @Summary Delete Product
// @Description Delete an Product by ID
// @Tags Product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Product ID"
// @Success 200 {string} string "message":"Product deleted successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/product/delete/{id} [delete]
func (h *Handler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	req := &pb.GetById{Id: id}
	_, err := h.Clients.Product.DeleteProduct(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Product deleted successfully"})
}
