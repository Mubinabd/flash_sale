package handlers

import (
	"context"
	pb "flashSale_gateway/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
)

// @Summary       Share a Deal
// @Description   Share a flash sale deal on a specified platform
// @Tags          Social
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         ShareDealReq body pb.ShareDealReq true "Share Deal Request"
// @Success       200  {object} pb.Void "Deal shared successfully"
// @Failure       400  {string}  string "Invalid request"
// @Failure       500  {string}  string "Internal server error"
// @Router        /v1/deals/share [post]

func (h *Handler) ShareDeal(c *gin.Context) {
	var req pb.ShareDealReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, gin.H{"message": "Invalid request: " + err.Error()})
		return
	}

	_, err := h.Clients.Social.ShareDeal(context.Background(), &req) // Pass a pointer to req
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Deal shared successfully"})
}

// @Summary       Get Sharing Stats
// @Description   Retrieve sharing statistics for a flash sale deal
// @Tags          Social
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         GetSharingStatsReq body pb.GetSharingStatsReq true "Get Sharing Stats Request"
// @Success       200  {object} pb.SharingStatsRes "Sharing statistics response"
// @Failure       400  {string}  string "Invalid request"
// @Failure       404  {string}  string "Flash sale not found"
// @Failure       500  {string}  string "Internal server error"
// @Router        /v1/deals/{flashSaleId}/sharing [get]

func (h *Handler) GetSharingStats(c *gin.Context) {
	var req pb.GetSharingStatsReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Logger.ERROR.Println("Failed to bind request:", err)
		c.JSON(400, gin.H{"message": "Invalid request: " + err.Error()})
		return
	}

	res, err := h.Clients.Social.GetSharingStats(context.Background(), &req) // Pass a pointer to req
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)
}
