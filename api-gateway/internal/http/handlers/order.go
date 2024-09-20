package handlers

import (
	"context"
	pb "flashSale_gateway/internal/pkg/genproto"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateOrder creates a new Order
// @Summary       Create Order
// @Description   Create a new Order
// @Tags          Order
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         Order body pb.CreateOrderReq true "Order data"
// @Success       200  {string}  string "Order created successfully"
// @Failure       400  {string}  string "Invalid request"
// @Failure       500  {string}  string "Internal server error"
// @Router        /v1/order/create [post]
func (h *Handler) CreateOrder(c *gin.Context) {
	var req pb.CreateOrderReq
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

	err = h.Producer.ProduceMessages("create-order", input)
	if err != nil {
		h.Logger.ERROR.Println("Failed to produce Kafka message:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "Order created successfully"})

}

// @Summary Get Order
// @Description Get an Order by ID
// @Tags Order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Order ID"
// @Success 200 {object} pb.Order "Order data"
// @Failure 400 {string} string "Invalid request"
// @Failure 404 {string} string "Order not found"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/order/{id} [get]
func (h *Handler) GetOrder(c *gin.Context) {
	req := pb.GetById{}
	id := c.Param("id")

	req.Id = id

	res, err := h.Clients.Order.GetOrder(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)

}

// @Summary Update Order
// @Description Update an existing Order by ID
// @Tags Order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Order body pb.UpdateOrderReq true "Order update data"
// @Success 200 {string} string "message":"Order updated successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/order/update/{id} [put]
func (h *Handler) UpdateOrder(c *gin.Context) {
	var req pb.UpdateOrderReq
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

	err = h.Producer.ProduceMessages("update-order", input)
	if err != nil {
		h.Logger.ERROR.Println("Failed to produce Kafka message:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}
	c.JSON(200, gin.H{"message": "Orders updated successfully"})
}

// @Summary List Orders
// @Description List Orders with filters
// @Tags Order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param userID query int false "UserID"
// @Param orderStatus query int false "Order Status"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} pb.ListAllOrdersRes "List of Orders"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/order/list [get]
func (h *Handler) ListOrders(c *gin.Context) {
	var filter pb.ListAllOrdersReq
	userID := c.Query("user_id")
	orderStatus := c.Query("status")
	filter.UserID = userID
	filter.OrderStatus = orderStatus

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

	resp, err := h.Clients.Order.ListAllOrders(context.Background(), &filter)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}

// @Summary Delete Order
// @Description Delete an Order by ID
// @Tags Order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Order ID"
// @Success 200 {string} string "message":"Order deleted successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/order/delete/{id} [delete]
func (h *Handler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")

	req := &pb.GetById{Id: id}
	_, err := h.Clients.Order.DeleteOrder(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Order deleted successfully"})
}

// @Summary       Get Order History
// @Description   Retrieve a user's order history with pagination
// @Tags          Order
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param UserId query int false "User ID"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success       200  {object} pb.OrderHistoryRes "Order history response"
// @Failure       400  {string}  string "Invalid request"
// @Failure       500  {string}  string "Internal server error"
// @Router        /v1/order/history [get]
func (h *Handler) GetOrderHistory(c *gin.Context) {
	var req pb.OrderHistoryReq
	id := c.Query("user_id")
	req.UserID = id

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

	res, err := h.Clients.Order.GetOrderHistory(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
}

// @Summary       Cancel Order
// @Description   Cancel an order and initiate a refund
// @Tags          Order
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         id path string true "Order ID"
// @Success       200  {object} pb.CancelOrderRes "Cancellation response"
// @Failure       400  {string}  string "Invalid request"
// @Failure       404  {string}  string "Order not found"
// @Failure       500  {string}  string "Internal server error"
// @Router        /v1/order/{id}/cancel [post]

func (h *Handler) CancelOrder(c *gin.Context) {
	var req pb.GetById
	id := c.Param("id")
	req.Id = id
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, gin.H{"message": "Invalid request: " + err.Error()})
		return
	}

	_, err := h.Clients.FlashSale.CancelFlashSale(context.Background(), &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to add product to flash sale:", err)
		c.JSON(500, "Internal server error: "+err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "Flash sale cancel successfully"})

}