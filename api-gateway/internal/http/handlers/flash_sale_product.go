package handlers

import (
	"context"
	pb "flashSale_gateway/internal/pkg/genproto"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateFlashSale creates a new FlashSale
// @Summary       Create FlashSale
// @Description   Create a new FlashSale
// @Tags          FlashSaleProduct
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         FlashSale body pb.CreateFlashSaleProductReq true "FlashSale data"
// @Success       200  {string}  string "Flash Sale Product created successfully"
// @Failure       400  {string}  string "Invalid request"
// @Failure       500  {string}  string "Internal server error"
// @Router        /v1/flashSaleProduct/create [post]
func (h *Handler) CreateFlashSaleProduct(c *gin.Context) {
	var req pb.CreateFlashSaleProductReq
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

	err = h.Producer.ProduceMessages("create-flash-sale", input)
	if err != nil {
		h.Logger.ERROR.Println("Failed to produce Kafka message:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "Flash sale product created successfully"})

}

// @Summary Get FlashSaleProduct
// @Description Get an FlashSaleProduct by ID
// @Tags FlashSaleProduct
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "FlashSaleProduct ID"
// @Success 200 {object} pb.FlashSaleProduct "FlashSaleProduct data"
// @Failure 400 {string} string "Invalid request"
// @Failure 404 {string} string "Flash Sale Product not found"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/flashSaleProduct/{id} [get]
func (h *Handler) GetFlashSaleProduct(c *gin.Context) {
	req := pb.GetById{}
	id := c.Param("id")

	req.Id = id

	res, err := h.Clients.FlashSaleProduct.GetFlashSaleProduct(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)

}

// @Summary Update Flash Sale Product
// @Description Update an existing Flash Sale Product by ID
// @Tags FlashSaleProduct
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param FlashSaleProduct body pb.UpdateFlashSaleProductReq true "FlashSaleProduct update data"
// @Success 200 {string} string "message":"Flash Sale Product updated successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/flashSaleProduct/update/{id} [put]
func (h *Handler) UpdateFlashSaleProduct(c *gin.Context) {
	var req pb.UpdateFlashSaleProductReq
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

	err = h.Producer.ProduceMessages("update-flash-sale", input)
	if err != nil {
		h.Logger.ERROR.Println("Failed to produce Kafka message:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "Flash sale product updated successfully"})
}

// @Summary List Flash Sale Products
// @Description List Flash Sale Products with filters
// @Tags FlashSaleProduct
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Param discountPrice query int false "DiscountPrice"
// @Success 200 {object} pb.ListAllFlashSaleProductsRes "List of Flash Sale Products"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/flashSaleProduct/list [get]
func (h *Handler) ListFlashSaleProducts(c *gin.Context) {
	var filter pb.ListAllFlashSaleProductsReq
	discountPriceSTR := c.Query("discountPrice")
	if discountPriceSTR != "" {
		discountPrice, err := strconv.ParseFloat(discountPriceSTR, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid discount price"})
			return
		}
		filter.DiscountedPrice = float32(discountPrice)
	} else {
		filter.DiscountedPrice = 0
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

	resp, err := h.Clients.FlashSaleProduct.ListAllFlashSaleProducts(context.Background(), &filter)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}

// @Summary Delete FlashSaleProduct
// @Description Delete an group by ID
// @Tags FlashSaleProduct
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "FlashSaleProduct ID"
// @Success 200 {string} string "message":"Flash Sale Product deleted successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/flashSaleProduct/delete/{id} [delete]
func (h *Handler) DeleteFlashSaleProduct(c *gin.Context) {
	id := c.Param("id")

	req := &pb.GetById{Id: id}
	_, err := h.Clients.FlashSaleProduct.DeleteFlashSaleProduct(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Flash Sale Product deleted successfully"})
}
