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
// @Tags          FlashSale
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         FlashSale body pb.CreateFlashSalesReq true "FlashSale data"
// @Success       200  {string}  string "Flash Sale created successfully"
// @Failure       400  {string}  string "Invalid request"
// @Failure       500  {string}  string "Internal server error"
// @Router        /v1/flashSale/create [post]
func (h *Handler) CreateFlashSale(c *gin.Context) {
	var req pb.CreateFlashSalesReq
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

	err = h.Producer.ProduceMessages("create-flash", input)
	if err != nil {
		h.Logger.ERROR.Println("Failed to produce Kafka message:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}
	c.JSON(200, gin.H{"message": "Flash sale created successfully"})

}

// @Summary Get FlashSale
// @Description Get an FlashSale by ID
// @Tags FlashSale
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "FlashSale ID"
// @Success 200 {object} pb.FlashSale "FlashSale data"
// @Failure 400 {string} string "Invalid request"
// @Failure 404 {string} string "FlashSale not found"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/flashSale/{id} [get]
func (h *Handler) GetFlashSale(c *gin.Context) {
	req := pb.GetById{}
	id := c.Param("id")

	req.Id = id

	res, err := h.Clients.FlashSale.GetFlashSale(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)

}

// @Summary Update FlashSale
// @Description Update an existing FlashSale by ID
// @Tags FlashSale
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param FlashSale body pb.UpdateFlashSalesReq true "FlashSale update data"
// @Success 200 {string} string "message":"Flash Sale updated successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/flashSale/update/{id} [put]
func (h *Handler) UpdateFlashSale(c *gin.Context) {
	var req pb.UpdateFlashSalesReq
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

	err = h.Producer.ProduceMessages("update-flash", input)
	if err != nil {
		h.Logger.ERROR.Println("Failed to produce Kafka message:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "Flash sale updated successfully"})
}

// @Summary List FlashSales
// @Description List FlashSales with filters
// @Tags FlashSale
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param name query int false "Name"
// @Param status query int false "Status"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} pb.ListAllFlashSalesRes "List of FlashSales"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/flashSale/list [get]
func (h *Handler) ListFlashSales(c *gin.Context) {
	var filter pb.ListAllFlashSalesReq
	name := c.Query("name")
	status := c.Query("status")
	filter.Name = name
	filter.Status = status

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

	resp, err := h.Clients.FlashSale.ListAllFlashSales(context.Background(), &filter)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}

// @Summary Delete FlashSale
// @Description Delete an FlashSale by ID
// @Tags FlashSale
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "FlashSale ID"
// @Success 200 {string} string "message":"Flash Sale deleted successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/FlashSale/delete/{id} [delete]
func (h *Handler) DeleteFlashSale(c *gin.Context) {
	id := c.Param("id")

	req := &pb.GetById{Id: id}
	_, err := h.Clients.FlashSale.DeleteFlashSale(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Flash Sale deleted successfully"})
}

// @Summary       Add Product to Flash Sale
// @Description   Add a product to an existing flash sale
// @Tags          FlashSale
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         AddProductReq body pb.AddProductReq true "Add Product Request"
// @Success       200  {object} pb.Void "Product added to flash sale successfully"
// @Failure       400  {string}  string "Invalid request"
// @Failure       500  {string}  string "Internal server error"
// @Router        /v1/flashSale/products [post]
func (h *Handler) AddProductToFlashSale(c *gin.Context) {
	var req pb.AddProductReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, gin.H{"message": "Invalid request: " + err.Error()})
		return
	}

	_, err := h.Clients.FlashSale.AddProductToFlashSale(context.Background(), &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to add product to flash sale:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "Flash sale added successfully"})

}

// @Summary       Remove Product from Flash Sale
// @Description   Remove a product from an existing flash sale
// @Tags          FlashSale
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         RemoveProductReq body pb.RemoveProductReq true "Remove Product Request"
// @Success       200  {object} pb.Void "Product removed from flash sale successfully"
// @Failure       400  {string}  string "Invalid request"
// @Failure       500  {string}  string "Internal server error"
// @Router        /v1/flashSale/products [delete]
func (h *Handler) RemoveProductFromFlashSale(c *gin.Context) {
	flashSaleId := c.Query("id")
	productId := c.Query("id")

	req := &pb.RemoveProductReq{FlashSaleId: flashSaleId, ProductId: productId}
	_, err := h.Clients.FlashSale.RemoveProductFromFlashSale(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Flash Sale removed successfully"})
}

// @Summary       Cancel Flash Sale
// @Description   Cancel an existing flash sale
// @Tags          FlashSale
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         id path string true "Flash Sale ID"
// @Success       200  {object} pb.CancelFlashSaleRes "Flash sale cancelled successfully"
// @Failure       400  {string}  string "Invalid request"
// @Failure       404  {string}  string "Flash Sale not found"
// @Failure       500  {string}  string "Internal server error"
// @Router        /v1/flashSale/{id}/cancel [post]
func (h *Handler) CancelFlashSale(c *gin.Context) {

	req := &pb.GetById{}
	id := c.Param("id")

	req.Id = id

	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, gin.H{"message": "Invalid request: " + err.Error()})
		return
	}

	_, err := h.Clients.FlashSale.CancelFlashSale(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Flash sale cancelled successfully"})

}

// @Summary       Get Nearby Flash Sales
// @Description   Retrieve flash sales near a specific location
// @Tags          FlashSale
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         Latitude   query float64 true "Latitude of the location"
// @Param         Longitude  query float64 true "Longitude of the location"
// @Param         Radius     query float64 true "Search radius in meters"
// @Param         Limit      query int     false "Limit number of results"
// @Param         Offset     query int     false "Offset for pagination"
// @Success       200  {object} pb.NearbyFlashSalesRes "List of nearby flash sales"
// @Failure       400  {string}  string "Invalid request"
// @Failure       500  {string}  string "Internal server error"
// @Router        /v1/flashSale/nearby [get]
func (h *Handler) GetNearbyFlashSales(c *gin.Context) {

	var req pb.GetNearbyFlashSalesReq

	latitude := c.Query("latitude")
	longitude := c.Query("longitude")
	radius := c.Query("radius")

	req.Latitude, _ = strconv.ParseFloat(latitude, 64)
	req.Longitude, _ = strconv.ParseFloat(longitude, 64)
	req.Radius, _ = strconv.ParseFloat(radius, 64)

	filter := &pb.Pagination{}

	if limit := c.Query("limit"); limit != "" {
		if value, err := strconv.Atoi(limit); err == nil {
			filter.Limit = int32(value)
		} else {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	}

	if offset := c.Query("offset"); offset != "" {
		if value, err := strconv.Atoi(offset); err == nil {
			filter.Offset = int32(value)
		} else {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	}

	res, err := h.Clients.FlashSale.GetNearbyFlashSales(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)
}

// @Summary       Get Flash Sale  Location
// @Description   Retrieve the location details of a Flash Sale  by its ID
// @Tags          FlashSale 
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         Flash Sale Id path string true "Flash Sale  ID"
// @Success       200  {object} pb.StoreLocation "Flash Sale  location details"
// @Failure       400  {string}  string "Invalid request"
// @Failure       404  {string}  string "Store not found"
// @Failure       500  {string}  string "Internal server error"
// @Router        /v1/flashSale/{id}/location [get]

func(h *Handler)GetStoreLocation(c *gin.Context){
	req := &pb.GetStoreLocationReq{
		StoreId: c.Param("storeId"),
	}

	res, err := h.Clients.FlashSale.GetStoreLocation(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)


}
